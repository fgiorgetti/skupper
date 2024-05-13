package common

import (
	"fmt"
	"net"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
	"github.com/skupperproject/skupper/pkg/utils"
)

var (
	validLinkAccessRoles = []string{"edge", "inter-router"}
)

type SiteStateValidator struct {
}

// Validate provides a common validation for non-kubernetes sites
// which do not benefit from the Kubernetes API. The goal is not
// to validate the site state against the spec (CRD), but to a more
// basic level, like to ensure that mandatory fields for each resource are
// populated and users will be able to operate the non-k8s site.
func (s *SiteStateValidator) Validate(siteState apis.SiteState) error {
	var err error
	if err = s.validateSite(siteState.Site); err != nil {
		return err
	}
	if err = s.validateLinkAccesses(siteState.LinkAccesses); err != nil {
		return err
	}
	if err = s.validateLinks(siteState); err != nil {
		return err
	}
	if err = s.validateClaims(siteState.Claims); err != nil {
		return err
	}
	if err = s.validateGrants(siteState.Grants); err != nil {
		return err
	}
	if err = s.validateListeners(siteState.Listeners); err != nil {
		return err
	}
	if err = s.validateConnectors(siteState.Connectors); err != nil {
		return err
	}

	return nil
}

func (s *SiteStateValidator) validateSite(site v1alpha1.Site) error {
	if site.Name == "" {
		return fmt.Errorf("invalid site name: empty string")
	}
	return nil
}

func (s *SiteStateValidator) validateLinkAccesses(linkAccesses map[string]v1alpha1.LinkAccess) error {
	for _, linkAccess := range linkAccesses {
		if linkAccess.Name == "" {
			return fmt.Errorf("invalid link access name: empty string")
		}
		if len(linkAccess.Spec.Roles) == 0 {
			return fmt.Errorf("invalid link access: roles are required")
		}
		for _, role := range linkAccess.Spec.Roles {
			if !utils.StringSliceContains(validLinkAccessRoles, role.Role) {
				return fmt.Errorf("invalid link access: %s - invalid role: %s (valid roles: %s)",
					linkAccess.Name, role.Role, validLinkAccessRoles)
			}
		}
	}
	return nil
}

func (s *SiteStateValidator) validateLinks(siteState apis.SiteState) error {
	for linkName, link := range siteState.Links {
		if link.Name == "" {
			return fmt.Errorf("invalid link name: empty string")
		}
		secretName := link.Spec.TlsCredentials
		if _, ok := siteState.Secrets[secretName]; !ok {
			return fmt.Errorf("invalid link %q - secret %s not found", linkName, secretName)
		}
	}
	return nil
}

func (s *SiteStateValidator) validateClaims(claims map[string]v1alpha1.Claim) error {
	for _, claim := range claims {
		if claim.Name == "" {
			return fmt.Errorf("invalid claim name: empty string")
		}
	}
	return nil
}

func (s *SiteStateValidator) validateGrants(grants map[string]v1alpha1.Grant) error {
	for _, grant := range grants {
		if grant.Name == "" {
			return fmt.Errorf("invalid grant name: empty string")
		}
	}
	return nil
}

func (s *SiteStateValidator) validateListeners(listeners map[string]v1alpha1.Listener) error {
	hostPorts := map[string][]int{}
	for name, listener := range listeners {
		if listener.Name == "" {
			return fmt.Errorf("invalid listener name: empty string")
		}
		if listener.Spec.Host == "" || listener.Spec.Port == 0 {
			return fmt.Errorf("host and port and required")
		}
		// TODO allow host field to expose a name (not an ip)
		//      this is related to iptables/proxy container and will also
		//      require a container network to be provided
		if ip := net.ParseIP(listener.Spec.Host); ip == nil {
			return fmt.Errorf("invalid listener host: %s - a valid IP address is expected", listener.Spec.Host)
		}

		if utils.IntSliceContains(hostPorts[listener.Spec.Host], listener.Spec.Port) {
			return fmt.Errorf("port %d is already mapped for host %q (listener: %q)", listener.Spec.Port, listener.Spec.Host, name)
		}
		hostPorts[listener.Spec.Host] = append(hostPorts[listener.Spec.Host], listener.Spec.Port)
	}
	return nil
}

func (s *SiteStateValidator) validateConnectors(connectors map[string]v1alpha1.Connector) error {
	hostPorts := map[string][]int{}
	for name, connector := range connectors {
		if connector.Name == "" {
			return fmt.Errorf("invalid connector name: empty string")
		}
		if connector.Spec.Host == "" || connector.Spec.Port == 0 {
			return fmt.Errorf("connector host and port are required")
		}
		// TODO allow host field to expose a name (not an ip)
		//      this is related to iptables/proxy container and will also
		//      require a container network to be provided
		if ip := net.ParseIP(connector.Spec.Host); ip == nil {
			return fmt.Errorf("invalid connector host: %s - a valid IP address is expected", connector.Spec.Host)
		}
		if utils.IntSliceContains(hostPorts[connector.Spec.Host], connector.Spec.Port) {
			return fmt.Errorf("port %d is already mapped for host %q (listener: %q)", connector.Spec.Port, connector.Spec.Host, name)
		}
		hostPorts[connector.Spec.Host] = append(hostPorts[connector.Spec.Host], connector.Spec.Port)
	}
	return nil
}
