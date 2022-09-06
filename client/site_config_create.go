package client

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/api/types"
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

func (cli *VanClient) SiteConfigCreate(ctx context.Context, spec types.SiteConfigSpec) (*types.SiteConfig, error) {
	siteConfig := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        types.SiteConfigMapName,
			Annotations: spec.Annotations,
			Labels:      spec.Labels,
		},
		Data: map[string]string{
			SiteConfigNameKey:                  cli.Namespace,
			SiteConfigRouterModeKey:            string(types.TransportModeInterior),
			SiteConfigServiceControllerKey:     "true",
			SiteConfigServiceSyncKey:           "true",
			SiteConfigConsoleKey:               "true",
			SiteConfigRouterConsoleKey:         "false",
			SiteConfigRouterLoggingKey:         "",
			SiteConfigConsoleAuthenticationKey: types.ConsoleAuthModeInternal,
			SiteConfigConsoleUserKey:           "",
			SiteConfigConsolePasswordKey:       "",
			SiteConfigIngressKey:               types.IngressLoadBalancerString,
		},
	}
	specAsMap, err := spec.ToMap()
	if err != nil {
		return nil, err
	}
	// merging with config map data
	for k, v := range specAsMap {
		siteConfig.Data[k] = v
	}

	// Site controlled verification
	if !spec.SiteControlled {
		if siteConfig.ObjectMeta.Labels == nil {
			siteConfig.ObjectMeta.Labels = map[string]string{}
		}
		siteConfig.ObjectMeta.Labels[types.SiteControllerIgnore] = "true"
	}
	// Extra labels
	if DefaultSkupperExtraLabels != "" {
		labelRegex := regexp.MustCompile(ValidRfc1123Label)
		if labelRegex.MatchString(DefaultSkupperExtraLabels) {
			s := strings.Split(DefaultSkupperExtraLabels, ",")
			for _, kv := range s {
				parts := strings.Split(kv, "=")
				if len(parts) > 1 {
					siteConfig.ObjectMeta.Labels[parts[0]] = parts[1]
				}
			}
		}
	}

	if spec.IsIngressRoute() && cli.RouteClient == nil {
		return nil, fmt.Errorf("OpenShift cluster not detected for --ingress type route")
	}

	actual, err := cli.KubeClient.CoreV1().ConfigMaps(cli.Namespace).Create(siteConfig)
	if err != nil {
		return nil, err
	}
	if actual.TypeMeta.Kind == "" || actual.TypeMeta.APIVersion == "" { // why??
		actual.TypeMeta = siteConfig.TypeMeta
	}
	return cli.SiteConfigInspect(ctx, actual)
}
