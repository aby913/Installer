package daemon

import (
	"fmt"
	"path/filepath"

	"bytetrade.io/web3os/installer/pkg/common"
	"bytetrade.io/web3os/installer/pkg/core/action"
	"bytetrade.io/web3os/installer/pkg/core/connector"
	"bytetrade.io/web3os/installer/pkg/core/util"
	"bytetrade.io/web3os/installer/pkg/daemon/templates"
	"bytetrade.io/web3os/installer/pkg/manifest"
	"bytetrade.io/web3os/installer/pkg/utils"
	"github.com/pkg/errors"
)

type InstallTerminusdBinary struct {
	common.KubeAction
	manifest.ManifestAction
}

func (g *InstallTerminusdBinary) Execute(runtime connector.Runtime) error {
	if err := utils.ResetTmpDir(runtime); err != nil {
		return err
	}

	binary, err := g.Manifest.Get("terminusd")
	if err != nil {
		return fmt.Errorf("get kube binary terminusd info failed: %w", err)
	}

	path := binary.FilePath(g.BaseDir)

	dst := filepath.Join(common.TmpDir, binary.Filename)
	if err := runtime.GetRunner().Scp(path, dst); err != nil {
		return errors.Wrap(errors.WithStack(err), "sync terminusd tar.gz failed")
	}

	installCmd := fmt.Sprintf("tar -zxf %s && cp -f terminusd /usr/local/bin/ && chmod +x /usr/local/bin/terminusd && rm -rf terminusd*", dst)
	if _, err := runtime.GetRunner().SudoCmd(installCmd, false, false); err != nil {
		return errors.Wrap(errors.WithStack(err), "install terminusd binaries failed")
	}
	return nil
}

type GenerateTerminusdServiceEnv struct {
	common.KubeAction
}

func (g *GenerateTerminusdServiceEnv) Execute(runtime connector.Runtime) error {
	templateAction := action.Template{
		Name:     "TerminusdServiceEnv",
		Template: templates.TerminusdEnv,
		Dst:      filepath.Join("/etc/systemd/system/", templates.TerminusdEnv.Name()),
		Data: util.Data{
			"Version":  g.KubeConf.Arg.TerminusVersion,
			"KubeType": g.KubeConf.Arg.Kubetype,
			"BaseDir":  g.KubeConf.Arg.BaseDir,
		},
	}

	templateAction.Init(nil, nil)
	if err := templateAction.Execute(runtime); err != nil {
		return err
	}
	return nil
}

type GenerateTerminusdService struct {
	common.KubeAction
}

func (g *GenerateTerminusdService) Execute(runtime connector.Runtime) error {
	templateAction := action.Template{
		Name:     "TerminusdService",
		Template: templates.TerminusdService,
		Dst:      filepath.Join("/etc/systemd/system/", templates.TerminusdService.Name()),
		Data:     util.Data{},
	}

	templateAction.Init(nil, nil)
	if err := templateAction.Execute(runtime); err != nil {
		return err
	}
	return nil
}

type EnableTerminusdService struct {
	common.KubeAction
}

func (e *EnableTerminusdService) Execute(runtime connector.Runtime) error {
	if _, err := runtime.GetRunner().SudoCmd("systemctl enable --now terminusd",
		false, false); err != nil {
		return errors.Wrap(errors.WithStack(err), "enable terminusd failed")
	}
	return nil
}
