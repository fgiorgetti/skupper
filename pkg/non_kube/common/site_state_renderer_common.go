package common

import (
	"fmt"
	"maps"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
	"github.com/skupperproject/skupper/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CopySiteState(siteState apis.SiteState) apis.SiteState {
	// Preserving loaded state
	var activeSiteState apis.SiteState
	siteState.Site.DeepCopyInto(&activeSiteState.Site)
	activeSiteState.Listeners = maps.Clone(siteState.Listeners)
	activeSiteState.Connectors = maps.Clone(siteState.Connectors)
	activeSiteState.LinkAccesses = maps.Clone(siteState.LinkAccesses)
	activeSiteState.Claims = maps.Clone(siteState.Claims)
	activeSiteState.Links = maps.Clone(siteState.Links)
	activeSiteState.Grants = maps.Clone(siteState.Grants)
	activeSiteState.SecuredAccesses = maps.Clone(siteState.SecuredAccesses)
	activeSiteState.Certificates = maps.Clone(siteState.Certificates)
	activeSiteState.Secrets = maps.Clone(siteState.Secrets)
	return activeSiteState
}

func PrepareCertificatesAndLinkAccess(siteState *apis.SiteState) error {
	/*
		Prepare:
		- normal listener port ($management)
		- generate Certificates
	*/
	// Preparing siteState

	// If normal (role) link access not specified, create one
	err := PrepareNormalListener(siteState)
	if err != nil {
		return err
	}

	// Generating certificates for all link accesses and adaptors
	PrepareLinkAccessesCertificates(siteState)
	PrepareTCPCertificates(siteState)
	return nil
}

func PrepareNormalListener(siteState *apis.SiteState) error {
	normalListenerFound := false
	for _, la := range siteState.LinkAccesses {
		for _, role := range la.Spec.Roles {
			if role.Role == "normal" {
				normalListenerFound = true
			}
		}
	}
	if !normalListenerFound {
		name := fmt.Sprintf("skupper-local")
		port, err := utils.TcpPortNextFree(5671)
		if err != nil {
			return err
		}
		tlsCaName := fmt.Sprintf("%s-ca", name)
		tlsServerName := fmt.Sprintf("%s-server", name)
		tlsClientName := fmt.Sprintf("%s-client", name)
		siteState.LinkAccesses[name] = v1alpha1.LinkAccess{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: v1alpha1.LinkAccessSpec{
				Roles: []v1alpha1.LinkAccessRole{
					{
						Role: "normal",
						Port: port,
					},
				},
				TlsCredentials: tlsServerName,
				Ca:             tlsCaName,
			},
		}
		// TODO Validate if CA Certificate (CR) is properly described
		siteState.Certificates[tlsCaName] = v1alpha1.Certificate{
			ObjectMeta: metav1.ObjectMeta{
				Name: tlsCaName,
			},
			Spec: v1alpha1.CertificateSpec{
				Subject: tlsCaName,
				Hosts:   []string{"127.0.0.1", "localhost"},
				Signing: true,
			},
		}
		// TODO Validate if Server certificate looks good
		siteState.Certificates[tlsServerName] = v1alpha1.Certificate{
			ObjectMeta: metav1.ObjectMeta{
				Name: tlsServerName,
			},
			Spec: v1alpha1.CertificateSpec{
				Subject: "127.0.0.1",
				Hosts:   []string{"127.0.0.1", "localhost"},
				Ca:      tlsCaName,
				Server:  true,
			},
		}
		// TODO Validate if client certificate looks good
		siteState.Certificates[tlsClientName] = v1alpha1.Certificate{
			ObjectMeta: metav1.ObjectMeta{
				Name: tlsClientName,
			},
			Spec: v1alpha1.CertificateSpec{
				Subject: "127.0.0.1",
				Hosts:   []string{"127.0.0.1", "localhost"},
				Ca:      tlsCaName,
				Client:  true,
			},
		}
	}
	return nil
}

func PrepareLinkAccessesCertificates(siteState *apis.SiteState) {
	caName := fmt.Sprintf("skupper-site-ca")
	siteState.Certificates[caName] = v1alpha1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Name: caName,
		},
		Spec: v1alpha1.CertificateSpec{
			Subject: caName,
			Signing: true,
		},
	}

	for name, linkAccess := range siteState.LinkAccesses {
		create := false
		for _, role := range linkAccess.Spec.Roles {
			if utils.StringSliceContains(validLinkAccessRoles, role.Role) {
				create = true
				break
			}
		}
		if !create {
			continue
		}
		hosts := linkAccess.Spec.SubjectAlternativeNames
		if linkAccess.Spec.BindHost != "" && !utils.StringSliceContains(hosts, linkAccess.Spec.BindHost) {
			hosts = append(hosts, linkAccess.Spec.BindHost)
		}
		linkAccessCaName := caName
		if linkAccess.Spec.Ca != "" {
			linkAccessCaName = linkAccess.Spec.Ca
		}
		siteState.Certificates[name] = v1alpha1.Certificate{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: v1alpha1.CertificateSpec{
				Ca:      linkAccessCaName,
				Subject: name,
				Hosts:   hosts,
				Server:  true,
			},
		}
		clientCertificateName := fmt.Sprintf("client-%s", name)
		siteState.Certificates[clientCertificateName] = v1alpha1.Certificate{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name: clientCertificateName,
			},
			Spec: v1alpha1.CertificateSpec{
				Ca:      linkAccessCaName,
				Subject: clientCertificateName,
				Client:  true,
			},
		}
	}
	return
}

func PrepareTCPCertificates(siteState *apis.SiteState) {
	caName := fmt.Sprintf("skupper-service-ca")
	siteState.Certificates[caName] = v1alpha1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Name: caName,
		},
		Spec: v1alpha1.CertificateSpec{
			Subject: caName,
			Signing: true,
		},
	}
	// TODO How can we differentiate a listener that does simple tls (CA only) vs mutual tls auth?
	// 	    Should we introduce a "CA" field or should we inspect the content of the tlsCredential?
	for _, listener := range siteState.Listeners {
		if listener.Spec.TlsCredentials != "" {
			siteState.Certificates[listener.Spec.TlsCredentials] = v1alpha1.Certificate{
				ObjectMeta: metav1.ObjectMeta{
					Name: listener.Spec.TlsCredentials,
				},
				Spec: v1alpha1.CertificateSpec{
					Ca:      caName,
					Subject: listener.Spec.Host,
					Hosts:   []string{listener.Spec.Host},
					Server:  true,
				},
			}
		}
	}
	for _, connector := range siteState.Connectors {
		if connector.Spec.TlsCredentials != "" {
			siteState.Certificates[connector.Spec.TlsCredentials] = v1alpha1.Certificate{
				ObjectMeta: metav1.ObjectMeta{
					Name: connector.Spec.TlsCredentials,
				},
				Spec: v1alpha1.CertificateSpec{
					Ca:      caName,
					Subject: connector.Spec.Host,
					Hosts:   []string{connector.Spec.Host},
					Server:  true,
				},
			}
		}
	}
}
