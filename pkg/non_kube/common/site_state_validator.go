package common

import (
	"fmt"
	"net"
	"regexp"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
	"github.com/skupperproject/skupper/pkg/utils"
)

var (
	validLinkAccessRoles = []string{"edge", "inter-router"}
	rfc1123Regex         = regexp.MustCompile("^[a-z0-9]([-a-z0-9]*[a-z0-9])?$")
)

const (
	rfc1123Error = `a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')`
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
	if err := ValidateName(site.Name); err != nil {
		return fmt.Errorf("invalid site name: %w", err)
	}
	return nil
}

func (s *SiteStateValidator) validateLinkAccesses(linkAccesses map[string]v1alpha1.LinkAccess) error {
	for _, linkAccess := range linkAccesses {
		if err := ValidateName(linkAccess.Name); err != nil {
			return fmt.Errorf("invalid link access name: %w", err)
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
		if err := ValidateName(link.Name); err != nil {
			return fmt.Errorf("invalid link name: %w", err)
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
		if err := ValidateName(claim.Name); err != nil {
			return fmt.Errorf("invalid claim name: %w", err)
		}
	}
	return nil
}

func (s *SiteStateValidator) validateGrants(grants map[string]v1alpha1.Grant) error {
	for _, grant := range grants {
		if err := ValidateName(grant.Name); err != nil {
			return fmt.Errorf("invalid grant name: %w", err)
		}
	}
	return nil
}

func (s *SiteStateValidator) validateListeners(listeners map[string]v1alpha1.Listener) error {
	hostPorts := map[string][]int{}
	for name, listener := range listeners {
		if err := ValidateName(listener.Name); err != nil {
			return fmt.Errorf("invalid listener name: %w", err)
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
		if err := ValidateName(connector.Name); err != nil {
			return fmt.Errorf("invalid connector name: %w", err)
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

func ValidateName(name string) error {
	if !rfc1123Regex.MatchString(name) {
		return fmt.Errorf("invalid name %q: %s", name, rfc1123Error)
	}
	return nil
}
