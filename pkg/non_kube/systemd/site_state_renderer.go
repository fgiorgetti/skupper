package systemd

import (
	"fmt"
	"os"
	"path"

	"github.com/skupperproject/skupper/pkg/config"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
	"github.com/skupperproject/skupper/pkg/non_kube/common"
	"github.com/skupperproject/skupper/pkg/utils"
)

type SiteStateRenderer struct {
	loadedSiteState apis.SiteState
	siteState       apis.SiteState
	configRenderer  *common.FileSystemConfigurationRenderer
}

func (s *SiteStateRenderer) Render(loadedSiteState apis.SiteState) error {
	var err error
	var validator apis.SiteStateValidator = &common.SiteStateValidator{}
	// TODO enhance site state validator (too basic yet)
	err = validator.Validate(loadedSiteState)
	if err != nil {
		return err
	}
	s.loadedSiteState = loadedSiteState
	// active (runtime) SiteState
	s.siteState = common.CopySiteState(s.loadedSiteState)
	err = common.RedeemClaims(&s.siteState)
	if err != nil {
		return fmt.Errorf("failed to redeem claims: %v", err)
	}
	err = common.PrepareCertificatesAndLinkAccess(&s.siteState)
	if err != nil {
		return fmt.Errorf("failed to prepare systemd site: %w", err)
	}
	// rendering non-kube configuration files and certificates
	siteHome, err := apis.GetHostSiteHome(s.siteState.Site)
	if err != nil {
		return fmt.Errorf("failed to get site home: %w", err)
	}
	s.configRenderer = &common.FileSystemConfigurationRenderer{
		Force:              false, // TODO discuss how this should be handled?
		SslProfileBasePath: siteHome,
	}
	err = s.configRenderer.Render(s.siteState)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Serializing loaded and runtime site states
	if err = apis.MarshalSiteState(s.loadedSiteState, path.Join(s.configRenderer.OutputPath, common.LoadedSiteStatePath)); err != nil {
		return err
	}
	// No more site state changes after this
	if err = apis.MarshalSiteState(s.siteState, path.Join(s.configRenderer.OutputPath, common.RuntimeSiteStatePath)); err != nil {
		return err
	}
	// Saving runtime platform
	platform := config.GetPlatform()
	content := fmt.Sprintf("platform: %s\n", string(platform))
	err = os.WriteFile(path.Join(s.configRenderer.OutputPath, common.RuntimeSiteStatePath, "platform.yaml"), []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write runtime platform: %w", err)
	}
	// TODO Controller, collector, claims api, console have not yet been planned for systemd sites

	// Create systemd service
	if err = s.createSystemdService(); err != nil {
		return err
	}
	return nil
}

func (s *SiteStateRenderer) createSystemdService() error {
	// Creating systemd user service
	systemd, err := common.NewSystemdServiceInfo(s.siteState.Site)
	if err != nil {
		return err
	}
	if err = systemd.Create(); err != nil {
		return fmt.Errorf("unable to create startup service %q - %v\n", systemd.GetServiceName(), err)
	}

	// Validate if lingering is enabled for current user
	if !apis.IsRunningInContainer() {
		username := utils.ReadUsername()
		if os.Getuid() != 0 && !common.IsLingeringEnabled(username) {
			fmt.Printf("It is recommended to enable lingering for %s, otherwise Skupper may not start on boot.\n", username)
		}
	}

	return nil
}
