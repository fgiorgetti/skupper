package client

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	v1alpha12 "github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/event"
	"github.com/skupperproject/skupper/pkg/generated/client/clientset/versioned/typed/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"
)

type PolicyValidationResult struct {
	err             error
	matchingAllowed []v1alpha12.SkupperClusterPolicy
}

func (p *PolicyValidationResult) Enabled() bool {
	restCfgAvail := p.err == nil || !strings.Contains(p.err.Error(), "RestConfig not defined")
	crdAvailable := p.err == nil || !strings.Contains(p.err.Error(), "the server could not find the requested resource")
	permissionGranted := p.err == nil || !strings.Contains(p.err.Error(), "is forbidden")
	return restCfgAvail && crdAvailable && permissionGranted
}

func (p *PolicyValidationResult) Allowed() bool {
	return !p.Enabled() || p.err == nil && len(p.matchingAllowed) > 0
}

func (p *PolicyValidationResult) AllowPolicies() []v1alpha12.SkupperClusterPolicy {
	return p.matchingAllowed
}

func (p *PolicyValidationResult) AllowPolicyNames() []string {
	var names []string
	for _, p := range p.matchingAllowed {
		names = append(names, p.Name)
	}
	return names
}

func (p *PolicyValidationResult) Error() error {
	return p.err
}

// ClusterPolicyValidator The policy validator component must be
// used internally by the service-controller only. Client applications
// must use the PolicyAPIClient (rest client).
type ClusterPolicyValidator struct {
	cli           *VanClient
	skupperPolicy v1alpha1.SkupperClusterPolicyInterface
	labelRegex    *regexp.Regexp
}

func NewClusterPolicyValidator(cli *VanClient) *ClusterPolicyValidator {
	return &ClusterPolicyValidator{
		cli:        cli,
		labelRegex: regexp.MustCompile(ValidRfc1123Label),
	}
}

func (p *PolicyValidationResult) addMatchingPolicy(policy v1alpha12.SkupperClusterPolicy) {
	p.matchingAllowed = append(p.matchingAllowed, policy)
}

func (p *ClusterPolicyValidator) LoadNamespacePolicies() ([]v1alpha12.SkupperClusterPolicy, error) {
	policies := []v1alpha12.SkupperClusterPolicy{}
	if p.skupperPolicy == nil {
		if p.cli.RestConfig == nil {
			return policies, fmt.Errorf("RestConfig not defined")
		}
		skupperCli, err := v1alpha1.NewForConfig(p.cli.RestConfig)
		if err != nil {
			return policies, err
		}
		p.skupperPolicy = skupperCli.SkupperClusterPolicies()
	}
	policyList, err := p.skupperPolicy.List(v1.ListOptions{})
	if err != nil {
		return policies, err
	}
	for _, pol := range policyList.Items {
		if p.appliesToNS(&pol) {
			policies = append(policies, pol)
		}
	}
	return policies, nil
}

func (p *ClusterPolicyValidator) AppliesToNS(policyName string) bool {
	pol, err := p.skupperPolicy.Get(policyName, v1.GetOptions{})
	// If policy not found, revalidate
	if err != nil {
		return true
	}
	return p.appliesToNS(pol)
}

func (p *ClusterPolicyValidator) appliesToNS(pol *v1alpha12.SkupperClusterPolicy) bool {
	var namespaces []string
	for _, ns := range pol.Spec.Namespaces {
		if p.labelRegex.MatchString(ns) {
			namespace, err := p.cli.KubeClient.CoreV1().Namespaces().Get(p.cli.Namespace, v1.GetOptions{})
			if err == nil {
				selector, _ := labels.Parse(ns)
				if selector.Matches(labels.Set(namespace.Labels)) {
					namespaces = append(namespaces, namespace.Name)
				}
			}
		} else {
			namespaces = append(namespaces, ns)
		}
	}
	return utils.StringSliceContains(namespaces, "*") ||
		utils.StringSliceContains(namespaces, p.cli.Namespace) ||
		utils.RegexpStringSliceContains(namespaces, p.cli.Namespace)
}

func (p *ClusterPolicyValidator) Enabled() bool {
	if p.cli.RestConfig == nil {
		return false
	}
	_, err := p.LoadNamespacePolicies()
	if err != nil && (!p.CrdDefined(err) || !p.NoPermission(err)) {
		return false
	}
	return true
}

func (p *ClusterPolicyValidator) NoPermission(err error) bool {
	if err == nil {
		return true
	}
	return !strings.Contains(err.Error(), "is forbidden")
}

func (p *ClusterPolicyValidator) CrdDefined(err error) bool {
	if err == nil {
		return true
	}
	return !strings.Contains(err.Error(), "the server could not find the requested resource")
}

func (p *ClusterPolicyValidator) ValidateIncomingLink() *PolicyValidationResult {
	policies, err := p.LoadNamespacePolicies()
	res := &PolicyValidationResult{
		err: err,
	}
	if err != nil || len(policies) == 0 {
		return res
	}

	for _, pol := range policies {
		if pol.Spec.AllowIncomingLinks {
			res.addMatchingPolicy(pol)
		}
	}

	return res
}

func (p *ClusterPolicyValidator) ValidateOutgoingLink(hostname string) *PolicyValidationResult {
	policies, err := p.LoadNamespacePolicies()
	res := &PolicyValidationResult{
		err: err,
	}
	if err != nil || len(policies) == 0 {
		return res
	}

	for _, pol := range policies {
		if utils.StringSliceContains(pol.Spec.AllowedOutgoingLinksHostnames, "*") {
			res.addMatchingPolicy(pol)
		} else if utils.StringSliceContains(pol.Spec.AllowedOutgoingLinksHostnames, hostname) {
			res.addMatchingPolicy(pol)
		} else if utils.RegexpStringSliceContains(pol.Spec.AllowedOutgoingLinksHostnames, hostname) {
			res.addMatchingPolicy(pol)
		}
	}

	return res
}

func (p *ClusterPolicyValidator) ValidateExpose(resourceType, resourceName string) *PolicyValidationResult {
	policies, err := p.LoadNamespacePolicies()
	res := &PolicyValidationResult{
		err: err,
	}
	if err != nil || len(policies) == 0 {
		return res
	}

	resource := resourceType + "/" + resourceName
	for _, pol := range policies {
		if utils.StringSliceContains(pol.Spec.AllowedExposedResources, "*") {
			res.addMatchingPolicy(pol)
		} else if utils.StringSliceContains(pol.Spec.AllowedExposedResources, resource) {
			res.addMatchingPolicy(pol)
		} else if resourceType == "" && utils.StringSliceEndsWith(pol.Spec.AllowedExposedResources, resource) {
			res.addMatchingPolicy(pol)
		}
	}

	return res
}

func (p *ClusterPolicyValidator) ValidateImportService(serviceName string) *PolicyValidationResult {
	policies, err := p.LoadNamespacePolicies()
	res := &PolicyValidationResult{
		err: err,
	}
	if err != nil || len(policies) == 0 {
		return res
	}

	for _, pol := range policies {
		if utils.StringSliceContains(pol.Spec.AllowedServices, "*") {
			res.addMatchingPolicy(pol)
		} else if utils.StringSliceContains(pol.Spec.AllowedServices, serviceName) {
			res.addMatchingPolicy(pol)
		} else if utils.RegexpStringSliceContains(pol.Spec.AllowedServices, serviceName) {
			res.addMatchingPolicy(pol)
		}
	}

	return res
}

func (p *ClusterPolicyValidator) ValidateCreateGateway() *PolicyValidationResult {
	policies, err := p.LoadNamespacePolicies()
	res := &PolicyValidationResult{
		err: err,
	}
	if err != nil || len(policies) == 0 {
		return res
	}

	for _, pol := range policies {
		if pol.Spec.AllowGateway {
			res.addMatchingPolicy(pol)
		}
	}

	return res
}

type PolicyAPIClient struct {
	cli *VanClient
}

type PolicyAPIResult struct {
	Allowed   bool     `json:"allowed"`
	AllowedBy []string `json:"allowedBy"`
	Enabled   bool     `json:"enabled"`
	Error     string   `json:"error"`
}

func NewPolicyValidatorAPI(cli *VanClient) *PolicyAPIClient {
	return &PolicyAPIClient{
		cli: cli,
	}
}

func (p *PolicyAPIClient) execGet(args ...string) (*PolicyAPIResult, error) {
	if _, mock := p.cli.KubeClient.(*fake.Clientset); mock {
		return &PolicyAPIResult{
			Allowed: true,
			Enabled: false,
		}, nil
	}
	ctx, cn := context.WithTimeout(context.Background(), time.Second*30)
	defer cn()
	err := utils.RetryWithContext(ctx, time.Millisecond*100, func() (bool, error) {
		_, err := p.cli.exec([]string{"get", "policies", "-h"}, p.cli.GetNamespace())
		if err != nil {
			if err.Error() == "Not ready" {
				return false, nil
			}
			return true, err
		}
		return true, nil
	})
	if err != nil {
		err := fmt.Errorf("Unable to communicate with the API: %v", err)
		if event.DefaultStore != nil {
			event.Recordf("PolicyAPIError", err.Error())
		}
		return &PolicyAPIResult{
			Allowed: false,
			Enabled: false,
		}, err
	}
	fullArgs := []string{"get", "policies"}
	fullArgs = append(fullArgs, args...)
	fullArgs = append(fullArgs, "-o", "json")
	out, err := p.cli.exec(fullArgs, p.cli.GetNamespace())
	if err != nil {
		return nil, err
	}
	res := &PolicyAPIResult{}
	err = json.Unmarshal(out.Bytes(), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PolicyAPIClient) Gateway() (*PolicyAPIResult, error) {
	return p.execGet("gateway")
}

func (p *PolicyAPIClient) Expose(resourceType, resourceName string) (*PolicyAPIResult, error) {
	return p.execGet("expose", resourceType, resourceName)
}

func (p *PolicyAPIClient) Service(name string) (*PolicyAPIResult, error) {
	return p.execGet("service", name)
}

func (p *PolicyAPIClient) IncomingLink() (*PolicyAPIResult, error) {
	return p.execGet("incominglink")
}

func (p *PolicyAPIClient) OutgoingLink(hostname string) (*PolicyAPIResult, error) {
	return p.execGet("outgoinglink", hostname)
}
