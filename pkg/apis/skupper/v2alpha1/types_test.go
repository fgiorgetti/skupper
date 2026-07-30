package v2alpha1

import (
	"errors"
	"testing"

	"gotest.tools/v3/assert"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestRouterAccessRole_GetPort(t *testing.T) {
	t.Run("defaults to 55671 for inter-router", func(t *testing.T) {
		role := RouterAccessRole{Name: "inter-router"}
		assert.Equal(t, role.GetPort(), int32(55671))
	})
	t.Run("defaults to 45671 for edge", func(t *testing.T) {
		role := RouterAccessRole{Name: "edge"}
		assert.Equal(t, role.GetPort(), int32(45671))
	})
	t.Run("returns explicit port when set", func(t *testing.T) {
		role := RouterAccessRole{Name: "inter-router", Port: 12345}
		assert.Equal(t, role.GetPort(), int32(12345))
	})
}

func TestRouterAccess_AllocatePort(t *testing.T) {
	t.Run("allocates a new port", func(t *testing.T) {
		r := &RouterAccess{}
		assert.Assert(t, r.AllocatePort("inter-router", 55671))
		assert.Equal(t, r.GetAllocatedPortForRole("inter-router"), int32(55671))
	})
	t.Run("no change for same port", func(t *testing.T) {
		r := &RouterAccess{}
		r.AllocatePort("inter-router", 55671)
		assert.Assert(t, !r.AllocatePort("inter-router", 55671))
	})
	t.Run("zero port returns false", func(t *testing.T) {
		r := &RouterAccess{}
		assert.Assert(t, !r.AllocatePort("inter-router", 0))
	})
}

func TestRouterAccess_GetAllocatedPorts(t *testing.T) {
	r := &RouterAccess{
		Spec: RouterAccessSpec{
			Roles: []RouterAccessRole{
				{Name: "inter-router", Port: 55671},
				{Name: "edge"},
			},
		},
	}
	r.AllocatePort("edge", 45671)

	ports := r.GetAllocatedPorts()
	assert.Equal(t, len(ports), 1)
	assert.Equal(t, ports[0], int32(45671))
}

func TestRouterAccess_GetUnusedPorts(t *testing.T) {
	t.Run("status role absent from spec", func(t *testing.T) {
		r := &RouterAccess{
			Spec: RouterAccessSpec{
				Roles: []RouterAccessRole{
					{Name: "inter-router"},
				},
			},
			Status: RouterAccessStatus{
				Roles: []RouterAccessRole{
					{Name: "inter-router", Port: 55671},
					{Name: "edge", Port: 45671},
				},
			},
		}
		unused := r.GetUnusedPorts()
		assert.Equal(t, len(unused), 1)
		assert.Equal(t, unused[0], int32(45671))
	})
	t.Run("status role port mismatches explicit spec port", func(t *testing.T) {
		r := &RouterAccess{
			Spec: RouterAccessSpec{
				Roles: []RouterAccessRole{
					{Name: "inter-router", Port: 55671},
				},
			},
			Status: RouterAccessStatus{
				Roles: []RouterAccessRole{
					{Name: "inter-router", Port: 11111},
				},
			},
		}
		unused := r.GetUnusedPorts()
		assert.Equal(t, len(unused), 1)
		assert.Equal(t, unused[0], int32(11111))
	})
	t.Run("status role port same as explicit spec port", func(t *testing.T) {
		r := &RouterAccess{
			Spec: RouterAccessSpec{
				Roles: []RouterAccessRole{
					{Name: "inter-router", Port: 55671},
					{Name: "edge", Port: 0},
				},
			},
			Status: RouterAccessStatus{
				Roles: []RouterAccessRole{
					{Name: "inter-router", Port: 55671},
					{Name: "edge", Port: 0},
				},
			},
		}
		unused := r.GetUnusedPorts()
		assert.Equal(t, len(unused), 0)
	})
	t.Run("status role dynamic port is kept", func(t *testing.T) {
		r := &RouterAccess{
			Spec: RouterAccessSpec{
				Roles: []RouterAccessRole{
					{Name: "edge", Port: 0},
				},
			},
			Status: RouterAccessStatus{
				Roles: []RouterAccessRole{
					{Name: "edge", Port: 45671},
				},
			},
		}
		unused := r.GetUnusedPorts()
		assert.Equal(t, len(unused), 0)
	})
}

func TestRouterAccess_ReleaseUnusedPorts(t *testing.T) {
	t.Run("removes a single unused port", func(t *testing.T) {
		r := &RouterAccess{}
		r.AllocatePort("inter-router", 55671)
		r.AllocatePort("edge", 45671)

		r.ReleaseUnusedPorts([]int32{55671})

		assert.Equal(t, r.GetAllocatedPortForRole("inter-router"), int32(0))
		assert.Equal(t, r.GetAllocatedPortForRole("edge"), int32(45671))
	})
	t.Run("removes multiple unused ports at once", func(t *testing.T) {
		r := &RouterAccess{}
		r.AllocatePort("inter-router", 55671)
		r.AllocatePort("edge", 45671)

		r.ReleaseUnusedPorts([]int32{55671, 45671})

		assert.Equal(t, len(r.Status.Roles), 0)
	})
}

func TestRouterAccess_GetPortForRole(t *testing.T) {
	t.Run("returns explicit spec port", func(t *testing.T) {
		r := &RouterAccess{
			Spec: RouterAccessSpec{
				Roles: []RouterAccessRole{
					{Name: "inter-router", Port: 12345},
				},
			},
		}
		assert.Equal(t, r.GetPortForRole("inter-router"), int32(12345))
	})
	t.Run("falls through to allocated port when spec port is 0", func(t *testing.T) {
		r := &RouterAccess{
			Spec: RouterAccessSpec{
				Roles: []RouterAccessRole{
					{Name: "edge"},
				},
			},
		}
		r.AllocatePort("edge", 45671)
		assert.Equal(t, r.GetPortForRole("edge"), int32(45671))
	})
}

func TestRouterAccess_SetConfigured(t *testing.T) {
	t.Run("nil error sets Configured=True", func(t *testing.T) {
		r := &RouterAccess{}
		assert.Assert(t, r.SetConfigured(nil))
		cond := meta.FindStatusCondition(r.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionTrue)
	})
	t.Run("repeated call with same state is a no-op", func(t *testing.T) {
		r := &RouterAccess{}
		r.SetConfigured(nil)
		assert.Assert(t, !r.SetConfigured(nil))
		r.SetConfigured(errors.New("some error"))
		assert.Assert(t, !r.SetConfigured(errors.New("some error")))
	})
	t.Run("non-nil error sets Configured=False", func(t *testing.T) {
		r := &RouterAccess{}
		assert.Assert(t, r.SetConfigured(errors.New("some error")))
		cond := meta.FindStatusCondition(r.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
	})
}

func TestRouterAccess_Resolve(t *testing.T) {
	t.Run("endpoints present sets Resolved=True", func(t *testing.T) {
		r := &RouterAccess{}
		r.SetConfigured(nil)
		assert.Assert(t, r.Resolve([]Endpoint{{Name: "inter-router", Host: "host1", Port: "443"}}, "skupper-router"))
		cond := meta.FindStatusCondition(r.Status.Conditions, CONDITION_TYPE_RESOLVED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionTrue)
	})
	t.Run("no endpoints sets Resolved=Pending", func(t *testing.T) {
		r := &RouterAccess{}
		r.SetConfigured(nil)
		r.Resolve([]Endpoint{{Name: "inter-router", Host: "host1", Port: "443"}}, "skupper-router")
		assert.Assert(t, r.Resolve([]Endpoint{}, "skupper-router"))
		cond := meta.FindStatusCondition(r.Status.Conditions, CONDITION_TYPE_RESOLVED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
	})
}

func TestRouterAccess_IsConfigured(t *testing.T) {
	r := &RouterAccess{}
	assert.Assert(t, !r.IsConfigured())
	r.SetConfigured(nil)
	assert.Assert(t, r.IsConfigured())
}

func TestNetwork_SetConfigured(t *testing.T) {
	t.Run("ready=true sets Configured=True", func(t *testing.T) {
		n := &Network{}
		assert.Assert(t, n.SetConfigured(true))
		cond := meta.FindStatusCondition(n.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionTrue)
	})
	t.Run("ready=false sets Configured=Pending", func(t *testing.T) {
		n := &Network{}
		assert.Assert(t, n.SetConfigured(false))
		cond := meta.FindStatusCondition(n.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
	})
}

func TestNetwork_SetError(t *testing.T) {
	expectedName := "network"
	t.Run("sets Configured=False with expected name in message", func(t *testing.T) {
		n := &Network{}
		assert.Assert(t, n.SetError(expectedName))
		cond := meta.FindStatusCondition(n.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
		assert.Equal(t, cond.Message, "name must be 'network'")
	})
	t.Run("repeated call with same state is a no-op", func(t *testing.T) {
		n := &Network{}
		n.SetError(expectedName)
		assert.Assert(t, !n.SetError(expectedName))
	})
}

func TestNetworkLink_SetConfigured(t *testing.T) {
	t.Run("ready=true sets Configured=True", func(t *testing.T) {
		nl := &NetworkLink{}
		assert.Assert(t, nl.SetConfigured(true))
		cond := meta.FindStatusCondition(nl.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionTrue)
	})
	t.Run("ready=false sets Configured=Pending", func(t *testing.T) {
		nl := &NetworkLink{}
		assert.Assert(t, nl.SetConfigured(false))
		cond := meta.FindStatusCondition(nl.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
	})
}

func TestNetworkLink_SetError(t *testing.T) {
	t.Run("sets Configured=False", func(t *testing.T) {
		nl := &NetworkLink{}
		assert.Assert(t, nl.SetError(errors.New("link error")))
		cond := meta.FindStatusCondition(nl.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
	})
	t.Run("repeated call with same state is a no-op", func(t *testing.T) {
		nl := &NetworkLink{}
		nl.SetError(errors.New("link error"))
		assert.Assert(t, !nl.SetError(errors.New("link error")))
	})
}

func TestNetworkLink_IsReady(t *testing.T) {
	t.Run("false before any condition is set", func(t *testing.T) {
		nl := &NetworkLink{}
		assert.Assert(t, !nl.IsReady())
	})
	t.Run("true after SetConfigured(true)", func(t *testing.T) {
		nl := &NetworkLink{}
		nl.SetConfigured(true)
		assert.Assert(t, nl.IsReady())
	})
}

func TestInterNetworkIngress_SetConfigured(t *testing.T) {
	t.Run("ready=true sets Configured=True", func(t *testing.T) {
		i := &InterNetworkIngress{}
		assert.Assert(t, i.SetConfigured(true))
		cond := meta.FindStatusCondition(i.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionTrue)
	})
	t.Run("ready=false sets Configured=Pending", func(t *testing.T) {
		i := &InterNetworkIngress{}
		assert.Assert(t, i.SetConfigured(false))
		cond := meta.FindStatusCondition(i.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
	})
}

func TestInterNetworkIngress_SetError(t *testing.T) {
	t.Run("sets Configured=False", func(t *testing.T) {
		i := &InterNetworkIngress{}
		assert.Assert(t, i.SetError(errors.New("ingress error")))
		cond := meta.FindStatusCondition(i.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
	})
	t.Run("repeated call with same state is a no-op", func(t *testing.T) {
		i := &InterNetworkIngress{}
		i.SetError(errors.New("ingress error"))
		assert.Assert(t, !i.SetError(errors.New("ingress error")))
	})
}

func TestNetworkAccess_SetConfigured(t *testing.T) {
	t.Run("first call stores networkId and sets Configured=True", func(t *testing.T) {
		r := &NetworkAccess{}
		assert.Assert(t, r.SetConfigured("net-1", nil))
		assert.Equal(t, r.Status.NetworkId, "net-1")
		cond := meta.FindStatusCondition(r.Status.Conditions, CONDITION_TYPE_CONFIGURED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionTrue)
	})
	t.Run("identical second call is a no-op", func(t *testing.T) {
		r := &NetworkAccess{}
		r.SetConfigured("net-1", nil)
		assert.Assert(t, !r.SetConfigured("net-1", nil))
	})
	t.Run("changed networkId returns true even when condition already set", func(t *testing.T) {
		r := &NetworkAccess{}
		r.SetConfigured("net-1", nil)
		assert.Assert(t, r.SetConfigured("net-2", nil))
		assert.Equal(t, r.Status.NetworkId, "net-2")
	})
}

func TestNetworkAccess_Resolve(t *testing.T) {
	t.Run("endpoints present sets Resolved=True", func(t *testing.T) {
		r := &NetworkAccess{}
		r.SetConfigured("net-1", nil)
		assert.Assert(t, r.Resolve([]Endpoint{{Name: "inter-router", Host: "host1", Port: "8080"}}, "skupper-router"))
		cond := meta.FindStatusCondition(r.Status.Conditions, CONDITION_TYPE_RESOLVED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionTrue)
	})
	t.Run("no endpoints sets Resolved=Pending", func(t *testing.T) {
		r := &NetworkAccess{}
		r.SetConfigured("net-1", nil)
		r.Resolve([]Endpoint{{Name: "inter-router", Host: "host1", Port: "8080"}}, "skupper-router")
		assert.Assert(t, r.Resolve([]Endpoint{}, "skupper-router"))
		cond := meta.FindStatusCondition(r.Status.Conditions, CONDITION_TYPE_RESOLVED)
		assert.Assert(t, cond != nil)
		assert.Equal(t, cond.Status, metav1.ConditionFalse)
	})
}

func TestNetworkAccess_Resolve_endpointMerge(t *testing.T) {
	setup := func() *NetworkAccess {
		r := &NetworkAccess{}
		r.SetConfigured("net-1", nil)
		r.Resolve([]Endpoint{
			{Name: "inter-network", Host: "host-original", Port: "443", Group: "skupper-router"},
		}, "skupper-router")
		return r
	}

	t.Run("endpoint retained unchanged reports no change", func(t *testing.T) {
		r := setup()
		assert.Assert(t, !r.Resolve([]Endpoint{
			{Name: "inter-network", Host: "host-original", Port: "443", Group: "skupper-router"},
		}, "skupper-router"))
	})
	t.Run("endpoint host updated reports change", func(t *testing.T) {
		r := setup()
		assert.Assert(t, r.Resolve([]Endpoint{
			{Name: "inter-network", Host: "host-new", Port: "443", Group: "skupper-router"},
		}, "skupper-router"))
		found := false
		for _, ep := range r.Status.Endpoints {
			if ep.Name == "inter-network" {
				found = true
				assert.Equal(t, ep.Host, "host-new")
			}
		}
		assert.Assert(t, found)
	})
	t.Run("endpoint removed from desired list reports change", func(t *testing.T) {
		r := setup()
		assert.Assert(t, r.Resolve([]Endpoint{
			{Name: "inter-network", Host: "host-replace", Port: "443", Group: "skupper-router"},
		}, "skupper-router"))
		for _, ep := range r.Status.Endpoints {
			assert.Assert(t, ep.Host != "host-original")
		}
	})
	t.Run("new endpoint added reports change", func(t *testing.T) {
		r := setup()
		assert.Assert(t, r.Resolve([]Endpoint{
			{Name: "inter-network", Host: "host-new", Port: "443", Group: "skupper-router-2"},
		}, "skupper-router-2"))
		found := false
		for _, ep := range r.Status.Endpoints {
			if ep.Group == "skupper-router-2" {
				found = true
			}
		}
		assert.Assert(t, found)
		assert.Equal(t, 2, len(r.Status.Endpoints))
	})
}

func TestNetworkAccess_IsConfigured(t *testing.T) {
	r := &NetworkAccess{}
	r.SetConfigured("net-1", nil)
	assert.Assert(t, r.IsConfigured("net-1"))
	assert.Assert(t, !r.IsConfigured("net-2"))
}

func TestNetworkAccess_AllocatePort(t *testing.T) {
	t.Run("allocates a new port", func(t *testing.T) {
		r := &NetworkAccess{}
		assert.Assert(t, r.AllocatePort(8080))
		assert.Equal(t, r.Status.Port, 8080)
	})
	t.Run("no change for same port", func(t *testing.T) {
		r := &NetworkAccess{}
		r.AllocatePort(8080)
		assert.Assert(t, !r.AllocatePort(8080))
	})
	t.Run("zero port returns false", func(t *testing.T) {
		r := &NetworkAccess{}
		assert.Assert(t, !r.AllocatePort(0))
	})
}

func TestNetworkAccess_GetPort(t *testing.T) {
	t.Run("spec port takes precedence", func(t *testing.T) {
		r := &NetworkAccess{
			Spec:   NetworkAccessSpec{Port: 9090},
			Status: NetworkAccessStatus{Port: 8080},
		}
		assert.Equal(t, r.GetPort(), 9090)
	})
	t.Run("falls back to status port when spec port is 0", func(t *testing.T) {
		r := &NetworkAccess{
			Status: NetworkAccessStatus{Port: 8080},
		}
		assert.Equal(t, r.GetPort(), 8080)
	})
}

func TestNetworkAccess_HasAllocatedPort(t *testing.T) {
	t.Run("false when both spec and status port are 0", func(t *testing.T) {
		r := &NetworkAccess{}
		assert.Assert(t, !r.HasAllocatedPort())
	})
	t.Run("false when spec port is set (not dynamically allocated)", func(t *testing.T) {
		r := &NetworkAccess{Spec: NetworkAccessSpec{Port: 9090}}
		assert.Assert(t, !r.HasAllocatedPort())
	})
	t.Run("true when spec port is 0 and status port is set", func(t *testing.T) {
		r := &NetworkAccess{}
		r.AllocatePort(8080)
		assert.Assert(t, r.HasAllocatedPort())
	})
}
