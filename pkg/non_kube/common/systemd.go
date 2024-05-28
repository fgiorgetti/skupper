package common

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"text/template"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
)

var (
	//go:embed systemd_service.template
	SystemdServiceTemplate string
)

type systemdServiceInfo struct {
	Site           v1alpha1.Site
	SiteScriptPath string
	RuntimeDir     string
}

func NewSystemdServiceInfo(site v1alpha1.Site) (*systemdServiceInfo, error) {
	siteHomeDir, err := apis.GetHostSiteHome(site)
	if err != nil {
		return nil, err
	}
	siteScriptPath := path.Join(siteHomeDir, RuntimeScriptsPath)
	return &systemdServiceInfo{
		Site:           site,
		SiteScriptPath: siteScriptPath,
		RuntimeDir:     apis.GetRuntimeDir(),
	}, nil
}

func (s *systemdServiceInfo) GetServiceName() string {
	return fmt.Sprintf("skupper-site-%s.service", s.Site.Name)
}

func (s *systemdServiceInfo) Create() error {
	if !apis.IsRunningInContainer() && !IsSystemdEnabled() {
		msg := "SystemD is not enabled"
		if os.Getuid() != 0 {
			msg += " at user level"
		}
		return fmt.Errorf(msg)
	}

	var buf = new(bytes.Buffer)
	service := template.Must(template.New(s.GetServiceName()).Parse(SystemdServiceTemplate))
	err := service.Execute(buf, s)
	if err != nil {
		return err
	}

	// Creating the base dir
	baseDir := filepath.Dir(s.getServiceFile())
	if _, err := os.Stat(baseDir); err != nil {
		if err = os.MkdirAll(baseDir, 0755); err != nil {
			return fmt.Errorf("unable to create base directory %s - %q", baseDir, err)
		}
	}

	// Saving systemd user service
	serviceName := s.GetServiceName()
	err = os.WriteFile(s.getServiceFile(), buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("unable to write unit file (%s): %w", s.getServiceFile(), err)
	}

	// Only enable when running locally
	if !apis.IsRunningInContainer() {
		return s.enableService(serviceName)
	}

	return nil
}

func (s *systemdServiceInfo) getServiceFile() string {
	if apis.IsRunningInContainer() {
		return path.Join(GetDefaultOutputPath(s.Site.Name), RuntimeScriptsPath, s.GetServiceName())
	}
	if os.Getuid() == 0 {
		return path.Join("/etc/systemd/system", s.GetServiceName())
	}
	return path.Join(apis.GetConfigHome(), "systemd/user", s.GetServiceName())
}

func (s *systemdServiceInfo) Remove() error {
	if !apis.IsRunningInContainer() && !IsSystemdEnabled() {
		return fmt.Errorf("SystemD is not enabled at user level")
	}

	// Stopping systemd user service
	if !apis.IsRunningInContainer() {
		cmd := GetCmdStopSystemdService(s.GetServiceName())
		_ = cmd.Run()

		// Disabling systemd user service
		cmd = GetCmdDisableSystemdService(s.GetServiceName())
		_ = cmd.Run()
	}

	// Removing the .service file
	_ = os.Remove(s.getServiceFile())

	// Reloading systemd user daemon
	if !apis.IsRunningInContainer() {
		cmd := GetCmdReloadSystemdDaemon()
		_ = cmd.Run()

		// Resetting failed status
		cmd = GetCmdResetFailedSystemService(s.GetServiceName())
		_ = cmd.Run()
	}

	return nil
}

func (s *systemdServiceInfo) enableService(serviceName string) error {
	// Enabling systemd user service
	cmd := GetCmdEnableSystemdService(serviceName)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("unable to enable service (%s): %w", s.getServiceFile(), err)
	}

	// Reloading systemd user daemon
	cmd = GetCmdReloadSystemdDaemon()
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Unable to user service daemon-reload: %w", err)
	}

	// Starting systemd user service
	cmd = GetCmdStartSystemdService(serviceName)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Unable to start user service: %w", err)
	}

	return nil
}

func GetCmdEnableSystemdService(serviceName string) *exec.Cmd {
	if os.Getuid() == 0 {
		return exec.Command("systemctl", "enable", serviceName)
	}
	return exec.Command("systemctl", "--user", "enable", serviceName)
}

func GetCmdDisableSystemdService(serviceName string) *exec.Cmd {
	if os.Getuid() == 0 {
		return exec.Command("systemctl", "disable", serviceName)
	}
	return exec.Command("systemctl", "--user", "disable", serviceName)
}

func GetCmdReloadSystemdDaemon() *exec.Cmd {
	if os.Getuid() == 0 {
		return exec.Command("systemctl", "daemon-reload")
	}
	return exec.Command("systemctl", "--user", "daemon-reload")
}

func GetCmdStartSystemdService(serviceName string) *exec.Cmd {
	if os.Getuid() == 0 {
		return exec.Command("systemctl", "start", serviceName)
	}
	return exec.Command("systemctl", "--user", "start", serviceName)
}

func GetCmdStopSystemdService(serviceName string) *exec.Cmd {
	if os.Getuid() == 0 {
		return exec.Command("systemctl", "stop", serviceName)
	}
	return exec.Command("systemctl", "--user", "stop", serviceName)
}

func GetCmdResetFailedSystemService(serviceName string) *exec.Cmd {
	if os.Getuid() == 0 {
		return exec.Command("systemctl", "reset-failed", serviceName)
	}
	return exec.Command("systemctl", "--user", "reset-failed", serviceName)
}

func GetCmdIsSystemdEnabled() *exec.Cmd {
	if os.Getuid() == 0 {
		return exec.Command("systemctl", []string{"list-units", "--no-pager"}...)
	}
	return exec.Command("systemctl", []string{"--user", "list-units", "--no-pager"}...)
}

func IsSystemdEnabled() bool {
	cmd := GetCmdIsSystemdEnabled()
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func IsLingeringEnabled(user string) bool {
	lingerFile := fmt.Sprintf("/var/lib/systemd/linger/%s", user)
	_, err := os.Stat(lingerFile)
	return err == nil
}
