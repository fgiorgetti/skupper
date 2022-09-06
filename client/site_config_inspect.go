package client

import (
	"context"
	"strconv"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/api/types"
)

func (cli *VanClient) SiteConfigInspect(ctx context.Context, input *corev1.ConfigMap) (*types.SiteConfig, error) {
	var namespace string
	if input == nil {
		namespace = cli.Namespace
	} else {
		namespace = input.ObjectMeta.Namespace
	}
	return cli.SiteConfigInspectInNamespace(ctx, input, namespace)
}

func (cli *VanClient) SiteConfigInspectInNamespace(ctx context.Context, input *corev1.ConfigMap, namespace string) (*types.SiteConfig, error) {
	var siteConfig *corev1.ConfigMap
	if input == nil {
		cm, err := cli.KubeClient.CoreV1().ConfigMaps(namespace).Get(types.SiteConfigMapName, metav1.GetOptions{})
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

	if result.Spec.SkupperName == "" {
		result.Spec.SkupperName = namespace
	}
	if result.Spec.Ingress == "" {
		result.Spec.Ingress = cli.GetIngressDefault()
	}
	if siteConfig.ObjectMeta.Labels == nil {
		result.Spec.SiteControlled = true
	} else if ignore, ok := siteConfig.ObjectMeta.Labels[types.SiteControllerIgnore]; ok {
		siteIgnore, _ := strconv.ParseBool(ignore)
		result.Spec.SiteControlled = !siteIgnore
	} else {
		result.Spec.SiteControlled = true
	}
	result.Reference.UID = string(siteConfig.ObjectMeta.UID)
	result.Reference.Name = siteConfig.ObjectMeta.Name
	result.Reference.Kind = siteConfig.TypeMeta.Kind
	result.Reference.APIVersion = siteConfig.TypeMeta.APIVersion

	if routerLogging, ok := siteConfig.Data[SiteConfigRouterLoggingKey]; ok && routerLogging != "" {
		logConf, err := ParseRouterLogConfig(routerLogging)
		if err != nil {
			return result, err
		}
		result.Spec.Router.Logging = logConf
	}

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
