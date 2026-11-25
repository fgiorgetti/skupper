package qdr

import (
	"fmt"
	"log"
)

func SyncSslProfilesToRouter(agentPool *AgentPool, desired map[string]SslProfile) error {

	agent, err := agentPool.Get()
	if err != nil {
		return err
	}
	defer agentPool.Put(agent)

	actual, err := agent.GetSslProfiles()
	if err != nil {
		return err
	}

	for _, profile := range desired {
		current, ok := actual[profile.Name]
		if !ok {
			if err := agent.CreateSslProfile(profile); err != nil {
				return err
			}
		}
		if current != profile {
			if err := agent.UpdateSslProfile(profile); err != nil {
				return err
			}
		}
	}
	for _, profile := range actual {
		if _, ok := desired[profile.Name]; !ok {
			if err := agent.Delete("io.skupper.router.sslProfile", profile.Name); err != nil {
				return err
			}
		}
	}
	return nil
}

type bridgeConfigAgent interface {
	GetLocalBridgeConfig() (*BridgeConfig, error)
	UpdateLocalBridgeConfig(changes *BridgeConfigDifference) error
}

func GetLocalBridgeConfig(agentPool *AgentPool) (*BridgeConfig, error) {
	agent, err := agentPool.Get()
	if err != nil {
		return nil, fmt.Errorf("Could not get management agent : %s", err)
	}
	defer agentPool.Put(agent)

	actual, err := agent.GetLocalBridgeConfig()
	if err != nil {
		return nil, fmt.Errorf("Error retrieving bridges: %s", err)
	}
	return actual, nil
}

func SyncBridgeConfig(agentPool *AgentPool, desired *BridgeConfig) error {
	_, err := syncBridgeConfigFromPool(agentPool, desired, false)
	return err
}

// SyncAndGetBridgeConfig reconciles bridge config and returns the live router bridge config.
func SyncAndGetBridgeConfig(agentPool *AgentPool, desired *BridgeConfig) (*BridgeConfig, error) {
	return syncBridgeConfigFromPool(agentPool, desired, true)
}

func syncBridgeConfigFromPool(agentPool *AgentPool, desired *BridgeConfig, fetchAfterUpdate bool) (*BridgeConfig, error) {
	agent, err := agentPool.Get()
	if err != nil {
		return nil, fmt.Errorf("Could not get management agent : %s", err)
	}
	defer agentPool.Put(agent)

	actual, synced, err := syncBridgeConfig(agent, desired, fetchAfterUpdate)
	if err != nil {
		return nil, fmt.Errorf("Error while syncing bridge config : %s", err)
	}
	if !synced {
		log.Default().Println("Bridge config is not synchronised yet")
	}
	return actual, nil
}

func SyncRouterConfig(agentPool *AgentPool, desired *RouterConfig, checkCertFilesExist bool) error {
	if err := syncConnectors(agentPool, desired, checkCertFilesExist); err != nil {
		return err
	}
	if err := syncListeners(agentPool, desired); err != nil {
		return err
	}
	return nil
}

func syncConnectors(agentPool *AgentPool, desired *RouterConfig, checkCertFilesExist bool) error {
	agent, err := agentPool.Get()
	if err != nil {
		return err
	}
	defer agentPool.Put(agent)

	actual, err := agent.GetLocalConnectors()
	if err != nil {
		return fmt.Errorf("Error retrieving local connectors: %s", err)
	}

	ignorePrefix := "auto-mesh"
	if differences := ConnectorsDifference(actual, desired, &ignorePrefix); !differences.Empty() {
		if err = agent.UpdateConnectorConfig(differences, checkCertFilesExist); err != nil {
			return fmt.Errorf("Error syncing connectors: %s", err)
		}
	}
	return nil
}

func syncListeners(agentPool *AgentPool, desired *RouterConfig) error {
	agent, err := agentPool.Get()
	if err != nil {
		return err
	}
	defer agentPool.Put(agent)

	actual, err := agent.GetLocalListeners()
	if err != nil {
		return fmt.Errorf("Error retrieving local listeners: %s", err)
	}

	if differences := ListenersDifference(FilterListeners(actual, IsNotProtectedListener), desired.GetMatchingListeners(IsNotProtectedListener)); !differences.Empty() {
		if err := agent.UpdateListenerConfig(differences); err != nil {
			return fmt.Errorf("Error syncing listeners: %s", err)
		}
	}
	return nil
}

func syncBridgeConfig(agent bridgeConfigAgent, desired *BridgeConfig, fetchAfterUpdate bool) (*BridgeConfig, bool, error) {
	actual, err := agent.GetLocalBridgeConfig()
	if err != nil {
		return nil, false, fmt.Errorf("Error retrieving bridges: %s", err)
	}
	differences := actual.Difference(desired)
	if differences.Empty() {
		return actual, true, nil
	}
	if err = agent.UpdateLocalBridgeConfig(differences); err != nil {
		return nil, false, fmt.Errorf("Error syncing bridges: %s", err)
	}
	if !fetchAfterUpdate {
		return nil, false, nil
	}
	updated, err := agent.GetLocalBridgeConfig()
	if err != nil {
		return nil, false, fmt.Errorf("Error retrieving updated bridges: %s", err)
	}
	return updated, false, nil
}

func SyncProxyProfilesToRouter(agentPool *AgentPool, desired map[string]ProxyProfile) error {
	agent, err := agentPool.Get()
	if err != nil {
		return err
	}
	defer agentPool.Put(agent)

	actual, err := agent.GetProxyProfiles()
	if err != nil {
		return err
	}

	for _, profile := range desired {
		current, ok := actual[profile.Name]
		if !ok {
			if err := agent.CreateProxyProfile(profile); err != nil {
				return err
			}
			continue
		}
		if current != profile {
			if err := agent.UpdateProxyProfile(profile); err != nil {
				return err
			}
		}
	}
	for _, profile := range actual {
		if _, ok := desired[profile.Name]; !ok {
			if err := agent.Delete("io.skupper.router.proxyProfile", profile.Name); err != nil {
				return err
			}
		}
	}
	return nil
}
