package storage

import (
	"fmt"
	"os/exec"
	"strings"

	kubekeyapiv1alpha2 "bytetrade.io/web3os/installer/apis/kubekey/v1alpha2"
	"bytetrade.io/web3os/installer/pkg/common"
	"bytetrade.io/web3os/installer/pkg/constants"
	cc "bytetrade.io/web3os/installer/pkg/core/common"
	"bytetrade.io/web3os/installer/pkg/core/connector"
	"bytetrade.io/web3os/installer/pkg/core/logger"
	"bytetrade.io/web3os/installer/pkg/core/util"
	"bytetrade.io/web3os/installer/pkg/files"
	"bytetrade.io/web3os/installer/pkg/utils"
	"github.com/pkg/errors"
)

// ~ MkStorageDir
type MkStorageDir struct {
	common.KubeAction
}

func (t *MkStorageDir) Execute(runtime connector.Runtime) error {
	var storageVendor, _ = t.PipelineCache.GetMustString(common.CacheStorageVendor)
	if storageVendor == "true" {
		if utils.IsExist(StorageDataDir) {
			if utils.IsExist(cc.TerminusDir) {
				_, _ = runtime.GetRunner().SudoCmdExt(fmt.Sprintf("rm -rf %s", cc.TerminusDir), false, false)
			}

			if _, err := runtime.GetRunner().SudoCmdExt(fmt.Sprintf("mkdir -p %s", StorageDataTerminusDir), false, false); err != nil {
				return err
			}
			if _, err := runtime.GetRunner().SudoCmdExt(fmt.Sprintf("ln -s %s %s", StorageDataTerminusDir, cc.TerminusDir), false, false); err != nil {
				return err
			}
		}

	}

	return nil
}

// ~ DownloadStorageCli
type DownloadStorageCli struct {
	common.KubeAction
}

func (t *DownloadStorageCli) Execute(runtime connector.Runtime) error {
	storageTypeIf, ok := t.PipelineCache.Get(common.CacheStorageType)
	if !ok || storageTypeIf == nil {
		return nil
	}
	var arch = fmt.Sprintf("%s-%s", constants.OsType, constants.OsArch)

	var binary *files.KubeBinary
	storageType := storageTypeIf.(string)
	switch storageType {
	case "s3":
		binary = files.NewKubeBinary("awscli", arch, "", runtime.GetWorkDir())
	case "oss":
		binary = files.NewKubeBinary("ossutil", arch, kubekeyapiv1alpha2.DefaultOssUtilVersion, runtime.GetWorkDir())
	default:
		return nil
	}

	binaries := []*files.KubeBinary{binary}
	binariesMap := make(map[string]*files.KubeBinary)
	for _, binary := range binaries {
		if err := binary.CreateBaseDir(); err != nil {
			return errors.Wrapf(errors.WithStack(err), "create file %s base dir failed", binary.FileName)
		}

		binariesMap[binary.ID] = binary
		var exists = util.IsExist(binary.Path())
		if exists {
			p := binary.Path()
			if err := binary.SHA256Check(); err != nil {
				_ = exec.Command("/bin/sh", "-c", fmt.Sprintf("rm -f %s", p)).Run()
			} else {
				continue
			}
		}

		if !exists || binary.OverWrite {
			logger.Infof("%s downloading %s %s %s ...", common.LocalHost, arch, binary.ID, binary.Version)
			if err := binary.Download(); err != nil {
				return fmt.Errorf("Failed to download %s binary: %s error: %w ", binary.ID, binary.Url, err)
			}
		}
	}

	t.PipelineCache.Set(common.KubeBinaries+"-"+arch, binariesMap)

	return nil
}

// ~ UnMountS3
type UnMountS3 struct {
	common.KubeAction
}

func (t *UnMountS3) Execute(runtime connector.Runtime) error {
	// exp https://terminus-os-us-west-1.s3.us-west-1.amazonaws.com
	// s3  s3://terminus-os-us-west-1

	storageBucket, _ := t.PipelineCache.GetMustString(common.CacheStorageBucket)
	if storageBucket == "" {
		return nil
	}

	storageAccessKey, _ := t.PipelineCache.GetMustString(common.CacheSTSAccessKey)
	storageSecretKey, _ := t.PipelineCache.GetMustString(common.CacheSTSSecretKey)
	storageToken, _ := t.PipelineCache.GetMustString(common.CacheSTSToken)
	storageClusterId, _ := t.PipelineCache.GetMustString(common.CacheSTSClusterId)

	_, a, f := strings.Cut(storageBucket, "://")
	if !f {
		logger.Errorf("get s3 bucket failed %s", storageBucket)
		return nil
	}
	sa := strings.Split(a, ".")
	if len(sa) < 2 {
		logger.Errorf("get s3 bucket failed %s", storageBucket)
		return nil
	}
	endpoint := fmt.Sprintf("s3://%s", sa[0])
	var cmd = fmt.Sprintf("AWS_ACCESS_KEY_ID=%s AWS_SECRET_ACCESS_KEY=%s AWS_SESSION_TOKEN=%s /usr/local/bin/aws s3 rm %s/%s --recursive",
		storageAccessKey, storageSecretKey, storageToken, endpoint, storageClusterId,
	)

	if _, err := runtime.GetRunner().SudoCmdExt(cmd, false, true); err != nil {
		logger.Errorf("failed to unmount s3 bucket %s: %v", storageBucket, err)
	}

	return nil
}

// ~ UnMountOSS
type UnMountOSS struct {
	common.KubeAction
}

func (t *UnMountOSS) Execute(runtime connector.Runtime) error {
	storageBucket, _ := t.PipelineCache.GetMustString(common.CacheStorageBucket)
	if storageBucket == "" {
		return nil
	}
	storageAccessKey, _ := t.PipelineCache.GetMustString(common.CacheSTSAccessKey)
	storageSecretKey, _ := t.PipelineCache.GetMustString(common.CacheSTSSecretKey)
	storageToken, _ := t.PipelineCache.GetMustString(common.CacheSTSToken)
	storageClusterId, _ := t.PipelineCache.GetMustString(common.CacheSTSClusterId)

	// exp: https://name.area.aliyuncs.com
	// oss  oss://name
	// endpoint: https://area.aliyuncs.com

	b, a, f := strings.Cut(storageBucket, "://")
	if !f {
		logger.Errorf("get oss bucket failed %s", storageBucket)
		return nil
	}

	s := strings.Split(a, ".")
	if len(s) != 4 {
		logger.Errorf("get oss bucket failed %s", storageBucket)
		return nil
	}
	ossName := fmt.Sprintf("oss://%s", s[0])
	ossEndpoint := fmt.Sprintf("%s://%s.%s.%s", b, s[1], s[2], s[3])

	var cmd = fmt.Sprintf("/usr/local/sbin/ossutil64 rm %s/%s/ --endpoint=%s --access-key-id=%s --access-key-secret=%s --sts-token=%s -r -f", ossName, storageClusterId, ossEndpoint, storageAccessKey, storageSecretKey, storageToken)

	if _, err := runtime.GetRunner().SudoCmdExt(cmd, false, true); err != nil {
		logger.Errorf("failed to unmount oss bucket %s: %v", storageBucket, err)
	}

	return nil
}

// ~ StopJuiceFS
type StopJuiceFS struct {
	common.KubeAction
}

func (t *StopJuiceFS) Execute(runtime connector.Runtime) error {
	_, _ = runtime.GetRunner().SudoCmdExt("systemctl stop juicefs; systemctl disable juicefs", false, false)

	_, _ = runtime.GetRunner().SudoCmdExt("rm -rf /var/jfsCache /terminus/jfscache", false, false)

	return nil
}

// ~ StopMinio
type StopMinio struct {
	common.KubeAction
}

func (t *StopMinio) Execute(runtime connector.Runtime) error {
	_, _ = runtime.GetRunner().SudoCmdExt("systemctl stop minio; systemctl disable minio", false, false)
	return nil
}

// ~ StopMinioOperator
type StopMinioOperator struct {
	common.KubeAction
}

func (t *StopMinioOperator) Execute(runtime connector.Runtime) error {
	var cmd = "systemctl stop minio-operator; systemctl disable minio-operator"
	_, _ = runtime.GetRunner().SudoCmdExt(cmd, false, false)
	return nil
}

// ~ StopRedis
type StopRedis struct {
	common.KubeAction
}

func (t *StopRedis) Execute(runtime connector.Runtime) error {
	var cmd = "systemctl stop redis-server; systemctl disable redis-server"
	_, _ = runtime.GetRunner().SudoCmdExt(cmd, false, false)
	_, _ = runtime.GetRunner().SudoCmd("killall -9 redis-server", false, true)
	_, _ = runtime.GetRunner().SudoCmd("unlink /usr/bin/redis-server; unlink /usr/bin/redis-cli", false, true)

	return nil
}

// ~ RemoveTerminusFiles
type RemoveTerminusFiles struct {
	common.KubeAction
}

func (t *RemoveTerminusFiles) Execute(runtime connector.Runtime) error {
	var files = []string{
		"/usr/local/bin/redis-*",
		"/usr/bin/redis-*",
		"/sbin/mount.juicefs",
		"/etc/init.d/redis-server",
		"/usr/local/bin/juicefs",
		"/usr/local/bin/minio",
		"/usr/local/bin/velero",
		"/etc/systemd/system/redis-server.service",
		"/etc/systemd/system/minio.service",
		"/etc/systemd/system/minio-operator.service",
		"/etc/systemd/system/juicefs.service",
		"/etc/systemd/system/containerd.service",
		"/terminus/",
	}

	for _, f := range files {
		runtime.GetRunner().SudoCmdExt(fmt.Sprintf("rm -rf %s", f), false, true)
	}

	return nil
}
