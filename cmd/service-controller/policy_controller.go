package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client"
	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/generated/client/clientset/versioned"
	v1alpha12 "github.com/skupperproject/skupper/pkg/generated/client/informers/externalversions/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/kube"
	"github.com/skupperproject/skupper/pkg/qdr"
	"github.com/skupperproject/skupper/pkg/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/skupperproject/skupper/pkg/event"
)

type PolicyController struct {
	name      string
	cli       *client.VanClient
	validator *client.ClusterPolicyValidator
	informer  cache.SharedIndexInformer
	queue     workqueue.RateLimitingInterface
}

func (c *PolicyController) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err == nil {
		c.queue.Add(key)
	} else {
		event.Recordf(c.name, "Error retrieving key: %s", err)
	}
}

func (c *PolicyController) OnAdd(obj interface{}) {
	c.enqueue(obj)
}

func (c *PolicyController) OnUpdate(a, b interface{}) {
	aa := a.(*v1alpha1.SkupperClusterPolicy)
	bb := b.(*v1alpha1.SkupperClusterPolicy)
	if aa.ResourceVersion != bb.ResourceVersion {
		c.enqueue(b)
	}
}

func (c *PolicyController) OnDelete(obj interface{}) {
	c.enqueue(obj)
}

func (c *PolicyController) start(stopCh <-chan struct{}) error {
	go c.informer.Run(stopCh)
	if ok := cache.WaitForCacheSync(stopCh, c.informer.HasSynced); !ok {
		return fmt.Errorf("Failed to wait for caches to sync")
	}
	go wait.Until(c.run, time.Second, stopCh)
	return nil
}

func (c *PolicyController) stop() {
	c.queue.ShutDown()
}

func (c *PolicyController) run() {
	for c.process() {
	}
}

func (c *PolicyController) process() bool {
	obj, shutdown := c.queue.Get()

	if shutdown {
		return false
	}

	defer c.queue.Done(obj)
	retry := false
	if key, ok := obj.(string); ok {

		log.Println("Policy has changed:", key)

		// Validate incomingLink stage changed
		c.validateIncomingLinkStateChanged()

		// Validate outgoingLink state changed
		c.validateOutgoingLinkStateChanged()

		// Valid gateway state changed
		c.validateGatewayStateChanged()

		// Validate expose state changed
		c.validateExposeStateChanged()

		// Validate service state changed
		c.validateServiceStateChanged()
	} else {
		event.Recordf(c.name, "Expected key to be string, was %#v", key)
	}
	if retry && c.queue.NumRequeues(obj) < 5 {
		c.queue.AddRateLimited(obj)
	} else {
		c.queue.Forget(obj)
	}

	return true
}

func (c *PolicyController) validateIncomingLinkStateChanged() {
	c.adjustListenerState("validateIncomingLinkStateChanged", "interior-listener", c.validator.ValidateIncomingLink)
}

func (c *PolicyController) adjustListenerState(source string, listenerName string, validationFunc func() *client.PolicyValidationResult) {
	// Retrieving listener info
	configmap, err := kube.GetConfigMap(types.TransportConfigMapName, c.cli.GetNamespace(), c.cli.KubeClient)
	if err != nil {
		event.Recordf(c.name, "[%s] Unable to read %s ConfigMap: %v", source, types.TransportConfigMapName, err)
		return
	}
	current, err := qdr.GetRouterConfigFromConfigMap(configmap)

	// If mode is edge, no need to proceed
	if current.IsEdge() {
		return
	}

	// Retrieving listener info
	listener, ok := current.Listeners[listenerName]
	if !ok {
		event.Recordf(c.name, "[%s] interior-listener not defined: %v", source, err)
		return
	}

	// Validate state
	public := listener.Host == "0.0.0.0"

	// Validating if given policy is allowed
	res := validationFunc()
	if res.Error() != nil {
		event.Recordf(c.name, "[%s] error validating policy: %v", source, err)
		return
	}

	// If nothing changed, just return
	if public == res.Allowed() {
		return
	}

	// Changed to allowed
	if res.Allowed() {
		event.Recordf(c.name, "[%s] allowing links", source)
		listener.Host = "0.0.0.0"
	} else {
		event.Recordf(c.name, "[%s] blocking links", source)
		listener.Host = "127.0.0.1"
	}
	current.AddListener(listener)

	// Update router config
	updated, err := current.UpdateConfigMap(configmap)
	if err != nil {
		event.Recordf(c.name, "[%s] error updating host on %s listener: %v", source, listenerName, err)
		return
	}

	if updated {
		_, err = c.cli.KubeClient.CoreV1().ConfigMaps(c.cli.GetNamespace()).Update(configmap)
		if err != nil {
			event.Recordf(c.name, "[%s] error updating %s ConfigMap: %v", source, configmap.Name, err)
			return
		}
		if err = c.cli.RouterRestart(context.Background(), c.cli.Namespace); err != nil {
			event.Recordf(c.name, "[%s] error restarting router: %v", source, err)
			return
		}
	}
}

func (c *PolicyController) validateOutgoingLinkStateChanged() {
	// Iterate through all links
	links, err := c.cli.ConnectorList(context.Background())
	if err != nil {
		event.Recordf(c.name, "[validateOutgoingLinkStateChanged] error reading existing links: %v", err)
		return
	}
	for _, link := range links {
		// Retrieving state of respective link (enabled/disabled)
		secret, err := c.cli.KubeClient.CoreV1().Secrets(c.cli.GetNamespace()).Get(link.Name, v1.GetOptions{})
		if err != nil {
			event.Recordf(c.name, "[validateOutgoingLinkStateChanged] error reading secret %s: %v", link.Name, err)
			return
		}
		disabledValue, ok := secret.ObjectMeta.Labels[types.SkupperDisabledQualifier]
		disabled := false
		if ok {
			disabled, _ = strconv.ParseBool(disabledValue)
		}
		linkUrl := strings.Split(link.Url, ":")
		hostname := linkUrl[0]

		// Validating if hostname is allowed
		res := c.validator.ValidateOutgoingLink(hostname)
		if res.Error() != nil {
			event.Recordf(c.name, "[validateOutgoingLinkStateChanged] error validating if outgoing link to %s is allowed: %v", hostname, res.Error())
			return
		}

		// Not changed, continue to next link
		if res.Allowed() != disabled {
			continue
		}

		// Rule has changed for the related hostname
		if res.Allowed() {
			event.Recordf(c.name, "[validateOutgoingLinkStateChanged] enabling link %s", link.Name)
			delete(secret.Labels, types.SkupperDisabledQualifier)
		} else {
			event.Recordf(c.name, "[validateOutgoingLinkStateChanged] disabling link %s", link.Name)
			secret.Labels[types.SkupperDisabledQualifier] = "true"
		}

		// Update secret
		_, err = c.cli.KubeClient.CoreV1().Secrets(c.cli.GetNamespace()).Update(secret)
		if err != nil {
			event.Recordf(c.name, "[validateOutgoingLinkStateChanged] error updating secret %s: %v", link.Name, res.Error())
			return
		}
	}
}

func (c *PolicyController) validateGatewayStateChanged() {
	c.adjustListenerState("validateGatewayStateChanged", "edge-listener", c.validator.ValidateCreateGateway)
}

func (c *PolicyController) validateExposeStateChanged() {
	policies, err := c.validator.LoadNamespacePolicies()
	if err != nil {
		event.Recordf(c.name, "[validateExposeStateChanged] error retrieving policies: %v", err)
		return
	}

	for _, policy := range policies {
		// If there is a policy allowing all resources, no need to continue
		if utils.StringSliceContains(policy.Spec.AllowedOutgoingLinksHostnames, "*") {
			return
		}
	}

	serviceList, err := c.cli.ServiceInterfaceList(context.Background())
	if err != nil {
		event.Recordf(c.name, "[validateExposeStateChanged] error retrieving service list: %v", err)
		return
	}

	// iterate through service list and inspect if respective targets are allowed (ignoring target type)
	for _, service := range serviceList {
		if service.Targets == nil || len(service.Targets) == 0 {
			continue
		}
		for _, target := range service.Targets {
			// TODO service interface target does not store target type (validation will not be precise)
			res := c.validator.ValidateExpose("", target.Name)
			if res.Error() != nil {
				event.Recordf(c.name, "[validateExposeStateChanged] error validating if target can still be exposed: %v", err)
				return
			}
			if !res.Allowed() {
				// resource is no longer allowed, unbinding
				event.Recordf(c.name, "[validateExposeStateChanged] exposed resource is no longer authorized - unbinding service %s: %v", service.Address, err)
				err = c.cli.ServiceInterfaceUnbind(context.Background(), "deployment", target.Name, service.Address, false)
				if err != nil {
					event.Recordf(c.name, "[validateExposeStateChanged] error unbinding service %s: %v", service.Address, err)
					return
				}
			}
		}
	}
}

func (c *PolicyController) validateServiceStateChanged() {
	serviceList, err := c.cli.ServiceInterfaceList(context.Background())
	if err != nil {
		event.Recordf(c.name, "[validateServiceStateChanged] error retrieving service list: %v", err)
		return
	}

	for _, service := range serviceList {
		res := c.validator.ValidateImportService(service.Address)
		if res.Error() != nil {
			event.Recordf(c.name, "[validateServiceStateChanged] error validating service policy: %v", res.Error())
			return
		}
		if !res.Allowed() {
			err = c.cli.ServiceInterfaceRemove(context.Background(), service.Address)
			if err != nil {
				event.Recordf(c.name, "[validateServiceStateChanged] error removing service definition %s: %v", service.Address, err)
				return
			}
		}
	}
}

func NewPolicyController(cli *client.VanClient) *PolicyController {
	skupperCli, err := versioned.NewForConfig(cli.RestConfig)
	if err != nil {
		return nil
	}

	informer := v1alpha12.NewSkupperClusterPolicyInformer(
		skupperCli,
		time.Second*30,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "PolicyHandler")

	controller := &PolicyController{
		cli:       cli,
		validator: client.NewClusterPolicyValidator(cli),
		informer:  informer,
		queue:     queue,
	}
	informer.AddEventHandler(controller)

	return controller
}
