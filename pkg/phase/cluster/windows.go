package cluster

import (
	"bytetrade.io/web3os/installer/pkg/common"
	"bytetrade.io/web3os/installer/pkg/core/module"
	"bytetrade.io/web3os/installer/pkg/windows"
)

type windowsInstallPhaseBuilder struct {
	runtime *common.KubeRuntime
}

func (w *windowsInstallPhaseBuilder) build() []module.Module {
	return []module.Module{
		&windows.InstallWSLModule{},
		&windows.InstallWSLUbuntuDistroModule{},
		&windows.ConfigWslModule{},
		&windows.InstallTerminusModule{},
	}
}
