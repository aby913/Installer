/*
 Copyright 2021 The KubeSphere Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package common

import (
	kubekeyapiv1alpha2 "bytetrade.io/web3os/installer/apis/kubekey/v1alpha2"
	kubekeyclientset "bytetrade.io/web3os/installer/clients/clientset/versioned"
	"bytetrade.io/web3os/installer/pkg/core/connector"
	"bytetrade.io/web3os/installer/pkg/core/storage"
)

type KubeRuntime struct {
	connector.BaseRuntime
	ClusterName string
	Cluster     *kubekeyapiv1alpha2.ClusterSpec
	Kubeconfig  string
	ClientSet   *kubekeyclientset.Clientset
	Arg         Argument
}

type Argument struct {
	NodeName            string
	FilePath            string
	KubernetesVersion   string
	KsEnable            bool
	KsVersion           string
	Debug               bool
	IgnoreErr           bool
	SkipPullImages      bool
	SKipPushImages      bool
	SecurityEnhancement bool
	DeployLocalStorage  *bool
	// DownloadCommand     func(path, url string) string
	SkipConfirmCheck bool
	InCluster        bool
	ContainerManager string
	FromCluster      bool
	KubeConfig       string
	Artifact         string
	InstallPackages  bool
	ImagesDir        string
	Namespace        string
	DeleteCRI        bool
	Role             string
	Type             string

	// Extra args
	ExtraAddon string // addon yaml config

	// Registry mirrors
	RegistryMirrors string
	Proxy           string

	// master node ssh config
	MasterHost              string
	MasterNodeName          string
	MasterSSHPort           int
	MasterSSHUser           string
	MasterSSHPassword       string
	MasterSSHPrivateKeyPath string
	LocalSSHPort            int

	SkipMasterPullImages bool

	// db
	Provider storage.Provider
	// User
	User *User
	// storage
	Storage *Storage
	AWS     *AwsHost
	// request
	Params  map[string]interface{}
	Request any

	Minikube        bool
	MinikubeProfile string
}

type AwsHost struct {
	PublicIp string
	Hostname string
}

type User struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	DomainName string `json:"domain_name"`
}

type Storage struct {
	StorageVendor    string `json:"storage_vendor"`
	StorageType      string `json:"storage_type"`
	StorageAccessKey string `json:"storage_access_key"`
	StorageSecretKey string `json:"storage_secret_key"`
	StorageToken     string `json:"storage_token"`
	StorageClusterId string `json:"storage_cluster_id"`
	StorageBucket    string `json:"storage_bucket"`
}

func NewKubeRuntime(flag string, arg Argument) (*KubeRuntime, error) {
	loader := NewLoader(flag, arg)
	cluster, err := loader.Load()
	if err != nil {
		return nil, err
	}

	if err = loadExtraAddons(cluster, arg.ExtraAddon); err != nil {
		return nil, err
	}

	base := connector.NewBaseRuntime(cluster.Name, connector.NewDialer(), arg.Debug, arg.IgnoreErr, arg.Provider)

	clusterSpec := &cluster.Spec
	defaultCluster, roleGroups := clusterSpec.SetDefaultClusterSpec(arg.InCluster, arg.Minikube)

	hostSet := make(map[string]struct{})
	for _, role := range roleGroups {
		for _, host := range role {
			if host.IsRole(Master) || host.IsRole(Worker) {
				host.SetRole(K8s)
			}
			if host.IsRole(Master) && arg.SkipMasterPullImages {
				host.GetCache().Set(SkipMasterNodePullImages, true)
			}
			if _, ok := hostSet[host.GetName()]; !ok {
				hostSet[host.GetName()] = struct{}{}
				base.AppendHost(host)
				base.AppendRoleMap(host)
			}
			host.SetMinikube(arg.Minikube)
			host.SetMinikubeProfile(arg.MinikubeProfile)
		}
	}

	arg.KsEnable = defaultCluster.KubeSphere.Enabled
	arg.KsVersion = defaultCluster.KubeSphere.Version
	r := &KubeRuntime{
		Cluster:     defaultCluster,
		ClusterName: cluster.Name,
		Arg:         arg,
	}
	r.BaseRuntime = base

	return r, nil
}

// Copy is used to create a copy for Runtime.
func (k *KubeRuntime) Copy() connector.Runtime {
	runtime := *k
	return &runtime
}
