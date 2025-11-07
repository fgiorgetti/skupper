package site

import (
	"log/slog"
	"slices"

	"github.com/skupperproject/skupper/internal/qdr"
	"github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
)

type ManagedSite struct {
	NetworkId string
	Image     string
}

func (m *ManagedSite) Apply(config *qdr.RouterConfig) bool {
	if config.Network.NetworkId == m.NetworkId {
		return false
	}
	config.Network.NetworkId = m.NetworkId
	var interNetworkConnectors []string
	var autoLinks []string
	if m.NetworkId == "" {
		for name, conn := range config.Connectors {
			if conn.IsInterNetworkConnector() {
				interNetworkConnectors = append(interNetworkConnectors, name)
			}
		}
		for name, autoLink := range config.AutoLinks {
			if slices.Contains(interNetworkConnectors, autoLink.Connection) {
				autoLinks = append(autoLinks, name)
			}
		}
		for _, name := range interNetworkConnectors {
			config.RemoveConnector(name)
		}
		for _, name := range autoLinks {
			config.RemoveAutoLink(name)
		}
	}
	return true
}

func (m *ManagedSite) Equals(other v2alpha1.ManagedSiteSpec) bool {
	return m.NetworkId == other.NetworkId &&
		m.Image == other.Image
}

func (m *ManagedSite) Update(namespace, name string, managedSite *v2alpha1.ManagedSite) bool {
	logger := slog.New(slog.Default().Handler()).With(
		slog.String("namespace", namespace),
		slog.String("name", name))
	var update bool
	if managedSite == nil {
		if m.NetworkId != "" {
			logger.Info("ManagedSite removed, removing network id",
				slog.String("networkId", m.NetworkId))
			m.NetworkId = ""
			update = true
		}
	} else {
		if !m.Equals(managedSite.Spec) {
			logger.Info("ManagedSite changed, updating network id",
				slog.String("networkId", managedSite.Spec.NetworkId))
			m.Image = managedSite.Spec.Image
			m.NetworkId = managedSite.Spec.NetworkId
			update = true
		}
	}
	return update
}

func (m *ManagedSite) init(config *qdr.RouterConfig) {
	if config == nil {
		return
	}
	m.NetworkId = config.Network.NetworkId
}
