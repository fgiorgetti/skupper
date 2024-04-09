package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/openshift/api/console/v1alpha1"
	console "github.com/openshift/client-go/console/clientset/versioned/typed/console/v1alpha1"
	"github.com/skupperproject/skupper/pkg/version"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	corev1informer "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/informers/internalinterfaces"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client"
	"github.com/skupperproject/skupper/pkg/kube"
)

var (
	//go:embed plugin-nginx.conf
	nginxConf string
)

type SiteController struct {
	vanClient            *client.VanClient
	siteInformer         cache.SharedIndexInformer
	tokenInformer        cache.SharedIndexInformer
	tokenRequestInformer cache.SharedIndexInformer
	workqueue            workqueue.RateLimitingInterface
}

func NewSiteController(cli *client.VanClient) (*SiteController, error) {
	var watchNamespace string

	// Startup message
	if os.Getenv("WATCH_NAMESPACE") != "" {
		watchNamespace = os.Getenv("WATCH_NAMESPACE")
		log.Println("Skupper site controller watching current namespace ", watchNamespace)
	} else {
		watchNamespace = metav1.NamespaceAll
		log.Println("Skupper site controller watching all namespaces")
	}
	log.Printf("Version: %s", version.Version)

	siteInformer := corev1informer.NewFilteredConfigMapInformer(
		cli.KubeClient,
		watchNamespace,
		time.Second*30,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		internalinterfaces.TweakListOptionsFunc(func(options *metav1.ListOptions) {
			options.FieldSelector = "metadata.name=skupper-site"
			options.LabelSelector = "!" + types.SiteControllerIgnore
		}))
	tokenInformer := corev1informer.NewFilteredSecretInformer(
		cli.KubeClient,
		watchNamespace,
		time.Second*30,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		internalinterfaces.TweakListOptionsFunc(func(options *metav1.ListOptions) {
			options.LabelSelector = types.TypeTokenQualifier
		}))
	tokenRequestInformer := corev1informer.NewFilteredSecretInformer(
		cli.KubeClient,
		watchNamespace,
		time.Second*30,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		internalinterfaces.TweakListOptionsFunc(func(options *metav1.ListOptions) {
			options.LabelSelector = types.TypeTokenRequestQualifier
		}))
	workqueue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "skupper-site-controller")

	controller := &SiteController{
		vanClient:            cli,
		siteInformer:         siteInformer,
		tokenInformer:        tokenInformer,
		tokenRequestInformer: tokenRequestInformer,
		workqueue:            workqueue,
	}

	siteInformer.AddEventHandler(controller.getHandlerFuncs(SiteConfig, configmapResourceVersionTest))
	tokenInformer.AddEventHandler(controller.getHandlerFuncs(Token, secretResourceVersionTest))
	tokenRequestInformer.AddEventHandler(controller.getHandlerFuncs(TokenRequest, secretResourceVersionTest))
	return controller, nil
}

type resourceVersionTest func(a interface{}, b interface{}) bool

func configmapResourceVersionTest(a interface{}, b interface{}) bool {
	aa := a.(*corev1.ConfigMap)
	bb := b.(*corev1.ConfigMap)
	return aa.ResourceVersion == bb.ResourceVersion
}

func secretResourceVersionTest(a interface{}, b interface{}) bool {
	aa := a.(*corev1.Secret)
	bb := b.(*corev1.Secret)
	return aa.ResourceVersion == bb.ResourceVersion
}

func (c *SiteController) getHandlerFuncs(category triggerType, test resourceVersionTest) *cache.ResourceEventHandlerFuncs {
	return &cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.enqueueTrigger(obj, category)
		},
		UpdateFunc: func(old, new interface{}) {
			if !test(old, new) {
				c.enqueueTrigger(new, category)
			}
		},
		DeleteFunc: func(obj interface{}) {
			c.enqueueTrigger(obj, category)
		},
	}
}

func (c *SiteController) updateControllerClusterRoles() {
	old := "skupper-service-controller"
	if deleted, err := kube.DeleteClusterRole(old, c.vanClient.KubeClient); deleted || err != nil {
		if err != nil {
			log.Printf("Unable to delete old cluster role %q: %s", old, err)
		} else {
			log.Printf("Deleted old cluster role %q", old)
		}
	}
	for _, clusterRole := range c.vanClient.ClusterRoles(isClusterPermissionAllowed()) {
		_, err := kube.CreateClusterRole(clusterRole, c.vanClient.KubeClient)
		if errors.IsAlreadyExists(err) {
			log.Printf("Cluster role %q already exists", clusterRole.Name)
		} else if err != nil {
			log.Printf("Unable to create new cluster role %q: %s", clusterRole.Name, err)
		} else {
			log.Printf("Cluster role %q created", clusterRole.Name)
		}
	}
}

func (c *SiteController) Run(stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	c.updateControllerClusterRoles()
	c.activateConsolePlugin()

	log.Println("Starting the Skupper site controller informers")
	go c.siteInformer.Run(stopCh)
	go c.tokenInformer.Run(stopCh)
	go c.tokenRequestInformer.Run(stopCh)

	log.Println("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.siteInformer.HasSynced); !ok {
		return fmt.Errorf("Failed to wait for caches to sync")
	}
	log.Printf("Checking if sites need updates (%s)", version.Version)
	c.updateChecks()
	log.Println("Starting workers")
	go wait.Until(c.run, time.Second, stopCh)
	log.Println("Started workers")

	<-stopCh
	log.Println("Shutting down workers")
	return nil
}

type triggerType int

const (
	SiteConfig triggerType = iota
	Token
	TokenRequest
)

type trigger struct {
	key      string
	category triggerType
}

func (c *SiteController) run() {
	for c.processNextTrigger() {
	}
}

func (c *SiteController) processNextTrigger() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	defer c.workqueue.Done(obj)
	var t trigger
	var ok bool
	if t, ok = obj.(trigger); !ok {
		// invalid item
		c.workqueue.Forget(obj)
		utilruntime.HandleError(fmt.Errorf("Invalid item on work queue %#v", obj))
		return true
	}

	err := c.dispatchTrigger(t)
	c.workqueue.Forget(obj)
	if err != nil {
		utilruntime.HandleError(err)
	}

	return true
}

func (c *SiteController) dispatchTrigger(trigger trigger) error {
	switch trigger.category {
	case SiteConfig:
		return c.checkSite(trigger.key)
	case Token:
		return c.checkToken(trigger.key)
	case TokenRequest:
		return c.checkTokenRequest(trigger.key)
	default:
		return fmt.Errorf("invalid trigger %d", trigger.category)
	}

}

func (c *SiteController) enqueueTrigger(obj interface{}, category triggerType) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(trigger{
		key:      key,
		category: category,
	})
}

func (c *SiteController) checkAllForSite() {
	// Now need to check whether there are any token requests already in place
	log.Println("Checking token requests...")
	c.checkAllToken()
	c.checkAllTokenRequests()
	log.Println("Done.")
}

const clusterPermissionEnvVarName string = "CLUSTER_PERMISSIONS_ALLOWED"

func (c *SiteController) checkSite(key string) error {
	// get site namespace
	siteNamespace, _, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		log.Println("Error checking skupper-site namespace: ", err)
		return err
	}
	// get skupper-site configmap
	obj, exists, err := c.siteInformer.GetStore().GetByKey(key)
	if err != nil {
		log.Println("Error checking skupper-site config map: ", err)
		return err
	} else if exists {
		configmap := obj.(*corev1.ConfigMap)
		_, err := c.vanClient.RouterInspectNamespace(context.Background(), configmap.ObjectMeta.Namespace)
		if err == nil {
			log.Println("Skupper site exists", key)
			updatedLogging, err := c.vanClient.RouterUpdateLogging(context.Background(), configmap, false)
			if err != nil {
				log.Println("Error checking router logging configuration:", err)
			}
			if updatedLogging {
				err = c.vanClient.RouterRestart(context.Background(), configmap.ObjectMeta.Namespace)
				if err != nil {
					log.Println("Error restarting router:", err)
				} else {
					log.Println("Updated router logging for", key)
				}
			}
			updatedAnnotations, err := c.vanClient.RouterUpdateAnnotations(context.Background(), configmap)
			if err != nil {
				log.Println("Error checking annotations:", err)
			} else if updatedAnnotations {
				log.Println("Updated annotations for", key)
			}

			c.checkAllForSite()
		} else if errors.IsNotFound(err) {
			log.Println("Initialising skupper site ...")
			siteConfig, _ := c.vanClient.SiteConfigInspect(context.Background(), configmap)
			siteConfig.Spec.SkupperNamespace = siteNamespace
			if siteConfig.Spec.EnableClusterPermissions && !isClusterPermissionAllowed() {
				siteConfig.Spec.EnableClusterPermissions = false
				log.Printf("Ignoring request for cluster permissions for %q; site will not be able to target workloads in other namespaces.", siteNamespace)
				if os.Getenv(clusterPermissionEnvVarName) == "" {
					log.Printf("To enable cluster permissions, set env var %s to 'true'", clusterPermissionEnvVarName)
				}
			}
			ctx, cancel := context.WithTimeout(context.Background(), types.DefaultTimeoutDuration)
			defer cancel()
			err = c.vanClient.RouterCreate(ctx, *siteConfig)
			if err != nil {
				log.Println("Error initialising skupper: ", err)
				return err
			} else {
				log.Println("Skupper site initialised")
				c.checkAllForSite()
			}
		} else {
			log.Println("Error inspecting VAN router: ", err)
			return err
		}
	}
	return nil
}

func isClusterPermissionAllowed() bool {
	enabled, _ := strconv.ParseBool(os.Getenv(clusterPermissionEnvVarName))
	return enabled
}

func getTokenCost(token *corev1.Secret) (int32, bool) {
	if token.ObjectMeta.Annotations == nil {
		return 0, false
	}
	if costString, ok := token.ObjectMeta.Annotations[types.TokenCost]; ok {
		cost, err := strconv.Atoi(costString)
		if err != nil {
			log.Printf("Ignoring invalid cost annotation %q", costString)
			return 0, false
		}
		return int32(cost), true
	}
	return 0, false
}

func (c *SiteController) connect(token *corev1.Secret, namespace string) error {
	log.Printf("Connecting site in %s using token %s", namespace, token.ObjectMeta.Name)
	var options types.ConnectorCreateOptions
	options.Name = token.ObjectMeta.Name
	options.SkupperNamespace = namespace
	if cost, ok := getTokenCost(token); ok {
		options.Cost = cost
	}
	return c.vanClient.ConnectorCreate(context.Background(), token, options)
}

func (c *SiteController) disconnect(name string, namespace string) error {
	log.Printf("Disconnecting connector %s from site in %s", name, namespace)
	var options types.ConnectorRemoveOptions
	options.Name = name
	options.SkupperNamespace = namespace
	// Secret has already been deleted so force update to current active secrets
	options.ForceCurrent = true
	return c.vanClient.ConnectorRemove(context.Background(), options)
}

func (c *SiteController) generate(token *corev1.Secret) error {
	log.Printf("Generating token for request %s...", token.ObjectMeta.Name)
	generated, _, err := c.vanClient.ConnectorTokenCreate(context.Background(), token.ObjectMeta.Name, token.ObjectMeta.Namespace)
	if err == nil {
		token.Data = generated.Data
		if token.ObjectMeta.Annotations == nil {
			token.ObjectMeta.Annotations = make(map[string]string)
		}
		for key, value := range generated.ObjectMeta.Annotations {
			token.ObjectMeta.Annotations[key] = value
		}
		token.ObjectMeta.Labels[types.SkupperTypeQualifier] = types.TypeToken
		siteId := c.getSiteIdForNamespace(token.ObjectMeta.Namespace)
		if siteId != "" {
			token.ObjectMeta.Annotations[types.TokenGeneratedBy] = siteId
		}
		_, err = c.vanClient.KubeClient.CoreV1().Secrets(token.ObjectMeta.Namespace).Update(context.TODO(), token, metav1.UpdateOptions{})
		return err
	} else {
		log.Printf("Failed to generate token for request %s: %s", token.ObjectMeta.Name, err)
		return err
	}
}

func (c *SiteController) checkAllToken() {
	// can we rely on the cache here?
	tokens := c.tokenInformer.GetStore().List()
	for _, t := range tokens {
		// service from workqueue
		c.enqueueTrigger(t, Token)
	}
}

func (c *SiteController) checkAllTokenRequests() {
	// can we rely on the cache here?
	tokens := c.tokenRequestInformer.GetStore().List()
	for _, t := range tokens {
		// service from workqueue
		c.enqueueTrigger(t, TokenRequest)
	}
}

func (c *SiteController) checkToken(key string) error {
	log.Printf("Handling token for %s", key)
	obj, exists, err := c.tokenInformer.GetStore().GetByKey(key)
	if err != nil {
		log.Println("Error checking connection-token secret: ", err)
		return err
	}
	if !exists {
		return nil
	}

	token := obj.(*corev1.Secret)
	if !c.isTokenValidInSite(token) {
		log.Println("Cannot handle token, as site not yet initialised")
		return nil
	}

	updated, _ := c.vanClient.RouterUpdateHostAliases(context.Background(), token)
	if updated {
		log.Printf("Router updated and restarted due to changes in token %s", key)
	} else {
		log.Printf("Changes in token %s didn't trigger router update", key)
	}
	return nil
}

func (c *SiteController) checkTokenRequest(key string) error {
	log.Printf("Handling token request for %s", key)
	obj, exists, err := c.tokenRequestInformer.GetStore().GetByKey(key)
	if err != nil {
		log.Println("Error checking connection-token-request secret: ", err)
		return err
	} else if exists {
		token := obj.(*corev1.Secret)
		if !c.isTokenRequestValidInSite(token) {
			log.Println("Cannot handle token request, as site not yet initialised")
			return nil
		}
		return c.generate(token)
	}
	return nil
}

func (c *SiteController) getSiteIdForNamespace(namespace string) string {
	cm, err := c.vanClient.KubeClient.CoreV1().ConfigMaps(namespace).Get(context.TODO(), types.SiteConfigMapName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Printf("Could not obtain siteid for namespace %q, assuming not yet initialised", namespace)
		} else {
			log.Printf("Error checking siteid for namespace %q: %s", namespace, err)
		}
		return ""
	}
	return string(cm.ObjectMeta.UID)
}

func (c *SiteController) isTokenValidInSite(token *corev1.Secret) bool {
	siteId := c.getSiteIdForNamespace(token.ObjectMeta.Namespace)
	if author, ok := token.ObjectMeta.Annotations[types.TokenGeneratedBy]; ok && author == siteId {
		// token was generated by this site so should not be applied
		return false
	} else {
		return true
	}
}

func (c *SiteController) isTokenRequestValidInSite(token *corev1.Secret) bool {
	siteId := c.getSiteIdForNamespace(token.ObjectMeta.Namespace)
	if siteId == "" {
		return false
	}
	return true
}

func (c *SiteController) updateChecks() {
	sites := c.siteInformer.GetStore().List()
	for _, s := range sites {
		if site, ok := s.(*corev1.ConfigMap); ok {
			updated, err := c.vanClient.RouterUpdateVersionInNamespace(context.Background(), false, site.ObjectMeta.Namespace)
			if err != nil {
				log.Printf("Version update check failed for namespace %q: %s", site.ObjectMeta.Namespace, err)
			} else if updated {
				log.Printf("Updated version for namespace %q", site.ObjectMeta.Namespace)
			} else {
				log.Printf("Version update not required for namespace %q", site.ObjectMeta.Namespace)
			}
		} else {
			log.Printf("Unexpected item in site informer store: %v", s)
		}
	}
}

func (c *SiteController) activateConsolePlugin() {
	// Create the SiteConsole resource
	if os.Getenv("WATCH_NAMESPACE") != "" || c.vanClient.RouteClient == nil {
		return
	}
	// Reading the current namespace
	namespaceBytes, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		log.Printf("unable read namespace: %s", err)
		return
	}
	namespace := string(namespaceBytes)
	// Reading the owner reference
	siteCtrlDeploy, err := c.vanClient.KubeClient.AppsV1().Deployments(namespace).Get(context.Background(), "skupper-site-controller", metav1.GetOptions{})
	if err != nil {
		log.Printf("Failed to retrieve site controller deployment: %s", err)
		log.Println("Unable to activate the console plugin")
		return
	}
	var ownerReference = metav1.OwnerReference{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Name:       siteCtrlDeploy.Name,
		UID:        siteCtrlDeploy.UID,
	}
	c.createConsolePluginConfig(namespace, ownerReference)
	c.createConsolePluginService(namespace, ownerReference)
	c.createConsolePluginDeployment(namespace, ownerReference)
	c.createConsolePlugin(namespace, ownerReference)
}

func (c *SiteController) createConsolePlugin(namespace string, reference metav1.OwnerReference) {
	//TODO Validate if ConsolePlugin CRD is available
	consoleCli, err := console.NewForConfig(c.vanClient.RestConfig)
	if err != nil {
		log.Printf("unable to manage console plugins: %s", err)
		return
	}

	consolePlugin := &v1alpha1.ConsolePlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: "skupper-site-console",
			OwnerReferences: []metav1.OwnerReference{
				reference,
			},
		},
		Spec: v1alpha1.ConsolePluginSpec{
			DisplayName: "Skupper site console",
			Service: v1alpha1.ConsolePluginService{
				Name:      "skupper-site-console",
				Namespace: namespace,
				Port:      9443,
				BasePath:  "/",
			},
		},
	}
	_, err = consoleCli.ConsolePlugins().Create(context.Background(), consolePlugin, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		log.Printf("unable to create console plugin: %s", err)
	}
}

func (c *SiteController) createConsolePluginDeployment(namespace string, reference metav1.OwnerReference) {
	int32p := func(i int) *int32 {
		i32 := int32(i)
		return &i32
	}
	boolp := func(b bool) *bool {
		return &b
	}
	consolePluginDeployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "skupper-site-console",
			Labels: map[string]string{
				"app":                                "skupper-site-console",
				"app.kubernetes.io/component":        "site-console",
				"app.kubernetes.io/instance":         "skupper-site-console",
				"app.kubernetes.io/name":             "skupper-site-console",
				"app.kubernetes.io/part-of":          "skupper",
				"app.openshift.io/runtime-namespace": "skupper-site-console",
			},
			OwnerReferences: []metav1.OwnerReference{
				reference,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32p(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "skupper-site-console"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":                         "skupper-site-console",
						"app.kubernetes.io/component": "site-console",
						"app.kubernetes.io/instance":  "skupper-site-console",
						"app.kubernetes.io/name":      "skupper-site-console",
						"app.kubernetes.io/part-of":   "skupper",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "site-console",
							Image: "quay.io/vbartoli/rhsi-plugin",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 9443,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							ImagePullPolicy: corev1.PullAlways,
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "plugin-serving-cert",
									ReadOnly:  true,
									MountPath: "/var/serving-cert",
								},
								{
									Name:      "nginx-conf",
									ReadOnly:  true,
									MountPath: "/etc/nginx/nginx.conf",
									SubPath:   "nginx.conf",
								},
							},
							SecurityContext: &corev1.SecurityContext{
								RunAsNonRoot: boolp(true),
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "plugin-serving-cert",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName:  "skupper-site-console-cert",
									DefaultMode: int32p(420),
								},
							},
						},
						{
							Name: "nginx-conf",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "skupper-site-console-nginx-conf",
									},
									DefaultMode: int32p(420),
								},
							},
						},
						{
							Name: "plugin-conf",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "skupper-site-console-plugin-conf",
									},
									DefaultMode: int32p(420),
								},
							},
						},
					},
					RestartPolicy:      corev1.RestartPolicyAlways,
					DNSPolicy:          corev1.DNSClusterFirst,
					ServiceAccountName: "skupper-site-controller",
				},
			},
		},
	}
	existingDeployment, err := c.vanClient.KubeClient.AppsV1().Deployments(namespace).Get(context.Background(), consolePluginDeployment.ObjectMeta.Name, metav1.GetOptions{})
	if existingDeployment != nil && (err == nil || !errors.IsNotFound(err)) {
		if existingDeployment.Spec.Template.Spec.Containers[0].Image != consolePluginDeployment.Spec.Template.Spec.Containers[0].Image {
			existingDeployment.Spec = consolePluginDeployment.Spec
			_, err = c.vanClient.KubeClient.AppsV1().Deployments(namespace).Update(context.Background(), existingDeployment, metav1.UpdateOptions{})
			if err != nil {
				log.Printf("unable to update site console deployment: %s", err)
			}
		}
	}
	_, err = c.vanClient.KubeClient.AppsV1().Deployments(namespace).Create(context.Background(), consolePluginDeployment, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		log.Printf("unable to create console plugin deployment: %s", err)
		return
	}
}

func (c *SiteController) createConsolePluginService(namespace string, reference metav1.OwnerReference) {
	consolePluginService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "skupper-site-console",
			Annotations: map[string]string{
				"service.alpha.openshift.io/serving-cert-secret-name": "skupper-site-console-cert",
			},
			Labels: map[string]string{
				"app":                         "skupper-site-console",
				"app.kubernetes.io/component": "site-console",
				"app.kubernetes.io/instance":  "skupper-site-console",
				"app.kubernetes.io/name":      "skupper-site-console",
				"app.kubernetes.io/part-of":   "skupper",
			},
			OwnerReferences: []metav1.OwnerReference{
				reference,
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "9443-tcp",
					Protocol:   corev1.ProtocolTCP,
					Port:       9443,
					TargetPort: intstr.IntOrString{IntVal: 9443},
				},
			},
			Selector: map[string]string{
				"app": "skupper-site-console",
			},
			Type:            corev1.ServiceTypeClusterIP,
			SessionAffinity: corev1.ServiceAffinityNone,
		},
	}
	_, err := c.vanClient.KubeClient.CoreV1().Services(namespace).Create(context.Background(), consolePluginService, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		log.Printf("unable to create skupper-site-console service: %s", err)
	}
}

func (c *SiteController) createConsolePluginConfig(namespace string, reference metav1.OwnerReference) {
	labels := map[string]string{
		"app.kubernetes.io/component": "site-console",
		"app.kubernetes.io/instance":  "skupper-site-console",
		"app.kubernetes.io/name":      "skupper-site-console",
		"app.kubernetes.io/part-of":   "skupper",
	}

	nginxConfigMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "skupper-site-console-nginx-conf",
			Labels: labels,
			OwnerReferences: []metav1.OwnerReference{
				reference,
			},
		},
		Data: map[string]string{
			"nginx.conf": nginxConf,
		},
	}
	pluginConfigMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "skupper-site-console-plugin-conf",
			Labels: labels,
		},
	}
	ctx := context.Background()
	_, err := c.vanClient.KubeClient.CoreV1().ConfigMaps(namespace).Create(ctx, nginxConfigMap, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		log.Printf("unable to create nginx configmap: %s", err)
	}
	_, err = c.vanClient.KubeClient.CoreV1().ConfigMaps(namespace).Create(ctx, pluginConfigMap, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		log.Printf("unable to create nginx configmap: %s", err)
	}
}
