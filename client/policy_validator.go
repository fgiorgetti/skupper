package client

import (
	"strings"

	v1alpha12 "github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/generated/client/clientset/versioned/typed/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PolicyValidationResult struct {
	err             error
	matchingAllowed []v1alpha12.SkupperClusterPolicy
}

func (p *PolicyValidationResult) Enabled() bool {
	crdAvailable := p.err == nil || !strings.Contains(p.err.Error(), "the server could not find the requested resource")
	permissionGranted := p.err == nil || !strings.Contains(p.err.Error(), "is forbidden")
	return crdAvailable && permissionGranted
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

type ClusterPolicyValidator struct {
	cli           *VanClient
	skupperPolicy v1alpha1.SkupperClusterPolicyInterface
}

func NewClusterPolicyValidator(cli *VanClient) *ClusterPolicyValidator {
	return &ClusterPolicyValidator{cli: cli}
}

func (p *PolicyValidationResult) addMatchingPolicy(policy v1alpha12.SkupperClusterPolicy) {
	p.matchingAllowed = append(p.matchingAllowed, policy)
}

func (p *ClusterPolicyValidator) loadNamespacePolicies() ([]v1alpha12.SkupperClusterPolicy, error) {
	policies := []v1alpha12.SkupperClusterPolicy{}
	if p.skupperPolicy == nil {
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
		if utils.StringSliceContains(pol.Spec.Namespaces, "*") || utils.StringSliceContains(pol.Spec.Namespaces, p.cli.Namespace) {
			policies = append(policies, pol)
		}
	}
	return policies, nil
}

func (p *ClusterPolicyValidator) Enabled() bool {
	_, err := p.loadNamespacePolicies()
	if err != nil && strings.Contains(err.Error(), "the server could not find the requested resource") {
		return false
	}
	return true
}

func (p *ClusterPolicyValidator) ValidateIncomingLink() *PolicyValidationResult {
	policies, err := p.loadNamespacePolicies()
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
	policies, err := p.loadNamespacePolicies()
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
	policies, err := p.loadNamespacePolicies()
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
		} else if utils.RegexpStringSliceContains(pol.Spec.AllowedExposedResources, resource) {
			res.addMatchingPolicy(pol)
		}
	}

	return res
}

func (p *ClusterPolicyValidator) ValidateImportService(serviceName string) *PolicyValidationResult {
	policies, err := p.loadNamespacePolicies()
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
	policies, err := p.loadNamespacePolicies()
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
