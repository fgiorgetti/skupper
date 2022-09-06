package types

import (
	"fmt"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	// core options
	SiteConfigNameKey                string = "name"
	SiteConfigRouterModeKey          string = "router-mode"
	SiteConfigIngressKey             string = "ingress"
	SiteConfigIngressAnnotationsKey  string = "ingress-annotations"
	SiteConfigIngressHostKey         string = "ingress-host"
	SiteConfigCreateNetworkPolicyKey string = "create-network-policy"
	SiteConfigRoutersKey             string = "routers"

	// console options
	SiteConfigConsoleKey               string = "console"
	SiteConfigConsoleAuthenticationKey string = "console-authentication"
	SiteConfigConsoleUserKey           string = "console-user"
	SiteConfigConsolePasswordKey       string = "console-password"
	SiteConfigConsoleIngressKey        string = "console-ingress"

	// router options
	SiteConfigRouterConsoleKey            string = "router-console"
	SiteConfigRouterLoggingKey            string = "router-logging"
	SiteConfigRouterDebugModeKey          string = "router-debug-mode"
	SiteConfigRouterCpuKey                string = "router-cpu"
	SiteConfigRouterMemoryKey             string = "router-memory"
	SiteConfigRouterCpuLimitKey           string = "router-cpu-limit"
	SiteConfigRouterMemoryLimitKey        string = "router-memory-limit"
	SiteConfigRouterAffinityKey           string = "router-pod-affinity"
	SiteConfigRouterAntiAffinityKey       string = "router-pod-antiaffinity"
	SiteConfigRouterNodeSelectorKey       string = "router-node-selector"
	SiteConfigRouterMaxFrameSizeKey       string = "xp-router-max-frame-size"
	SiteConfigRouterMaxSessionFramesKey   string = "xp-router-max-session-frames"
	SiteConfigRouterIngressHostKey        string = "router-ingress-host"
	SiteConfigRouterServiceAnnotationsKey string = "router-service-annotations"
	SiteConfigRouterLoadBalancerIp        string = "router-load-balancer-ip"

	// controller options
	SiteConfigServiceControllerKey            string = "service-controller"
	SiteConfigServiceSyncKey                  string = "service-sync"
	SiteConfigControllerCpuKey                string = "controller-cpu"
	SiteConfigControllerMemoryKey             string = "controller-memory"
	SiteConfigControllerCpuLimitKey           string = "controller-cpu-limit"
	SiteConfigControllerMemoryLimitKey        string = "controller-memory-limit"
	SiteConfigControllerAffinityKey           string = "controller-pod-affinity"
	SiteConfigControllerAntiAffinityKey       string = "controller-pod-antiaffinity"
	SiteConfigControllerNodeSelectorKey       string = "controller-node-selector"
	SiteConfigControllerIngressHostKey        string = "controller-ingress-host"
	SiteConfigControllerServiceAnnotationsKey string = "controller-service-annotations"
	SiteConfigControllerLoadBalancerIp        string = "controller-load-balancer-ip"

	// config-sync options
	SiteConfigConfigSyncCpuKey         string = "config-sync-cpu"
	SiteConfigConfigSyncMemoryKey      string = "config-sync-memory"
	SiteConfigConfigSyncCpuLimitKey    string = "config-sync-cpu-limit"
	SiteConfigConfigSyncMemoryLimitKey string = "config-sync-memory-limit"
)

type SiteConfig struct {
	Spec      SiteConfigSpec
	Reference SiteConfigReference
}

type SiteConfigReference struct {
	UID        string
	Name       string
	APIVersion string
	Kind       string
}

type SiteConfigSpec struct {
	SkupperName         string
	SkupperNamespace    string
	RouterMode          string
	Routers             int
	EnableController    bool
	EnableServiceSync   bool
	EnableConsole       bool
	AuthMode            string
	User                string
	Password            string
	Ingress             string
	IngressAnnotations  map[string]string
	ConsoleIngress      string
	IngressHost         string
	Replicas            int32
	SiteControlled      bool
	CreateNetworkPolicy bool
	Annotations         map[string]string
	Labels              map[string]string
	Router              RouterOptions
	Controller          ControllerOptions
	ConfigSync          ConfigSyncOptions
	Platform            Platform
}

func (s *SiteConfigSpec) IsIngressRoute() bool {
	return s.Ingress == IngressRouteString
}
func (s *SiteConfigSpec) IsIngressLoadBalancer() bool {
	return s.Ingress == IngressLoadBalancerString
}
func (s *SiteConfigSpec) IsIngressNodePort() bool {
	return s.Ingress == IngressNodePortString
}
func (s *SiteConfigSpec) IsIngressNginxIngress() bool {
	return s.Ingress == IngressNginxIngressString
}
func (s *SiteConfigSpec) IsIngressContourHttpProxy() bool {
	return s.Ingress == IngressContourHttpProxyString
}
func (s *SiteConfigSpec) IsIngressKubernetes() bool {
	return s.Ingress == IngressKubernetes
}
func (s *SiteConfigSpec) IsIngressNone() bool {
	return s.Ingress == IngressNoneString
}

func (s *SiteConfigSpec) IsConsoleIngressRoute() bool {
	return s.getConsoleIngress() == IngressRouteString
}
func (s *SiteConfigSpec) IsConsoleIngressLoadBalancer() bool {
	return s.getConsoleIngress() == IngressLoadBalancerString
}
func (s *SiteConfigSpec) IsConsoleIngressNodePort() bool {
	return s.getConsoleIngress() == IngressNodePortString
}
func (s *SiteConfigSpec) IsConsoleIngressNginxIngress() bool {
	return s.getConsoleIngress() == IngressNginxIngressString
}
func (s *SiteConfigSpec) IsConsoleIngressContourHttpProxy() bool {
	return s.getConsoleIngress() == IngressContourHttpProxyString
}
func (s *SiteConfigSpec) IsConsoleIngressKubernetes() bool {
	return s.getConsoleIngress() == IngressKubernetes
}
func (s *SiteConfigSpec) IsConsoleIngressNone() bool {
	return s.getConsoleIngress() == IngressNoneString
}
func (s *SiteConfigSpec) getConsoleIngress() string {
	if s.ConsoleIngress == "" {
		return s.Ingress
	}
	return s.ConsoleIngress
}

func (s *SiteConfigSpec) CheckIngress() error {
	if !isValidIngress(s.Platform, s.Ingress) {
		return fmt.Errorf("Invalid value for ingress: %s", s.Ingress)
	}
	return nil
}

func (s *SiteConfigSpec) CheckConsoleIngress() error {
	if !isValidIngress(s.Platform, s.ConsoleIngress) {
		return fmt.Errorf("Invalid value for console-ingress: %s", s.ConsoleIngress)
	}
	return nil
}

func (s *SiteConfigSpec) GetRouterIngressHost() string {
	if s.Router.IngressHost != "" {
		return s.Router.IngressHost
	}
	return s.IngressHost
}

func (s *SiteConfigSpec) GetControllerIngressHost() string {
	if s.Controller.IngressHost != "" {
		return s.Controller.IngressHost
	}
	return s.IngressHost
}

func (s *SiteConfigSpec) ToMap() (map[string]string, error) {
	data := map[string]string{}

	if s.SkupperName != "" {
		data[SiteConfigNameKey] = s.SkupperName
	}
	if s.RouterMode != "" {
		data[SiteConfigRouterModeKey] = s.RouterMode
	}
	if s.Routers != 0 {
		data[SiteConfigRoutersKey] = strconv.Itoa(s.Routers)
	}
	if !s.EnableController {
		data[SiteConfigServiceControllerKey] = "false"
	}
	if !s.EnableServiceSync {
		data[SiteConfigServiceSyncKey] = "false"
	}
	if !s.EnableConsole {
		data[SiteConfigConsoleKey] = "false"
	}
	if s.AuthMode != "" {
		data[SiteConfigConsoleAuthenticationKey] = s.AuthMode
	}
	if s.User != "" {
		data[SiteConfigConsoleUserKey] = s.User
	}
	if s.Password != "" {
		data[SiteConfigConsolePasswordKey] = s.Password
	}
	if s.Ingress != "" {
		data[SiteConfigIngressKey] = s.Ingress
	}
	if len(s.IngressAnnotations) > 0 {
		var annotations []string
		for key, value := range s.IngressAnnotations {
			annotations = append(annotations, key+"="+value)
		}
		data[SiteConfigIngressAnnotationsKey] = strings.Join(annotations, ",")
	}
	if s.ConsoleIngress != "" {
		data[SiteConfigConsoleIngressKey] = s.ConsoleIngress
	}
	if s.IngressHost != "" {
		data[SiteConfigIngressHostKey] = s.IngressHost
	}
	if s.CreateNetworkPolicy {
		data[SiteConfigCreateNetworkPolicyKey] = "true"
	}
	if s.Router.Logging != nil {
		data[SiteConfigRouterLoggingKey] = RouterLogConfigToString(s.Router.Logging)
	}
	if s.Router.DebugMode != "" {
		data[SiteConfigRouterDebugModeKey] = s.Router.DebugMode
	}
	if s.Router.Cpu != "" {
		if _, err := resource.ParseQuantity(s.Router.Cpu); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigRouterCpuKey, s.Router.Cpu, err)
		}
		data[SiteConfigRouterCpuKey] = s.Router.Cpu
	}
	if s.Router.Memory != "" {
		if _, err := resource.ParseQuantity(s.Router.Memory); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigRouterMemoryKey, s.Router.Memory, err)
		}
		data[SiteConfigRouterMemoryKey] = s.Router.Memory
	}
	if s.Router.CpuLimit != "" {
		if _, err := resource.ParseQuantity(s.Router.CpuLimit); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigRouterCpuLimitKey, s.Router.CpuLimit, err)
		}
		data[SiteConfigRouterCpuLimitKey] = s.Router.CpuLimit
	}
	if s.Router.MemoryLimit != "" {
		if _, err := resource.ParseQuantity(s.Router.MemoryLimit); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigRouterMemoryLimitKey, s.Router.MemoryLimit, err)
		}
		data[SiteConfigRouterMemoryLimitKey] = s.Router.MemoryLimit
	}
	if s.Router.Affinity != "" {
		data[SiteConfigRouterAffinityKey] = s.Router.Affinity
	}
	if s.Router.AntiAffinity != "" {
		data[SiteConfigRouterAntiAffinityKey] = s.Router.AntiAffinity
	}
	if s.Router.NodeSelector != "" {
		data[SiteConfigRouterNodeSelectorKey] = s.Router.NodeSelector
	}
	if s.Router.IngressHost != "" {
		data[SiteConfigRouterIngressHostKey] = s.Router.IngressHost
	}
	if s.Router.MaxFrameSize != RouterMaxFrameSizeDefault {
		data[SiteConfigRouterMaxFrameSizeKey] = strconv.Itoa(s.Router.MaxFrameSize)
	}
	if s.Router.MaxSessionFrames != RouterMaxSessionFramesDefault {
		data[SiteConfigRouterMaxSessionFramesKey] = strconv.Itoa(s.Router.MaxSessionFrames)
	}
	if len(s.Router.ServiceAnnotations) > 0 {
		var annotations []string
		for key, value := range s.Router.ServiceAnnotations {
			annotations = append(annotations, key+"="+value)
		}
		data[SiteConfigRouterServiceAnnotationsKey] = strings.Join(annotations, ",")
	}
	if s.Router.LoadBalancerIp != "" {
		data[SiteConfigRouterLoadBalancerIp] = s.Router.LoadBalancerIp
	}
	if s.Controller.Cpu != "" {
		if _, err := resource.ParseQuantity(s.Controller.Cpu); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigControllerCpuKey, s.Controller.Cpu, err)
		}
		data[SiteConfigControllerCpuKey] = s.Controller.Cpu
	}
	if s.Controller.Memory != "" {
		if _, err := resource.ParseQuantity(s.Controller.Memory); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigControllerMemoryKey, s.Controller.Memory, err)
		}
		data[SiteConfigControllerMemoryKey] = s.Controller.Memory
	}
	if s.Controller.CpuLimit != "" {
		if _, err := resource.ParseQuantity(s.Controller.CpuLimit); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigControllerCpuLimitKey, s.Controller.CpuLimit, err)
		}
		data[SiteConfigControllerCpuLimitKey] = s.Controller.CpuLimit
	}
	if s.Controller.MemoryLimit != "" {
		if _, err := resource.ParseQuantity(s.Controller.MemoryLimit); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigControllerMemoryLimitKey, s.Controller.MemoryLimit, err)
		}
		data[SiteConfigControllerMemoryLimitKey] = s.Controller.MemoryLimit
	}
	if s.Controller.Affinity != "" {
		data[SiteConfigControllerAffinityKey] = s.Controller.Affinity
	}
	if s.Controller.AntiAffinity != "" {
		data[SiteConfigControllerAntiAffinityKey] = s.Controller.AntiAffinity
	}
	if s.Controller.NodeSelector != "" {
		data[SiteConfigControllerNodeSelectorKey] = s.Controller.NodeSelector
	}
	if s.Controller.IngressHost != "" {
		data[SiteConfigControllerIngressHostKey] = s.Controller.IngressHost
	}
	if len(s.Controller.ServiceAnnotations) > 0 {
		var annotations []string
		for key, value := range s.Controller.ServiceAnnotations {
			annotations = append(annotations, key+"="+value)
		}
		data[SiteConfigControllerServiceAnnotationsKey] = strings.Join(annotations, ",")
	}
	if s.Controller.LoadBalancerIp != "" {
		data[SiteConfigControllerLoadBalancerIp] = s.Controller.LoadBalancerIp
	}

	if s.ConfigSync.Cpu != "" {
		if _, err := resource.ParseQuantity(s.ConfigSync.Cpu); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigConfigSyncCpuKey, s.ConfigSync.Cpu, err)
		}
		data[SiteConfigConfigSyncCpuKey] = s.ConfigSync.Cpu
	}
	if s.ConfigSync.Memory != "" {
		if _, err := resource.ParseQuantity(s.ConfigSync.Memory); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigConfigSyncMemoryKey, s.ConfigSync.Memory, err)
		}
		data[SiteConfigConfigSyncMemoryKey] = s.ConfigSync.Memory
	}
	if s.ConfigSync.CpuLimit != "" {
		if _, err := resource.ParseQuantity(s.ConfigSync.CpuLimit); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigConfigSyncCpuLimitKey, s.ConfigSync.CpuLimit, err)
		}
		data[SiteConfigConfigSyncCpuLimitKey] = s.ConfigSync.CpuLimit
	}
	if s.ConfigSync.MemoryLimit != "" {
		if _, err := resource.ParseQuantity(s.ConfigSync.MemoryLimit); err != nil {
			return nil, fmt.Errorf("Invalid value for %s %q: %s", SiteConfigConfigSyncMemoryLimitKey, s.ConfigSync.MemoryLimit, err)
		}
		data[SiteConfigConfigSyncMemoryLimitKey] = s.ConfigSync.MemoryLimit
	}
	return data, nil
}

func MapToSiteConfig(data map[string]string) (*SiteConfig, error) {
	var result SiteConfig
	if skupperName, ok := data[SiteConfigNameKey]; ok {
		result.Spec.SkupperName = skupperName
	}

	if routerMode, ok := data[SiteConfigRouterModeKey]; ok {
		result.Spec.RouterMode = routerMode
	} else {
		// check for deprecated key
		if isEdge, ok := data["edge"]; ok {
			if isEdge == "true" {
				result.Spec.RouterMode = string(TransportModeEdge)
			} else {
				result.Spec.RouterMode = string(TransportModeInterior)
			}
		} else {
			result.Spec.RouterMode = string(TransportModeInterior)
		}
	}
	if routers, ok := data[SiteConfigRoutersKey]; ok {
		result.Spec.Routers, _ = strconv.Atoi(routers)
	}
	if enableController, ok := data[SiteConfigServiceControllerKey]; ok {
		result.Spec.EnableController, _ = strconv.ParseBool(enableController)
	} else {
		result.Spec.EnableController = true
	}
	if enableServiceSync, ok := data[SiteConfigServiceSyncKey]; ok {
		result.Spec.EnableServiceSync, _ = strconv.ParseBool(enableServiceSync)
	} else {
		result.Spec.EnableServiceSync = true
	}
	if enableConsole, ok := data[SiteConfigConsoleKey]; ok {
		result.Spec.EnableConsole, _ = strconv.ParseBool(enableConsole)
	} else {
		result.Spec.EnableConsole = true
	}
	if createNetworkPolicy, ok := data[SiteConfigCreateNetworkPolicyKey]; ok {
		result.Spec.CreateNetworkPolicy, _ = strconv.ParseBool(createNetworkPolicy)
	} else {
		result.Spec.CreateNetworkPolicy = false
	}
	if authMode, ok := data[SiteConfigConsoleAuthenticationKey]; ok {
		result.Spec.AuthMode = authMode
	} else {
		result.Spec.AuthMode = ConsoleAuthModeInternal
	}
	if user, ok := data[SiteConfigConsoleUserKey]; ok {
		result.Spec.User = user
	} else {
		result.Spec.User = ""
	}
	if password, ok := data[SiteConfigConsolePasswordKey]; ok {
		result.Spec.Password = password
	} else {
		result.Spec.Password = ""
	}
	if ingress, ok := data[SiteConfigIngressKey]; ok {
		result.Spec.Ingress = ingress
	} else {
		// check for deprecated key
		if clusterLocal, ok := data["cluster-local"]; ok {
			if clusterLocal == "true" {
				result.Spec.Ingress = IngressNoneString
			} else {
				result.Spec.Ingress = IngressLoadBalancerString
			}
		}
	}
	if ingressAnnotations, ok := data[SiteConfigIngressAnnotationsKey]; ok {
		result.Spec.IngressAnnotations = asMap(splitWithEscaping(ingressAnnotations, ',', '\\'))
	}
	if consoleIngress, ok := data[SiteConfigConsoleIngressKey]; ok {
		result.Spec.ConsoleIngress = consoleIngress
	}
	if ingressHost, ok := data[SiteConfigIngressHostKey]; ok {
		result.Spec.IngressHost = ingressHost
	}

	if routerDebugMode, ok := data[SiteConfigRouterDebugModeKey]; ok && routerDebugMode != "" {
		result.Spec.Router.DebugMode = routerDebugMode
	}
	if routerCpu, ok := data[SiteConfigRouterCpuKey]; ok && routerCpu != "" {
		result.Spec.Router.Cpu = routerCpu
	}
	if routerMemory, ok := data[SiteConfigRouterMemoryKey]; ok && routerMemory != "" {
		result.Spec.Router.Memory = routerMemory
	}
	if routerCpuLimit, ok := data[SiteConfigRouterCpuLimitKey]; ok && routerCpuLimit != "" {
		result.Spec.Router.CpuLimit = routerCpuLimit
	}
	if routerMemoryLimit, ok := data[SiteConfigRouterMemoryLimitKey]; ok && routerMemoryLimit != "" {
		result.Spec.Router.MemoryLimit = routerMemoryLimit
	}
	if routerNodeSelector, ok := data[SiteConfigRouterNodeSelectorKey]; ok && routerNodeSelector != "" {
		result.Spec.Router.NodeSelector = routerNodeSelector
	}
	if routerAffinity, ok := data[SiteConfigRouterAffinityKey]; ok && routerAffinity != "" {
		result.Spec.Router.Affinity = routerAffinity
	}
	if routerAntiAffinity, ok := data[SiteConfigRouterAntiAffinityKey]; ok && routerAntiAffinity != "" {
		result.Spec.Router.AntiAffinity = routerAntiAffinity
	}
	if routerIngressHost, ok := data[SiteConfigRouterIngressHostKey]; ok && routerIngressHost != "" {
		result.Spec.Router.IngressHost = routerIngressHost
	}

	if routerMaxFrameSize, ok := data[SiteConfigRouterMaxFrameSizeKey]; ok && routerMaxFrameSize != "" {
		val, err := strconv.Atoi(routerMaxFrameSize)
		if err != nil {
			return &result, err
		}
		result.Spec.Router.MaxFrameSize = val
	} else {
		result.Spec.Router.MaxFrameSize = RouterMaxFrameSizeDefault
	}
	if routerMaxSessionFrames, ok := data[SiteConfigRouterMaxSessionFramesKey]; ok && routerMaxSessionFrames != "" {
		val, err := strconv.Atoi(routerMaxSessionFrames)
		if err != nil {
			return &result, err
		}
		result.Spec.Router.MaxSessionFrames = val
	} else {
		result.Spec.Router.MaxSessionFrames = RouterMaxSessionFramesDefault
	}

	if routerServiceAnnotations, ok := data[SiteConfigRouterServiceAnnotationsKey]; ok {
		result.Spec.Router.ServiceAnnotations = asMap(splitWithEscaping(routerServiceAnnotations, ',', '\\'))
	}
	if routerServiceLoadBalancerIp, ok := data[SiteConfigRouterLoadBalancerIp]; ok {
		result.Spec.Router.LoadBalancerIp = routerServiceLoadBalancerIp
	}

	if controllerCpu, ok := data[SiteConfigControllerCpuKey]; ok && controllerCpu != "" {
		result.Spec.Controller.Cpu = controllerCpu
	}
	if controllerMemory, ok := data[SiteConfigControllerMemoryKey]; ok && controllerMemory != "" {
		result.Spec.Controller.Memory = controllerMemory
	}
	if controllerCpuLimit, ok := data[SiteConfigControllerCpuLimitKey]; ok && controllerCpuLimit != "" {
		result.Spec.Controller.CpuLimit = controllerCpuLimit
	}
	if controllerMemoryLimit, ok := data[SiteConfigControllerMemoryLimitKey]; ok && controllerMemoryLimit != "" {
		result.Spec.Controller.MemoryLimit = controllerMemoryLimit
	}
	if controllerNodeSelector, ok := data[SiteConfigControllerNodeSelectorKey]; ok && controllerNodeSelector != "" {
		result.Spec.Controller.NodeSelector = controllerNodeSelector
	}
	if controllerAffinity, ok := data[SiteConfigControllerAffinityKey]; ok && controllerAffinity != "" {
		result.Spec.Controller.Affinity = controllerAffinity
	}
	if controllerAntiAffinity, ok := data[SiteConfigControllerAntiAffinityKey]; ok && controllerAntiAffinity != "" {
		result.Spec.Controller.AntiAffinity = controllerAntiAffinity
	}
	if controllerIngressHost, ok := data[SiteConfigControllerIngressHostKey]; ok && controllerIngressHost != "" {
		result.Spec.Controller.IngressHost = controllerIngressHost
	}
	if controllerServiceAnnotations, ok := data[SiteConfigControllerServiceAnnotationsKey]; ok {
		result.Spec.Controller.ServiceAnnotations = asMap(splitWithEscaping(controllerServiceAnnotations, ',', '\\'))
	}
	if controllerServiceLoadBalancerIp, ok := data[SiteConfigControllerLoadBalancerIp]; ok {
		result.Spec.Controller.LoadBalancerIp = controllerServiceLoadBalancerIp
	}

	if configSyncCpu, ok := data[SiteConfigConfigSyncCpuKey]; ok && configSyncCpu != "" {
		result.Spec.ConfigSync.Cpu = configSyncCpu
	}
	if configSyncMemory, ok := data[SiteConfigConfigSyncMemoryKey]; ok && configSyncMemory != "" {
		result.Spec.ConfigSync.Memory = configSyncMemory
	}
	if configSyncCpuLimit, ok := data[SiteConfigConfigSyncCpuLimitKey]; ok && configSyncCpuLimit != "" {
		result.Spec.ConfigSync.CpuLimit = configSyncCpuLimit
	}
	if configSyncMemoryLimit, ok := data[SiteConfigConfigSyncMemoryLimitKey]; ok && configSyncMemoryLimit != "" {
		result.Spec.ConfigSync.MemoryLimit = configSyncMemoryLimit
	}

	return &result, nil
}

func splitWithEscaping(s string, separator, escape byte) []string {
	var token []byte
	var tokens []string
	for i := 0; i < len(s); i++ {
		if s[i] == separator {
			tokens = append(tokens, strings.TrimSpace(string(token)))
			token = token[:0]
		} else if s[i] == escape && i+1 < len(s) {
			i++
			token = append(token, s[i])
		} else {
			token = append(token, s[i])
		}
	}
	tokens = append(tokens, strings.TrimSpace(string(token)))
	return tokens
}

func asMap(entries []string) map[string]string {
	result := map[string]string{}
	for _, entry := range entries {
		parts := strings.Split(entry, "=")
		if len(parts) > 1 {
			result[parts[0]] = parts[1]
		} else {
			result[parts[0]] = ""
		}
	}
	return result
}
