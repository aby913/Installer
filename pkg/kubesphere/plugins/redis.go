package plugins

import (
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"bytetrade.io/web3os/installer/pkg/common"
	cc "bytetrade.io/web3os/installer/pkg/core/common"
	"bytetrade.io/web3os/installer/pkg/core/connector"
	"bytetrade.io/web3os/installer/pkg/core/logger"
	"bytetrade.io/web3os/installer/pkg/core/prepare"
	"bytetrade.io/web3os/installer/pkg/core/task"
	"bytetrade.io/web3os/installer/pkg/utils"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

// ~ CreateRedisSecret
type CreateRedisSecret struct {
	common.KubeAction
}

func (t *CreateRedisSecret) Execute(runtime connector.Runtime) error {
	redisPwd, ok := t.ModuleCache.Get(common.CacheRedisPassword)
	if !ok {
		return fmt.Errorf("get redis password from module cache failed")
	}

	if stdout, err := runtime.GetRunner().SudoCmdExt(fmt.Sprintf("/usr/local/bin/kubectl -n %s create secret generic redis-secret --from-literal=auth=%s", common.NamespaceKubesphereSystem, redisPwd), false, true); err != nil {
		if err != nil && !strings.Contains(stdout, "already exists") {
			return errors.Wrap(errors.WithStack(err), "create redis secret failed")
		}
	}

	return nil
}

// ~ BackupRedisManifests
type BackupRedisManifests struct {
	common.KubeAction
}

func (t *BackupRedisManifests) Execute(runtime connector.Runtime) error {
	rver, err := runtime.GetRunner().SudoCmdExt(fmt.Sprintf("/usr/local/bin/kubectl get pod -n %s -l app=%s,tier=database,version=%s-4.0 | wc -l",
		common.NamespaceKubesphereSystem, common.ChartNameRedis, common.ChartNameRedis), false, false)

	if err != nil || strings.Contains(rver, "No resources found") {
		return nil
	}
	rver = strings.ReplaceAll(rver, "No resources found in kubesphere-system namespace.", "")
	rver = strings.ReplaceAll(rver, "\r\n", "")
	rver = strings.ReplaceAll(rver, "\n", "")
	if rver != "0" {
		var cmd = fmt.Sprintf("/usr/local/bin/kubectl get svc -n %s %s -o yaml > %s/redis-svc-backup.yaml && /usr/local/bin/kubectl delete svc -n %s %s", common.NamespaceKubesphereSystem, common.ChartNameRedis,
			common.KubeManifestDir, common.NamespaceKubesphereSystem,
			common.ChartNameRedis)

		if _, err := runtime.GetRunner().SudoCmd(cmd, false, true); err != nil {
			logger.Errorf("failed to backup %s svc: %v", common.ChartNameRedis, err)
			return errors.Wrap(errors.WithStack(err), "backup redis svc failed")
		}
	}
	return nil
}

// ~ DeployRedisHA
type DeployRedisHA struct {
	common.KubeAction
}

func (t *DeployRedisHA) Execute(runtime connector.Runtime) error {

	return nil
}

// ~ DeployRedis
type DeployRedis struct {
	common.KubeAction
}

func (t *DeployRedis) Execute(runtime connector.Runtime) error {
	config, err := ctrl.GetConfig()
	if err != nil {
		return err
	}

	var appName = common.ChartNameRedis
	var appPath = path.Join(runtime.GetFilesDir(), cc.BuildDir, appName)

	actionConfig, settings, err := utils.InitConfig(config, common.NamespaceKubesphereSystem)
	if err != nil {
		return err
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	if err := utils.InstallCharts(ctx, actionConfig, settings, appName, appPath, "", common.NamespaceKubesphereSystem, nil); err != nil {
		return err
	}

	return nil
}

// ~ PatchRedisStatus
type PatchRedisStatus struct {
	common.KubeAction
}

func (t *PatchRedisStatus) Execute(runtime connector.Runtime) error {
	// var jsonPatch = fmt.Sprintf(`{\"status\": {\"redis\": {\"status\": \"enabled\", \"enabledTime\": \"%s\"}}}`,
	// 	time.Now().Format("2006-01-02T15:04:05Z"))
	// if runtime.GetRunner().Host.GetMinikube() {

	// }
	// todo fix
	var jsonPatch = fmt.Sprintf(`{"status": {"redis": {"status": "enabled", "enabledTime": "%s"}}}`, time.Now().Format("2006-01-02T15:04:05Z"))
	var cmd = fmt.Sprintf("/usr/local/bin/kubectl patch cc ks-installer --type merge -p '%s' -n %s", jsonPatch, common.NamespaceKubesphereSystem)

	_, err := runtime.GetRunner().SudoCmd(cmd, false, true)
	if err != nil {
		return errors.Wrap(errors.WithStack(err), "patch redis status failed")
	}

	return nil
}

// +++++

// ~ DeployRedisModule
type DeployRedisModule struct {
	common.KubeModule
}

func (m *DeployRedisModule) Init() {
	m.Name = "DeployRedis"

	createRedisSecret := &task.RemoteTask{
		Name:  "CreateRedisSecret",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
			new(NotEqualDesiredVersion),
			new(GenerateRedisPassword),
		},
		Action:   new(CreateRedisSecret),
		Parallel: false,
		Retry:    0,
	}

	backupRedisManifests := &task.RemoteTask{
		Name:  "BackupRedisManifests",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
			new(NotEqualDesiredVersion),
		},
		Action:   new(BackupRedisManifests),
		Parallel: false,
		Retry:    0,
	}

	deployRedisHA := &task.RemoteTask{
		Name:  "DeployRedisHA",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
			new(NotEqualDesiredVersion),
			new(common.GetMasterNum),
		},
		Action:   new(DeployRedisHA), // todo skip
		Parallel: false,
		Retry:    0,
	}

	deployRedis := &task.RemoteTask{
		Name:  "DeployRedis",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
			new(NotEqualDesiredVersion),
			new(CheckStorageClass),
		},
		Action:   new(DeployRedis),
		Parallel: false,
		Retry:    0,
	}

	patchRedis := &task.RemoteTask{
		Name:  "PatchRedisStatus",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
			new(NotEqualDesiredVersion),
		},
		Action:   new(PatchRedisStatus),
		Parallel: false,
		Retry:    50,
		Delay:    2 * time.Second,
	}

	m.Tasks = []task.Interface{
		createRedisSecret,
		backupRedisManifests,
		deployRedisHA, // todo
		deployRedis,
		patchRedis,
	}
}
