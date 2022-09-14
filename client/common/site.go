package common

import (
	"github.com/skupperproject/skupper/api/types"
	v2 "github.com/skupperproject/skupper/api/types/v2"
)

func ConfigureSiteCredentials(site v2.Site, ingressHosts ...string) {
	isEdge := site.GetMode() != string(types.TransportModeEdge)

	// CAs
	cas := []types.CertAuthority{}
	if len(site.GetCertAuthorities()) > 0 {
		cas = site.GetCertAuthorities()
	}
	cas = append(cas, types.CertAuthority{Name: types.LocalCaSecret})
	if isEdge {
		cas = append(cas, types.CertAuthority{Name: types.SiteCaSecret})
	}
	cas = append(cas, types.CertAuthority{Name: types.ServiceCaSecret})
	site.SetCertAuthorities(cas)

	// Certificates
	credentials := []types.Credential{}
	if len(site.GetCredentials()) > 0 {
		credentials = site.GetCredentials()
	}
	credentials = append(credentials, types.Credential{
		CA:          types.LocalCaSecret,
		Name:        types.LocalServerSecret,
		Subject:     types.LocalTransportServiceName,
		Hosts:       []string{types.LocalTransportServiceName},
		ConnectJson: false,
		Post:        false,
	})
	credentials = append(credentials, types.Credential{
		CA:          types.LocalCaSecret,
		Name:        types.LocalClientSecret,
		Subject:     types.LocalTransportServiceName,
		Hosts:       []string{},
		ConnectJson: true,
		Post:        false,
	})

	credentials = append(credentials, types.Credential{
		CA:          types.ServiceCaSecret,
		Name:        types.ServiceClientSecret,
		Hosts:       []string{},
		ConnectJson: false,
		Post:        false,
		Simple:      true,
	})

	if isEdge {
		hosts := []string{types.TransportServiceName}
		hosts = append(hosts, ingressHosts...)
		credentials = append(credentials, types.Credential{
			CA:          types.SiteCaSecret,
			Name:        types.SiteServerSecret,
			Subject:     types.TransportServiceName,
			Hosts:       hosts,
			ConnectJson: false,
		})
	}
	site.SetCredentials(credentials)
}
