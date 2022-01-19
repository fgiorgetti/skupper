package client

import (
	"testing"

	"github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	"github.com/skupperproject/skupper/pkg/generated/client/clientset/versioned/typed/skupper/v1alpha1/fake"
	"github.com/skupperproject/skupper/pkg/utils"
	"gotest.tools/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	testing2 "k8s.io/client-go/testing"
)

// Unit test using mock data

var (
	emptyPolicy = []v1alpha1.SkupperClusterPolicy{
		{
			ObjectMeta: v1.ObjectMeta{
				Name: "policy-empty",
			},
			Spec: v1alpha1.SkupperClusterPolicySpec{},
		},
	}

	allNs = []string{"*"}
	abcNs = []string{"aaa", "bbb", "ccc"}
)

func NewClusterPolicyValidatorMock(ns string, policies []v1alpha1.SkupperClusterPolicy) *ClusterPolicyValidator {
	policyCli := &fake.FakeSkupperV1alpha1{Fake: &testing2.Fake{}}
	policyCli.Fake.ClearActions()
	policyCli.Fake.PrependReactor("list", "skupperclusterpolicies", func(action testing2.Action) (handled bool, ret runtime.Object, err error) {
		return true, &v1alpha1.SkupperClusterPolicyList{
			Items: policies,
		}, nil
	})

	cli, _ := newMockClient(ns, "", "")
	return &ClusterPolicyValidator{
		cli:           cli,
		skupperPolicy: policyCli.SkupperClusterPolicies(),
	}
}

func TestMockValidateIncomingLink(t *testing.T) {
	type tc struct {
		name                string
		ns                  string
		policies            []v1alpha1.SkupperClusterPolicy
		expAllowed          bool
		expAllowPolicyNames []string
	}
	type policyData struct {
		name       string
		namespaces []string
		allow      bool
	}
	addIncomingLinkPolicy := func(data []policyData) []v1alpha1.SkupperClusterPolicy {
		var policies []v1alpha1.SkupperClusterPolicy
		for _, d := range data {
			policies = append(policies, v1alpha1.SkupperClusterPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name: d.name,
				},
				Spec: v1alpha1.SkupperClusterPolicySpec{
					Namespaces:         d.namespaces,
					AllowIncomingLinks: d.allow,
				},
			})
		}
		return policies
	}

	scenarios := []tc{
		{
			name: "allow-all-ns",
			ns:   "aaa",
			policies: addIncomingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"*"}, allow: true},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-1"},
		},
		{
			name: "allow-bbb-ns",
			ns:   "bbb",
			policies: addIncomingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
				{name: "policy-4", namespaces: []string{"*"}, allow: true},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-2", "policy-4"},
		},
		{
			name: "allow-ccc-ns",
			ns:   "ccc",
			policies: addIncomingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-3"},
		},
		{
			name: "deny-ddd-ns",
			ns:   "ddd",
			policies: addIncomingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
		{
			name: "deny-explicit-ddd-ns",
			ns:   "ddd",
			policies: addIncomingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
				{name: "policy-4", namespaces: []string{"ddd"}, allow: false},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			policyMock := NewClusterPolicyValidatorMock(scenario.ns, scenario.policies)
			res := policyMock.ValidateIncomingLink()

			// asserting results
			assert.Equal(t, scenario.expAllowed, res.Allowed())
			assert.Equal(t, len(res.matchingAllowed), len(scenario.expAllowPolicyNames))
			if scenario.expAllowed {
				for _, expPolicyName := range scenario.expAllowPolicyNames {
					assert.Assert(t, utils.StringSliceContains(res.AllowPolicyNames(), expPolicyName))
				}
			}
		})
	}

}

func TestMockValidateOutgoingLink(t *testing.T) {
	type tc struct {
		name                string
		ns                  string
		hostname            string
		policies            []v1alpha1.SkupperClusterPolicy
		expAllowed          bool
		expAllowPolicyNames []string
	}
	type policyData struct {
		name       string
		namespaces []string
		hostnames  []string
	}
	addOutgoingLinkPolicy := func(data []policyData) []v1alpha1.SkupperClusterPolicy {
		var policies []v1alpha1.SkupperClusterPolicy
		for _, d := range data {
			policies = append(policies, v1alpha1.SkupperClusterPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name: d.name,
				},
				Spec: v1alpha1.SkupperClusterPolicySpec{
					Namespaces:                    d.namespaces,
					AllowedOutgoingLinksHostnames: d.hostnames,
				},
			})
		}
		return policies
	}

	scenarios := []tc{
		{
			name:     "allow-all-ns-all-hosts",
			ns:       "aaa",
			hostname: "hostname1.domain1",
			policies: addOutgoingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"*"}, hostnames: []string{"*"}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-1"},
		},
		{
			name:     "allow-ns-and-hostname",
			ns:       "aaa",
			hostname: "hostname1.domain1",
			policies: addOutgoingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, hostnames: []string{"hostname1.domain1"}},
				{name: "policy-2", namespaces: []string{"bbb"}, hostnames: []string{"hostname2.domain2"}},
				{name: "policy-3", namespaces: []string{"ccc"}, hostnames: []string{"hostname3.domain3"}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-1"},
		},
		{
			name:     "allow-ns-and-hostname-regex",
			ns:       "ccc",
			hostname: "hostname3.domain3",
			policies: addOutgoingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, hostnames: []string{"hostname1.domain1"}},
				{name: "policy-2", namespaces: []string{"bbb"}, hostnames: []string{"hostname2.domain2"}},
				{name: "policy-3", namespaces: []string{"ccc"}, hostnames: []string{`hostname[0-9]\.domain[0-9]`}},
				{name: "policy-4", namespaces: []string{"ccc"}, hostnames: []string{`hostname.*`}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-3", "policy-4"},
		},
		{
			name:     "deny-all-ns-no-hosts",
			ns:       "aaa",
			hostname: "hostname1.domain1",
			policies: addOutgoingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"*"}, hostnames: []string{}},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
		{
			name:     "deny-hostname",
			ns:       "bbb",
			hostname: "hostname2.domain2",
			policies: addOutgoingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, hostnames: []string{"hostname1.domain1"}},
				{name: "policy-2", namespaces: []string{"bbb"}, hostnames: []string{"hostname1.domain1"}},
				{name: "policy-3", namespaces: []string{"ccc"}, hostnames: []string{"hostname1.domain1"}},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
		{
			name:     "deny-no-policy-for-ns",
			ns:       "bbb",
			hostname: "hostname2.domain2",
			policies: addOutgoingLinkPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, hostnames: []string{"hostname2.domain2"}},
				{name: "policy-3", namespaces: []string{"ccc"}, hostnames: []string{"hostname2.domain2"}},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			policyMock := NewClusterPolicyValidatorMock(scenario.ns, scenario.policies)
			res := policyMock.ValidateOutgoingLink(scenario.hostname)

			// asserting results
			assert.Equal(t, scenario.expAllowed, res.Allowed())
			assert.Equal(t, len(res.matchingAllowed), len(scenario.expAllowPolicyNames))
			if scenario.expAllowed {
				for _, expPolicyName := range scenario.expAllowPolicyNames {
					assert.Assert(t, utils.StringSliceContains(res.AllowPolicyNames(), expPolicyName))
				}
			}
		})
	}
}

func TestMockValidateExpose(t *testing.T) {
	type tc struct {
		name                string
		ns                  string
		resource            string
		policies            []v1alpha1.SkupperClusterPolicy
		expAllowed          bool
		expAllowPolicyNames []string
	}
	type policyData struct {
		name       string
		namespaces []string
		resources  []string
	}
	addExposePolicy := func(data []policyData) []v1alpha1.SkupperClusterPolicy {
		var policies []v1alpha1.SkupperClusterPolicy
		for _, d := range data {
			policies = append(policies, v1alpha1.SkupperClusterPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name: d.name,
				},
				Spec: v1alpha1.SkupperClusterPolicySpec{
					Namespaces:              d.namespaces,
					AllowedExposedResources: d.resources,
				},
			})
		}
		return policies
	}

	scenarios := []tc{
		{
			name:     "allow-all-ns-all-resources",
			ns:       "aaa",
			resource: "deployment/my-app-1",
			policies: addExposePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"*"}, resources: []string{"*"}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-1"},
		},
		{
			name:     "allow-ns-resource",
			ns:       "aaa",
			resource: "deployment/my-app-3",
			policies: addExposePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-1"}},
				{name: "policy-2", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-2"}},
				{name: "policy-3", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-3"}},
				{name: "policy-4", namespaces: []string{"bbb"}, resources: []string{"*"}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-3"},
		},
		{
			name:     "allow-ns-statefulset-regex",
			ns:       "aaa",
			resource: "statefulset/my-app",
			policies: addExposePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-1"}},
				{name: "policy-2", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-2"}},
				{name: "policy-3", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-3"}},
				{name: "policy-4", namespaces: []string{"aaa"}, resources: []string{`statefulset\/.*`}},
				{name: "policy-5", namespaces: []string{"bbb"}, resources: []string{"*"}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-4"},
		},
		{
			name:     "deny-no-ns-policy",
			ns:       "ddd",
			resource: "deployment/my-app-1",
			policies: addExposePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-1"}},
				{name: "policy-2", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-2"}},
				{name: "policy-3", namespaces: []string{"bbb"}, resources: []string{"deployment/my-app-3"}},
				{name: "policy-4", namespaces: []string{"bbb"}, resources: []string{`statefulset\/.*`}},
				{name: "policy-5", namespaces: []string{"ccc"}, resources: []string{"*"}},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
		{
			name:     "deny-no-resource-policy",
			ns:       "bbb",
			resource: "deployment/my-app-1",
			policies: addExposePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-1"}},
				{name: "policy-2", namespaces: []string{"aaa"}, resources: []string{"deployment/my-app-2"}},
				{name: "policy-3", namespaces: []string{"bbb"}, resources: []string{"deployment/my-app-3"}},
				{name: "policy-4", namespaces: []string{"bbb"}, resources: []string{`statefulset\/.*`}},
				{name: "policy-5", namespaces: []string{"ccc"}, resources: []string{"*"}},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			policyMock := NewClusterPolicyValidatorMock(scenario.ns, scenario.policies)
			res := policyMock.ValidateExpose(scenario.resource)

			// asserting results
			assert.Equal(t, scenario.expAllowed, res.Allowed())
			assert.Equal(t, len(res.matchingAllowed), len(scenario.expAllowPolicyNames))
			if scenario.expAllowed {
				for _, expPolicyName := range scenario.expAllowPolicyNames {
					assert.Assert(t, utils.StringSliceContains(res.AllowPolicyNames(), expPolicyName))
				}
			}
		})
	}
}

func TestMockValidateImportService(t *testing.T) {
	type tc struct {
		name                string
		ns                  string
		service             string
		policies            []v1alpha1.SkupperClusterPolicy
		expAllowed          bool
		expAllowPolicyNames []string
	}
	type policyData struct {
		name       string
		namespaces []string
		services   []string
	}
	addServicePolicy := func(data []policyData) []v1alpha1.SkupperClusterPolicy {
		var policies []v1alpha1.SkupperClusterPolicy
		for _, d := range data {
			policies = append(policies, v1alpha1.SkupperClusterPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name: d.name,
				},
				Spec: v1alpha1.SkupperClusterPolicySpec{
					Namespaces:      d.namespaces,
					AllowedServices: d.services,
				},
			})
		}
		return policies
	}

	scenarios := []tc{
		{
			name:    "allow-all-ns-all-services",
			ns:      "aaa",
			service: "service-1",
			policies: addServicePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"*"}, services: []string{"*"}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-1"},
		},
		{
			name:    "allow-ns-service",
			ns:      "aaa",
			service: "service-4",
			policies: addServicePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, services: []string{"service-1", "service-2"}},
				{name: "policy-2", namespaces: []string{"aaa"}, services: []string{"service-3", "service-4"}},
				{name: "policy-3", namespaces: []string{"bbb"}, services: []string{"service-1"}},
				{name: "policy-4", namespaces: []string{"bbb"}, services: []string{"service-2"}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-2"},
		},
		{
			name:    "allow-ns-service-regex",
			ns:      "aaa",
			service: "my-service-99",
			policies: addServicePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, services: []string{"service-1"}},
				{name: "policy-2", namespaces: []string{"aaa"}, services: []string{"service-2"}},
				{name: "policy-3", namespaces: []string{"aaa"}, services: []string{`service-[0-9]+`, `my-service-[0-9]+`}},
				{name: "policy-5", namespaces: []string{"bbb"}, services: []string{"*"}},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-3"},
		},
		{
			name:    "deny-no-ns-policy",
			ns:      "ddd",
			service: "service-1",
			policies: addServicePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, services: []string{"service-1"}},
				{name: "policy-2", namespaces: []string{"aaa"}, services: []string{"service-1", "service-2"}},
				{name: "policy-3", namespaces: []string{"bbb"}, services: []string{"service-3"}},
				{name: "policy-4", namespaces: []string{"bbb"}, services: []string{`.*`}},
				{name: "policy-5", namespaces: []string{"ccc"}, services: []string{"*"}},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
		{
			name:    "deny-no-service-policy",
			ns:      "bbb",
			service: "service-1",
			policies: addServicePolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, services: []string{"service-1"}},
				{name: "policy-2", namespaces: []string{"aaa"}, services: []string{"service-2"}},
				{name: "policy-3", namespaces: []string{"bbb"}, services: []string{"service-3"}},
				{name: "policy-4", namespaces: []string{"bbb"}, services: []string{`my-service-[0-9]+`}},
				{name: "policy-5", namespaces: []string{"ccc"}, services: []string{"*"}},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			policyMock := NewClusterPolicyValidatorMock(scenario.ns, scenario.policies)
			res := policyMock.ValidateImportService(scenario.service)

			// asserting results
			assert.Equal(t, scenario.expAllowed, res.Allowed())
			assert.Equal(t, len(res.matchingAllowed), len(scenario.expAllowPolicyNames))
			if scenario.expAllowed {
				for _, expPolicyName := range scenario.expAllowPolicyNames {
					assert.Assert(t, utils.StringSliceContains(res.AllowPolicyNames(), expPolicyName))
				}
			}
		})
	}
}

func TestMockValidateCreateGateway(t *testing.T) {
	type tc struct {
		name                string
		ns                  string
		policies            []v1alpha1.SkupperClusterPolicy
		expAllowed          bool
		expAllowPolicyNames []string
	}
	type policyData struct {
		name       string
		namespaces []string
		allow      bool
	}
	addGatewayPolicy := func(data []policyData) []v1alpha1.SkupperClusterPolicy {
		var policies []v1alpha1.SkupperClusterPolicy
		for _, d := range data {
			policies = append(policies, v1alpha1.SkupperClusterPolicy{
				ObjectMeta: v1.ObjectMeta{
					Name: d.name,
				},
				Spec: v1alpha1.SkupperClusterPolicySpec{
					Namespaces:   d.namespaces,
					AllowGateway: d.allow,
				},
			})
		}
		return policies
	}

	scenarios := []tc{
		{
			name: "allow-all-ns",
			ns:   "aaa",
			policies: addGatewayPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"*"}, allow: true},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-1"},
		},
		{
			name: "allow-bbb-ns",
			ns:   "bbb",
			policies: addGatewayPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
				{name: "policy-4", namespaces: []string{"*"}, allow: true},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-2", "policy-4"},
		},
		{
			name: "allow-ccc-ns",
			ns:   "ccc",
			policies: addGatewayPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-3"},
		},
		{
			name: "allow-ddd-ns",
			ns:   "ccc",
			policies: addGatewayPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa", "bbb", "ccc", "ddd", "eee"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
				{name: "policy-4", namespaces: []string{"*"}, allow: true},
			}),
			expAllowed:          true,
			expAllowPolicyNames: []string{"policy-1", "policy-3", "policy-4"},
		},
		{
			name: "deny-ddd-ns",
			ns:   "ddd",
			policies: addGatewayPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
		{
			name: "deny-explicit-ddd-ns",
			ns:   "ddd",
			policies: addGatewayPolicy([]policyData{
				{name: "policy-1", namespaces: []string{"aaa"}, allow: true},
				{name: "policy-2", namespaces: []string{"bbb"}, allow: true},
				{name: "policy-3", namespaces: []string{"ccc"}, allow: true},
				{name: "policy-4", namespaces: []string{"ddd"}, allow: false},
			}),
			expAllowed:          false,
			expAllowPolicyNames: []string{},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			policyMock := NewClusterPolicyValidatorMock(scenario.ns, scenario.policies)
			res := policyMock.ValidateCreateGateway()

			// asserting results
			assert.Equal(t, scenario.expAllowed, res.Allowed())
			assert.Equal(t, len(res.matchingAllowed), len(scenario.expAllowPolicyNames))
			if scenario.expAllowed {
				for _, expPolicyName := range scenario.expAllowPolicyNames {
					assert.Assert(t, utils.StringSliceContains(res.AllowPolicyNames(), expPolicyName))
				}
			}
		})
	}

}
