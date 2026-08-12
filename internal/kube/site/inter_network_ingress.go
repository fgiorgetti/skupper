package site

import (
	"fmt"

	"github.com/skupperproject/skupper/internal/qdr"
	"github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
)

type InterNetworkIngress struct {
	Name       string
	RoutingKey string
	LinkName   string
	Ingress    *v2alpha1.InterNetworkIngress
	Link       *NetworkLink
	AccessName string
	Access     *v2alpha1.NetworkAccess
}

func NewInterNetworkIngress(ingress *v2alpha1.InterNetworkIngress, networkLink *NetworkLink, networkAccess *v2alpha1.NetworkAccess) *InterNetworkIngress {
	inetIngress := &InterNetworkIngress{
		Name:       ingress.Name,
		RoutingKey: ingress.Spec.RoutingKey,
		LinkName:   ingress.Spec.NetworkLink,
		AccessName: ingress.Spec.NetworkAccess,
		Link:       networkLink,
		Access:     networkAccess,
		Ingress:    ingress,
	}
	return inetIngress
}

func (m *InterNetworkIngress) Update(desired *v2alpha1.InterNetworkIngress, networkLink *NetworkLink, networkAccess *v2alpha1.NetworkAccess) bool {
	if desired == nil {
		if m.Ingress == nil {
			return false
		}
		m.Ingress = nil
		return true
	}
	var update bool
	if m.RoutingKey != desired.Spec.RoutingKey {
		m.RoutingKey = desired.Spec.RoutingKey
		update = true
	}
	if m.LinkName != desired.Spec.NetworkLink {
		m.LinkName = desired.Spec.NetworkLink
		update = true
	}
	if m.AccessName != desired.Spec.NetworkAccess {
		m.AccessName = desired.Spec.NetworkAccess
		update = true
	}
	if networkLink != nil {
		if m.Link == nil {
			m.Link = networkLink
			update = true
		}
	} else if m.Link != nil {
		m.Link = nil
		update = true
	}
	if networkAccess != nil {
		if m.Access == nil {
			m.Access = networkAccess
			update = true
		}
	} else if m.Access != nil {
		m.Access = nil
		update = true
	}
	m.Ingress = desired
	return update
}

func (m *InterNetworkIngress) Apply(config *qdr.RouterConfig) bool {
	var updated bool
	if m.Ingress == nil || m.Link == nil {
		if config.RemoveAutoLink(m.AutoLinkNameForNetworkLink()) {
			updated = true
		}
	}
	if m.Ingress == nil || m.Access == nil {
		if config.RemoveAutoLink(m.AutoLinkNameForNetworkAccess()) {
			updated = true
		}
	}
	for _, autoLink := range m.DesiredAutoLinks() {
		if config.AddAutoLink(autoLink) {
			updated = true
		}
	}
	return updated
}

func (m *InterNetworkIngress) AutoLinkNameForNetworkLink() string {
	return fmt.Sprintf("%s-networklink-ingress", m.Name)
}

func (m *InterNetworkIngress) AutoLinkNameForNetworkAccess() string {
	return fmt.Sprintf("%s-networkaccess-ingress", m.Name)
}

func (m *InterNetworkIngress) DesiredAutoLinks() []qdr.AutoLink {
	var autoLinks []qdr.AutoLink
	if m.Ingress == nil {
		return autoLinks
	}
	if m.Link != nil {
		autoLinks = append(autoLinks, qdr.AutoLink{
			Name:       m.AutoLinkNameForNetworkLink(),
			Address:    m.RoutingKey,
			Direction:  qdr.DirectionIn,
			Connection: m.Link.connectorName(),
		})
	}
	if m.Access != nil {
		autoLinks = append(autoLinks, qdr.AutoLink{
			Name:       m.AutoLinkNameForNetworkAccess(),
			Address:    m.RoutingKey,
			Direction:  qdr.DirectionIn,
			Connection: m.Access.ListenerName(),
		})
	}
	return autoLinks
}
