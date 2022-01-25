package main

import (
	"fmt"
	"time"

	"github.com/skupperproject/skupper/client"
	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/generated/client/clientset/versioned"
	v1alpha12 "github.com/skupperproject/skupper/pkg/generated/client/informers/externalversions/skupper/v1alpha1"
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
		fmt.Println("Policy has changed:", key)
		// TODO implement how to react for policy changes
		fmt.Println("Validate incomingLink state changed")
		fmt.Println("Validate outgoingLink state changed (go through tokens)")
		fmt.Println("Validate gateway state changed")
		fmt.Println("Validate expose state changed (load namespace policies and iterate through allowed resources - unique list - then validate if current targets [kube.GetServiceInterfaceTarget] differ from authorized kube.GetServiceInterfaceTarget")
		// c.validator.LoadNamespacePolicies()
		fmt.Println("Validate allowed services then iterate through all of them and if a local service with the given name exists, remove it")
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
