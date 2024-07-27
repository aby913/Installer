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

package dns

import (
	"fmt"
	"path/filepath"
	"strings"

	"bytetrade.io/web3os/installer/pkg/common"
	"bytetrade.io/web3os/installer/pkg/core/action"
	"bytetrade.io/web3os/installer/pkg/core/connector"
	"bytetrade.io/web3os/installer/pkg/core/logger"
	"bytetrade.io/web3os/installer/pkg/core/util"
	"bytetrade.io/web3os/installer/pkg/plugins/dns/templates"
	"bytetrade.io/web3os/installer/pkg/utils"
	"github.com/pkg/errors"
)

// ~ SetProxyNameServer
type SetProxyNameServer struct {
	common.KubeAction
}

func (s *SetProxyNameServer) Execute(runtime connector.Runtime) error {
	proxy, ok := s.PipelineCache.Get(common.CacheProxy)
	if !ok || proxy == nil {
		return nil
	}
	if addr := proxy.(string); len(addr) != 0 {
		if !utils.IsValidIP(addr) {
			// todo set nameserver
			return nil
		}

		if _, err := runtime.GetRunner().SudoCmd("cat /etc/resolv.conf > /etc/resolv.conf.bak", false, false); err != nil {
			logger.Errorf("backup /etc/resolv.conf failed: %v", err)
		}
		if _, err := runtime.GetRunner().SudoCmd(fmt.Sprintf("echo nameserver %s > /etc/resolv.conf", addr), false, true); err != nil {
			logger.Errorf("set nameserver %s failed: %v", addr, err)
		}
	}
	return nil
}

// ~ OverrideCoreDNS
type OverrideCoreDNS struct {
	common.KubeAction
}

func (o *OverrideCoreDNS) Execute(runtime connector.Runtime) error {
	if _, err := runtime.GetRunner().SudoCmd("/usr/local/bin/kubectl delete -n kube-system svc kube-dns", false, true); err != nil {
		if !strings.Contains(err.Error(), "NotFound") {
			return errors.Wrap(errors.WithStack(err), "delete kube-dns failed")
		}
	}

	if _, err := runtime.GetRunner().SudoCmd("/usr/local/bin/kubectl apply -f /etc/kubernetes/coredns-svc.yaml", false, true); err != nil {
		return errors.Wrap(errors.WithStack(err), "create coredns service failed")
	}
	return nil
}

// ~ DeployNodeLocalDNS
type DeployNodeLocalDNS struct {
	common.KubeAction
}

func (d *DeployNodeLocalDNS) Execute(runtime connector.Runtime) error {
	if _, err := runtime.GetRunner().SudoCmdExt("/usr/local/bin/kubectl apply -f /etc/kubernetes/nodelocaldns.yaml", false, false); err != nil {
		return errors.Wrap(errors.WithStack(err), "deploy nodelocaldns failed")
	}
	return nil
}

// ~ GenerateNodeLocalDNSConfigMap
type GenerateNodeLocalDNSConfigMap struct {
	common.KubeAction
}

func (g *GenerateNodeLocalDNSConfigMap) Execute(runtime connector.Runtime) error {
	clusterIP, err := runtime.GetRunner().SudoCmd("/usr/local/bin/kubectl get svc -n kube-system coredns -o jsonpath='{.spec.clusterIP}'", false, false)
	if err != nil {
		return errors.Wrap(errors.WithStack(err), "get clusterIP failed")
	}

	if len(clusterIP) == 0 {
		clusterIP = g.KubeConf.Cluster.CorednsClusterIP()
	}

	templateAction := action.Template{
		Name:     "GenerateNodeLocalDNSConfigMap",
		Template: templates.NodeLocalDNSConfigMap,
		Dst:      filepath.Join(common.KubeConfigDir, templates.NodeLocalDNSConfigMap.Name()),
		Data: util.Data{
			"ForwardTarget": clusterIP,
			"DNSDomain":     g.KubeConf.Cluster.Kubernetes.DNSDomain,
		},
	}

	templateAction.Init(nil, nil)
	if err := templateAction.Execute(runtime); err != nil {
		return err
	}
	return nil
}

// ~ ApplyNodeLocalDNSConfigMap
type ApplyNodeLocalDNSConfigMap struct {
	common.KubeAction
}

func (a *ApplyNodeLocalDNSConfigMap) Execute(runtime connector.Runtime) error {
	if _, err := runtime.GetRunner().SudoCmdExt("/usr/local/bin/kubectl apply -f /etc/kubernetes/nodelocaldnsConfigmap.yaml", false, false); err != nil {
		return errors.Wrap(errors.WithStack(err), "apply nodelocaldns configmap failed")
	}
	return nil
}
