package site

import (
	"fmt"

	"github.com/skupperproject/skupper/internal/qdr"
	"github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
)

type InterNetworkIngress struct {
	Name     string
	Address  string
	LinkName string
	Ingress  *v2alpha1.InterNetworkIngress
	Link     *NetworkLink
}

func NewInterNetworkIngress(ingress *v2alpha1.InterNetworkIngress, networkLink *NetworkLink) *InterNetworkIngress {
	inetIngress := &InterNetworkIngress{
		Name:     ingress.Name,
		Address:  ingress.Spec.Address,
		LinkName: ingress.Spec.NetworkLink,
		Link:     networkLink,
		Ingress:  ingress,
	}
	return inetIngress
}

func (m *InterNetworkIngress) Update(desired *v2alpha1.InterNetworkIngress, networkLink *NetworkLink) bool {
	if desired == nil {
		m.Ingress = nil
		return true
	}
	var update bool
	if m.Address != desired.Spec.Address {
		m.Address = desired.Spec.Address
		update = true
	}
	if m.LinkName != desired.Spec.NetworkLink {
		m.LinkName = desired.Spec.NetworkLink
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
	m.Ingress = desired
	return update
}

func (m *InterNetworkIngress) Apply(config *qdr.RouterConfig) bool {
	if m.Ingress == nil || m.Link == nil {
		config.RemoveAutoLink(m.AutoLinkName())
		return true
	}
	autoLink := qdr.AutoLink{
		Name:       m.AutoLinkName(),
		Address:    m.Address,
		Direction:  qdr.DirectionIn,
		Connection: m.Link.connectorName(),
	}
	if config.AddAutoLink(autoLink) {
		return true
	}
	return false
}

func (m *InterNetworkIngress) AutoLinkName() string {
	return fmt.Sprintf("%s-inter-network-ingress", m.Name)
}
