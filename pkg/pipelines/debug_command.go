package pipelines

import (
	"bytetrade.io/web3os/installer/pkg/common"
	"bytetrade.io/web3os/installer/pkg/core/module"
	"bytetrade.io/web3os/installer/pkg/core/pipeline"
	"bytetrade.io/web3os/installer/pkg/terminus"
)

func DebugCommand() error {
	args := common.Argument{
		KsEnable:          true,
		KsVersion:         common.DefaultKubeSphereVersion,
		KubernetesVersion: common.DefaultK3sVersion,
		InstallPackages:   false,
		SKipPushImages:    false,
		ContainerManager:  common.Containerd,
	}

	runtime, err := common.NewKubeRuntime(common.AllInOne, args)
	if err != nil {
		return err
	}

	m := []module.Module{
		&terminus.SetupWs{},
	}

	p := pipeline.Pipeline{
		Name:    "Debug Command",
		Modules: m,
		Runtime: runtime,
	}

	return p.Start()
}
