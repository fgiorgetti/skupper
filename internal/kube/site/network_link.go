package site

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/skupperproject/skupper/internal/qdr"
	"github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
)

type NetworkLink struct {
	Name      string
	Link      *v2alpha1.NetworkLink
	NetworkId string
}

func (m *NetworkLink) Update(networkId string, desired *v2alpha1.NetworkLink) bool {
	var update bool
	if m.NetworkId != networkId {
		m.NetworkId = networkId
		update = true
	}
	if desired == nil {
		update = true
		m.Link = nil
	} else {
		if !reflect.DeepEqual(m.Link.Spec, desired.Spec) {
			update = true
		}
		m.Link = desired
	}
	return update
}

func (m *NetworkLink) Apply(config *qdr.RouterConfig) bool {
	if m.Link == nil || m.NetworkId == "" {
		return m.delete(config)
	}
	var update bool
	if config.AddSslProfile(qdr.ConfigureSslProfile(m.sslProfileName(), qdr.SSL_PROFILE_PATH, true)) {
		update = true
	}
	if config.AddConnector(qdr.Connector{
		Name:           m.connectorName(),
		Role:           qdr.RoleInterNetwork,
		Host:           m.Link.Spec.Hostname,
		Port:           strconv.Itoa(m.Link.Spec.Port),
		VerifyHostname: true,
		SslProfile:     m.sslProfileName(),
	}) {
		update = true
	}
	if m.NetworkId != "" {
		if config.AddAutoLink(qdr.AutoLink{
			Name:            m.autoLinkName(),
			ExternalAddress: m.autoLinkExternalAddress(),
			Direction:       qdr.DirectionIn,
			Connection:      m.connectorName(),
		}) {
			update = true
		}
	}
	return update
}

func (m *NetworkLink) delete(config *qdr.RouterConfig) bool {
	var update bool
	if ok, _ := config.RemoveConnector(m.connectorName()); ok {
		update = true
	}
	if config.RemoveAutoLink(m.autoLinkName()) {
		update = true
	}
	sslProfileName := m.sslProfileName()
	for name := range config.UnreferencedSslProfiles() {
		if name == sslProfileName {
			if config.RemoveSslProfile(sslProfileName) {
				update = true
			}
		}
	}
	return update
}

func (m *NetworkLink) sslProfileName() string {
	name := m.Name
	if m.Link != nil && m.Link.Spec.TlsCredentials != "" {
		name = m.Link.Spec.TlsCredentials
	}
	return fmt.Sprintf("%s-profile", name)
}

func (m *NetworkLink) connectorName() string {
	return fmt.Sprintf("%s-networklink-connector", m.Name)
}

func (m *NetworkLink) autoLinkName() string {
	return fmt.Sprintf("%s-connector-autoLink", m.Name)
}

func (m *NetworkLink) autoLinkExternalAddress() string {
	return fmt.Sprintf("_xtopo/%s", m.NetworkId)
}
