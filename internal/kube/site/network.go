package site

import (
	"log/slog"
	"slices"

	"github.com/skupperproject/skupper/internal/qdr"
	"github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
)

type Network struct {
	NetworkId string
	Image     string
}

func (m *Network) Apply(config *qdr.RouterConfig) bool {
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

func (m *Network) Equals(other v2alpha1.NetworkSpec) bool {
	return m.NetworkId == other.NetworkId &&
		m.Image == other.Image
}

func (m *Network) Update(namespace, name string, network *v2alpha1.Network) bool {
	logger := slog.New(slog.Default().Handler()).With(
		slog.String("namespace", namespace),
		slog.String("name", name))
	var update bool
	if network == nil {
		if m.NetworkId != "" {
			logger.Info("Network removed, removing network id",
				slog.String("networkId", m.NetworkId))
			m.NetworkId = ""
			update = true
		}
	} else {
		if !m.Equals(network.Spec) {
			logger.Info("Network changed, updating network id",
				slog.String("networkId", network.Spec.NetworkId))
			m.Image = network.Spec.Image
			m.NetworkId = network.Spec.NetworkId
			update = true
		}
	}
	return update
}

func (m *Network) init(config *qdr.RouterConfig) {
	if config == nil {
		return
	}
	m.NetworkId = config.Network.NetworkId
}
