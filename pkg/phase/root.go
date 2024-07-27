package phase

import (
	"fmt"
	"os"
	"strings"

	"bytetrade.io/web3os/installer/pkg/common"
	"bytetrade.io/web3os/installer/pkg/constants"
	"bytetrade.io/web3os/installer/pkg/core/logger"
	"bytetrade.io/web3os/installer/pkg/core/util"
)

func GetCurrentKubeVersion() string {
	stdout, _, err := util.Exec("/usr/local/bin/kubectl get nodes -o jsonpath='{.items[0].status.nodeInfo.kubeletVersion}'", false, false)
	if err != nil {
		goto SKIP
	}
	if stdout != "" {
		if strings.Contains(stdout, "+k3s1") {
			stdout = strings.ReplaceAll(stdout, "+k3s1", "-k3s")
		} else if strings.Contains(stdout, "+k3s2") {
			stdout = strings.ReplaceAll(stdout, "+k3s2", "-k3s")
		}
	}

	constants.InstalledKubeVersion = stdout
	goto SKIP

SKIP:
	if constants.InstalledKubeVersion != "" {
		fmt.Printf("KUBE: version: %s\n", constants.InstalledKubeVersion)
	}

	return constants.InstalledKubeVersion
}

func UserParameters() *common.User {
	var u = &common.User{
		UserName:   "zhaoyu",
		Password:   "",
		Email:      "zhaoyu@bytetrade.io",
		DomainName: "myterminus.com",
	}

	return u
}

func StorageParameters() *common.Storage {
	var storageAccessKey, storageSecretKey, storageToken, storageClusterId string

	if stdout, _, err := util.Exec("/usr/local/bin/kubectl get terminus terminus -o jsonpath='{.metadata.annotations.bytetrade\\.io/s3-ak}'", false, false); stdout == "" {
		storageAccessKey = os.Getenv(common.EnvStorageAccessKeyName)
		if storageAccessKey == "" {
			logger.Errorf("storage access key not found")
		}
	} else if err == nil {
		storageAccessKey = stdout
	}

	if stdout, _, err := util.Exec("/usr/local/bin/kubectl get terminus terminus -o jsonpath='{.metadata.annotations.bytetrade\\.io/s3-sk}'", false, false); stdout == "" {
		storageSecretKey = os.Getenv(common.EnvStorageSecretKeyName)
		if storageSecretKey == "" {
			logger.Errorf("storage secret key not found")
		}
	} else if err == nil {
		storageSecretKey = stdout
	}

	if stdout, _, err := util.Exec("/usr/local/bin/kubectl get terminus terminus -o jsonpath='{.metadata.annotations.bytetrade\\.io/s3-sts}'", false, false); stdout == "" {
		storageToken = os.Getenv(common.EnvStorageTokenName)
		if storageToken == "" {
			logger.Errorf("storage token not found")
		}
	} else if err == nil {
		storageToken = stdout
	}

	if stdout, _, err := util.Exec("/usr/local/bin/kubectl get terminus terminus -o jsonpath='{.metadata.annotations.bytetrade\\.io/cluster-id}'", false, false); stdout == "" {
		storageClusterId = os.Getenv(common.EnvStorageClusterIdName)
		if storageClusterId == "" {
			logger.Errorf("storage cluster id not found")
		}
	} else if err == nil {
		storageClusterId = stdout
	}

	storageVendor := os.Getenv("TERMINUS_IS_CLOUD_VERSION")
	storageType := os.Getenv("STORAGE")
	storageBucket := os.Getenv("S3_BUCKET")

	if storageType == "" {
		storageType = common.Minio
	}

	return &common.Storage{
		StorageAccessKey: storageAccessKey,
		StorageSecretKey: storageSecretKey,
		StorageToken:     storageToken,
		StorageClusterId: storageClusterId,
		StorageVendor:    storageVendor,
		StorageType:      storageType,
		StorageBucket:    storageBucket,
	}
}
