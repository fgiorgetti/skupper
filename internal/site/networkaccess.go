package site

import (
	"fmt"

	"github.com/skupperproject/skupper/internal/qdr"
	skupperv2alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
)

type NetworkAccessMap map[string]*skupperv2alpha1.NetworkAccess

func (m NetworkAccessMap) desiredListeners() map[string]qdr.Listener {
	role := "inter-network"
	desired := map[string]qdr.Listener{}
	for _, na := range m {
		name := na.ListenerName()
		desired[name] = qdr.Listener{
			Name:             name,
			Role:             qdr.GetRole(role),
			Host:             na.Spec.BindHost,
			Port:             int32(na.GetPort()),
			SslProfile:       na.Spec.TlsCredentials,
			SaslMechanisms:   "EXTERNAL",
			AuthenticatePeer: true,
		}
	}
	return desired
}

func (m NetworkAccessMap) DesiredConfig(targetGroups []string, profilePath string) *NetworkAccessConfig {
	return &NetworkAccessConfig{
		listeners:   m.desiredListeners(),
		profilePath: profilePath,
	}
}

type NetworkAccessConfig struct {
	listeners   map[string]qdr.Listener
	profilePath string
}

func (g *NetworkAccessConfig) Apply(config *qdr.RouterConfig) bool {
	changed := false
	networkId := config.Network.NetworkId
	lc := qdr.ListenersDifference(config.GetMatchingListeners(qdr.IsInterNetworkNotProtectedListener), g.listeners)
	// delete before add with listeners, as changes are handled as delete and add
	for _, value := range lc.Deleted {
		if removed, _ := config.RemoveListener(value.Name); removed {
			delete(config.Listeners, value.Name)
			config.RemoveAutoLink(g.autoLinkName(value.Name))
			changed = true
		}
	}
	for _, value := range lc.Added {
		if config.AddListener(value) {
			config.AddSslProfile(qdr.ConfigureSslProfile(value.SslProfile, g.profilePath, true))
			changed = true
		}
	}
	// SslProfiles may be shared, so only delete those that are now unreferenced
	for name, _ := range config.UnreferencedSslProfiles() {
		config.RemoveSslProfile(name)
		changed = true
	}
	if networkId == "" {
		if g.removeListenerAutoLinks(config) {
			changed = true
		}
	} else {
		if g.addListenerAutoLinks(config) {
			changed = true
		}
	}
	return changed
}

func (g *NetworkAccessConfig) autoLinkName(listenerName string) string {
	return fmt.Sprintf("%s-listener-autoLink", listenerName)
}

func (g *NetworkAccessConfig) autoLinkExternalAddress(networkId string) string {
	return fmt.Sprintf("_xtopo/%s", networkId)
}

func (g *NetworkAccessConfig) removeListenerAutoLinks(config *qdr.RouterConfig) bool {
	changed := false
	for name, _ := range config.GetMatchingListeners(qdr.IsInterNetworkNotProtectedListener) {
		if config.RemoveAutoLink(g.autoLinkName(name)) {
			changed = true
		}
	}
	return changed
}

func (g *NetworkAccessConfig) addListenerAutoLinks(config *qdr.RouterConfig) bool {
	changed := false
	for name, _ := range config.GetMatchingListeners(qdr.IsInterNetworkNotProtectedListener) {
		if config.AddAutoLink(qdr.AutoLink{
			Name:            g.autoLinkName(name),
			ExternalAddress: g.autoLinkExternalAddress(config.Network.NetworkId),
			Direction:       qdr.DirectionIn,
			Connection:      name,
		}) {
			changed = true
		}
	}
	return changed
}
