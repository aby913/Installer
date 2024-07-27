package plugins

import (
	"fmt"
	"path"

	"bytetrade.io/web3os/installer/pkg/common"
	cc "bytetrade.io/web3os/installer/pkg/core/common"
	"bytetrade.io/web3os/installer/pkg/core/connector"
	"bytetrade.io/web3os/installer/pkg/core/prepare"
	"bytetrade.io/web3os/installer/pkg/core/task"
)

// ~ InstallMonitorDashboardCrd
type InstallMonitorDashboardCrd struct {
	common.KubeAction
}

func (t *InstallMonitorDashboardCrd) Execute(runtime connector.Runtime) error {
	var p = path.Join(runtime.GetFilesDir(), cc.BuildDir, "ks-monitor", "monitoring-dashboard")
	var cmd = fmt.Sprintf("/usr/local/bin/kubectl apply -f %s", p)
	if _, err := runtime.GetRunner().SudoCmd(cmd, false, true); err != nil {
		return err
	}
	return nil
}

// +

// ~ CreateMonitorDashboardModule
type CreateMonitorDashboardModule struct {
	common.KubeModule
}

func (m *CreateMonitorDashboardModule) Init() {
	m.Name = "CreateMonitorDashboardModule"

	installMonitorDashboardCrd := &task.RemoteTask{
		Name:  "InstallMonitorDashboardCrd",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
			new(NotEqualDesiredVersion),
		},
		Action:   new(InstallMonitorDashboardCrd),
		Parallel: false,
		Retry:    0,
	}

	m.Tasks = []task.Interface{
		installMonitorDashboardCrd,
	}

}
