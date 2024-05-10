package common

import (
	"fmt"
	"os"
	"path"

	"github.com/prometheus/procfs"
	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/config"
)

func IsRunningInContainer() bool {
	// See: https://docs.podman.io/en/latest/markdown/podman-run.1.html
	if _, err := os.Stat("/run/.containerenv"); err == nil {
		return true
	}
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

func GetHostDataHome() (string, error) {
	if IsRunningInContainer() {
		mounts, err := procfs.GetProcMounts(1)
		if err != nil {
			return "", fmt.Errorf("error getting mount points: %v", err)
		}
		for _, mount := range mounts {
			if mount.MountPoint == "/output" {
				return mount.Root, nil
			}
		}
		return "", fmt.Errorf("unable to determine host data home directory")
	} else {
		return config.GetDataHome(), nil
	}
}

func GetHostSiteHome(site v1alpha1.Site) (string, error) {
	dataHome, err := GetHostDataHome()
	if err != nil {
		return "", err
	}
	return path.Join(dataHome, "sites", site.Name), nil
}
