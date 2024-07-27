package storage

import (
	"bytetrade.io/web3os/installer/pkg/common"
	"bytetrade.io/web3os/installer/pkg/core/prepare"
	"bytetrade.io/web3os/installer/pkg/core/task"
)

// ~ InitStorageModule
type InitStorageModule struct {
	common.KubeModule
	Skip bool
}

func (m *InitStorageModule) IsSkip() bool {
	return m.Skip
}

func (m *InitStorageModule) Init() {
	m.Name = "InitStorage"

	mkStorageDir := &task.RemoteTask{
		Name:  "CreateStorageDir",
		Hosts: m.Runtime.GetAllHosts(),
		Prepare: &prepare.PrepareCollection{
			&CheckStorageVendor{},
		},
		Action:   new(MkStorageDir),
		Parallel: false,
		Retry:    0,
	}

	m.Tasks = []task.Interface{
		mkStorageDir,
	}
}

// ~ RemoveMountModule
type RemoveMountModule struct {
	common.KubeModule
}

func (m *RemoveMountModule) Init() {
	// 如果是云上，则需要下载 oss 或 s3 的工具，进行挂载解绑
	m.Name = "RemoveMount"

	downloadStorageCli := &task.RemoteTask{
		Name:  "Download",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(CheckStorageVendor),
		},
		Action:   new(DownloadStorageCli),
		Parallel: false,
		Retry:    0,
	}

	unMountOSS := &task.RemoteTask{
		Name:  "UnMountOSS",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			&CheckStorageType{
				StorageType: "oss",
			},
		},
		Action:   new(UnMountOSS),
		Parallel: false,
		Retry:    0,
	}

	unMountS3 := &task.RemoteTask{
		Name:  "UnMountS3",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			&CheckStorageType{
				StorageType: "s3",
			},
		},
		Action:   new(UnMountS3),
		Parallel: false,
		Retry:    0,
	}

	m.Tasks = []task.Interface{
		downloadStorageCli,
		unMountOSS,
		unMountS3,
	}
}

// ~ RemoveStorage
type RemoveStorageModule struct {
	common.KubeModule
}

func (m *RemoveStorageModule) Init() {
	m.Name = "RemoveStorage"

	stopJuiceFS := &task.RemoteTask{
		Name:  "StopJuiceFS",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
		},
		Action:   new(StopJuiceFS),
		Parallel: false,
		Retry:    0,
	}

	stopMinio := &task.RemoteTask{
		Name:  "StopMinio",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
		},
		Action:   new(StopMinio),
		Parallel: false,
		Retry:    0,
	}

	stopMinioOperator := &task.RemoteTask{
		Name:  "StopMinioOperator",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
		},
		Action:   new(StopMinioOperator),
		Parallel: false,
		Retry:    0,
	}

	stopRedis := &task.RemoteTask{
		Name:  "StopRedis",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
		},
		Action:   new(StopRedis),
		Parallel: false,
		Retry:    0,
	}

	removeTerminusFiles := &task.RemoteTask{
		Name:  "RemoveTerminusFiles",
		Hosts: m.Runtime.GetHostsByRole(common.Master),
		Prepare: &prepare.PrepareCollection{
			new(common.OnlyFirstMaster),
		},
		Action:   new(RemoveTerminusFiles),
		Parallel: false,
		Retry:    0,
	}

	m.Tasks = []task.Interface{
		stopJuiceFS,
		stopMinio,
		stopMinioOperator,
		stopRedis,
		removeTerminusFiles,
	}
}
