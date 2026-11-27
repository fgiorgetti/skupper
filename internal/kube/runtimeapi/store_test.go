package runtimeapi

import (
	"context"
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/skupperproject/skupper/internal/qdr"
)

func TestRuntimeStateStoreMultiPodAggregationAndFirstSeen(t *testing.T) {
	store := NewRuntimeStateStore(time.Minute, func() []string { return []string{"west", "east"} })

	store.Upsert("west", "skupper-router-pod-1", RuntimeStatePushRequest{
		SiteName: "site-west",
		SiteID:   "uid-west",
		TCPListeners: map[string]qdr.TcpEndpointRuntime{
			"backend:8080": {
				TcpEndpoint:       qdr.TcpEndpoint{Name: "backend:8080", Address: "backend:8080", Host: "0.0.0.0", Port: "8080"},
				OperStatus:        "up",
				ConnectionsOpened: 5,
				ConnectionsClosed: 4,
				BytesIn:           1024,
				BytesOut:          2048,
			},
		},
		TCPConnectors: map[string]qdr.TcpEndpointRuntime{
			"pod1@10.0.0.1": {
				TcpEndpoint: qdr.TcpEndpoint{Name: "pod1@10.0.0.1", Address: "svc:80", Host: "10.0.0.1", Port: "80"},
				OperStatus:  "up",
			},
		},
	})

	initialItems, err := store.List(context.Background(), "west")
	assert.NilError(t, err)
	assert.Equal(t, len(initialItems), 2)
	assert.Equal(t, initialItems[0].Name, "backend-8080")
	assert.Equal(t, initialItems[0].Status.PodName, "skupper-router-pod-1")
	assert.Equal(t, initialItems[0].Status.OperStatus, "up")
	assert.Equal(t, initialItems[0].Status.ConnectionsOpened, uint64(5))
	assert.Equal(t, initialItems[0].Status.BytesOut, uint64(2048))
	assert.Equal(t, initialItems[0].Annotations[RouterNameAnnotation], "backend:8080")
	assert.Equal(t, initialItems[1].Name, "pod1-10.0.0.1")
	firstSeenListener := initialItems[0].CreationTimestamp

	time.Sleep(15 * time.Millisecond)

	// Second pod reports the shared listener (with updated counters) plus its own distinct connector.
	store.Upsert("west", "skupper-router-pod-2", RuntimeStatePushRequest{
		SiteName: "site-west",
		SiteID:   "uid-west",
		TCPListeners: map[string]qdr.TcpEndpointRuntime{
			"backend:8080": {
				TcpEndpoint:       qdr.TcpEndpoint{Name: "backend:8080", Address: "backend:8080", Host: "0.0.0.0", Port: "9090"},
				OperStatus:        "up",
				ConnectionsOpened: 9,
			},
		},
		TCPConnectors: map[string]qdr.TcpEndpointRuntime{
			"pod2@10.0.0.2": {
				TcpEndpoint: qdr.TcpEndpoint{Name: "pod2@10.0.0.2", Address: "svc:80", Host: "10.0.0.2", Port: "80"},
			},
		},
	})

	items, err := store.List(context.Background(), "west")
	assert.NilError(t, err)
	assert.Equal(t, len(items), 3)
	assert.Equal(t, items[0].Name, "backend-8080")
	assert.Equal(t, items[0].Status.PodName, "skupper-router-pod-2")
	assert.Equal(t, items[0].Status.Port, "9090")
	assert.Equal(t, items[0].Status.ConnectionsOpened, uint64(9))
	assert.Equal(t, items[0].CreationTimestamp, firstSeenListener)
	assert.Equal(t, items[1].Name, "pod1-10.0.0.1")
	assert.Equal(t, items[1].Status.PodName, "skupper-router-pod-1")
	assert.Equal(t, items[2].Name, "pod2-10.0.0.2")
	assert.Equal(t, items[2].Status.PodName, "skupper-router-pod-2")
}

func TestRuntimeStateStoreTTLExpirationAndNamespaceCleanup(t *testing.T) {
	controlled := []string{"west", "east"}
	store := NewRuntimeStateStore(60*time.Millisecond, func() []string { return controlled })

	store.Upsert("west", "pod-a", RuntimeStatePushRequest{
		TCPListeners: map[string]qdr.TcpEndpointRuntime{
			"ep-a": {TcpEndpoint: qdr.TcpEndpoint{Name: "ep-a", Address: "a:80", Host: "0.0.0.0", Port: "80"}},
		},
	})
	store.Upsert("east", "pod-east", RuntimeStatePushRequest{
		TCPListeners: map[string]qdr.TcpEndpointRuntime{
			"ep-east": {TcpEndpoint: qdr.TcpEndpoint{Name: "ep-east", Address: "e:80", Host: "0.0.0.0", Port: "80"}},
		},
	})

	time.Sleep(40 * time.Millisecond)

	// Refresh only pod-b in west; pod-a is left to expire.
	store.Upsert("west", "pod-b", RuntimeStatePushRequest{
		TCPListeners: map[string]qdr.TcpEndpointRuntime{
			"ep-b": {TcpEndpoint: qdr.TcpEndpoint{Name: "ep-b", Address: "b:80", Host: "0.0.0.0", Port: "80"}},
		},
	})

	time.Sleep(35 * time.Millisecond)

	westItems, err := store.List(context.Background(), "west")
	assert.NilError(t, err)
	assert.Equal(t, len(westItems), 1)
	assert.Equal(t, westItems[0].Name, "ep-b")

	_, err = store.Get(context.Background(), "west", "ep-a")
	assert.ErrorIs(t, err, ErrNotFound)

	// Removing east from controlled namespaces hides and prunes its entries.
	controlled = []string{"west"}
	store.Prune()

	eastItems, err := store.List(context.Background(), "east")
	assert.NilError(t, err)
	assert.Equal(t, len(eastItems), 0)
}

func TestRuntimeStateStoreCollisionDisambiguation(t *testing.T) {
	store := NewRuntimeStateStore(time.Minute, nil)
	store.Upsert("west", "pod-1", RuntimeStatePushRequest{
		TCPListeners: map[string]qdr.TcpEndpointRuntime{
			"echo": {TcpEndpoint: qdr.TcpEndpoint{Name: "echo", Address: "echo:9000", Host: "0.0.0.0", Port: "9000"}},
		},
		TCPConnectors: map[string]qdr.TcpEndpointRuntime{
			"echo": {TcpEndpoint: qdr.TcpEndpoint{Name: "echo", Address: "echo:9000", Host: "10.0.0.5", Port: "9000"}},
		},
	})

	items, err := store.List(context.Background(), "west")
	assert.NilError(t, err)
	assert.Equal(t, len(items), 2)
	assert.Equal(t, items[0].Name, "echo")
	assert.Equal(t, items[0].Status.Type, AdaptorTypeTCPListener)
	assert.Equal(t, items[1].Name, "echo-connector")
	assert.Equal(t, items[1].Status.Type, AdaptorTypeTCPConnector)
}
