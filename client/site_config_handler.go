package client

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/skupperproject/skupper/api/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SiteConfigHandler struct {
	Cli         *VanClient
	inspectFrom *corev1.ConfigMap
}

func NewSiteConfigHandler(cli *VanClient) *SiteConfigHandler {
	return &SiteConfigHandler{
		Cli: cli,
	}
}

func NewSiteConfigHandlerFor(cli *VanClient, namespace string) *SiteConfigHandler {
	return &SiteConfigHandler{
		Cli: &VanClient{
			Namespace:     namespace,
			KubeClient:    cli.KubeClient,
			RouteClient:   cli.RouteClient,
			RestConfig:    cli.RestConfig,
			DynamicClient: cli.DynamicClient,
		},
	}
}

func (s *SiteConfigHandler) Create(spec *types.SiteConfigSpec) (*types.SiteConfig, error) {
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
			types.SiteConfigNameKey:                  s.Cli.Namespace,
			types.SiteConfigRouterModeKey:            string(types.TransportModeInterior),
			types.SiteConfigServiceControllerKey:     "true",
			types.SiteConfigServiceSyncKey:           "true",
			types.SiteConfigConsoleKey:               "true",
			types.SiteConfigRouterConsoleKey:         "false",
			types.SiteConfigRouterLoggingKey:         "",
			types.SiteConfigConsoleAuthenticationKey: types.ConsoleAuthModeInternal,
			types.SiteConfigConsoleUserKey:           "",
			types.SiteConfigConsolePasswordKey:       "",
			types.SiteConfigIngressKey:               types.IngressLoadBalancerString,
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

	// TODO: allow Replicas to be set through skupper-site configmap?
	if !spec.SiteControlled {
		if siteConfig.ObjectMeta.Labels == nil {
			siteConfig.ObjectMeta.Labels = map[string]string{}
		}
		siteConfig.ObjectMeta.Labels[types.SiteControllerIgnore] = "true"
	}
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

	if spec.IsIngressRoute() && s.Cli.RouteClient == nil {
		return nil, fmt.Errorf("OpenShift cluster not detected for --ingress type route")
	}

	actual, err := s.Cli.KubeClient.CoreV1().ConfigMaps(s.Cli.Namespace).Create(siteConfig)
	if err != nil {
		return nil, err
	}

	result := types.SiteConfig{
		Spec: *spec,
	}

	if actual.TypeMeta.Kind == "" || actual.TypeMeta.APIVersion == "" { // why??
		actual.TypeMeta = siteConfig.TypeMeta
	}

	result.Reference.UID = string(actual.ObjectMeta.UID)
	result.Reference.Name = actual.ObjectMeta.Name
	result.Reference.Kind = actual.TypeMeta.Kind
	result.Reference.APIVersion = actual.TypeMeta.APIVersion

	return &result, nil
}

func (s *SiteConfigHandler) InspectFrom(input *corev1.ConfigMap) (*types.SiteConfig, error) {
	s.inspectFrom = input
	return s.Inspect()
}

func (s *SiteConfigHandler) Inspect() (*types.SiteConfig, error) {
	var siteConfig *corev1.ConfigMap
	input := s.inspectFrom
	if input == nil {
		cm, err := s.Cli.KubeClient.CoreV1().ConfigMaps(s.Cli.Namespace).Get(types.SiteConfigMapName, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		siteConfig = cm
	} else {
		siteConfig = input
	}

	var result *types.SiteConfig
	var err error
	result, err = types.MapToSiteConfig(siteConfig.Data)
	if err != nil {
		return nil, err
	}

	// Setting the namespace
	result.Spec.SkupperNamespace = siteConfig.Namespace

	// Defining ingress if default method did not find it
	if result.Spec.Ingress == "" {
		result.Spec.Ingress = s.Cli.GetIngressDefault()
	}

	// Verifying if site-controlled
	if siteConfig.ObjectMeta.Labels == nil {
		result.Spec.SiteControlled = true
	} else if ignore, ok := siteConfig.ObjectMeta.Labels[types.SiteControllerIgnore]; ok {
		siteIgnore, _ := strconv.ParseBool(ignore)
		result.Spec.SiteControlled = !siteIgnore
	} else {
		result.Spec.SiteControlled = true
	}

	// Reference info
	result.Reference.UID = string(siteConfig.ObjectMeta.UID)
	result.Reference.Name = siteConfig.ObjectMeta.Name
	result.Reference.Kind = siteConfig.TypeMeta.Kind
	result.Reference.APIVersion = siteConfig.TypeMeta.APIVersion

	// Annotations
	annotationExclusions := []string{}
	labelExclusions := []string{}
	annotations := map[string]string{}
	for key, value := range siteConfig.ObjectMeta.Annotations {
		if key == types.AnnotationExcludes {
			annotationExclusions = strings.Split(value, ",")
		} else if key == types.LabelExcludes {
			labelExclusions = strings.Split(value, ",")
		} else {
			annotations[key] = value
		}
	}
	for _, key := range annotationExclusions {
		delete(annotations, key)
	}
	result.Spec.Annotations = annotations

	// Labels
	labels := map[string]string{}
	for key, value := range siteConfig.ObjectMeta.Labels {
		if key != types.SiteControllerIgnore {
			labels[key] = value
		}
	}
	for _, key := range labelExclusions {
		delete(labels, key)
	}
	result.Spec.Labels = labels

	return result, nil
}

func (s *SiteConfigHandler) Update(spec *types.SiteConfigSpec) (*types.SiteConfig, error) {
	configmap, err := s.Cli.KubeClient.CoreV1().ConfigMaps(s.Cli.Namespace).Get(types.SiteConfigMapName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	// For now, only update router-logging and/or router-debug-mode (TODO: update of other options)
	latestLogging := types.RouterLogConfigToString(spec.Router.Logging)
	update := false
	if configmap.Data[types.SiteConfigRouterLoggingKey] != latestLogging {
		configmap.Data[types.SiteConfigRouterLoggingKey] = latestLogging
		update = true
	}
	if configmap.Data[types.SiteConfigRouterDebugModeKey] != spec.Router.DebugMode {
		configmap.Data[types.SiteConfigRouterDebugModeKey] = spec.Router.DebugMode
		update = true
	}
	if update {
		configmap, err = s.Cli.KubeClient.CoreV1().ConfigMaps(s.Cli.Namespace).Update(configmap)
		if err != nil {
			return s.Inspect()
		}
	}
	return nil, nil
}
