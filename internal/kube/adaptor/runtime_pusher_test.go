package adaptor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/skupperproject/skupper/internal/kube/runtimeapi"
	"github.com/skupperproject/skupper/internal/qdr"
)

func TestRuntimeStatePusherTokenRotationAndPayload(t *testing.T) {
	var (
		mu       sync.Mutex
		tokens   []string
		payloads []runtimeapi.RuntimeStatePushRequest
	)

	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)
		var req runtimeapi.RuntimeStatePushRequest
		assert.NilError(t, json.NewDecoder(r.Body).Decode(&req))

		mu.Lock()
		tokens = append(tokens, r.Header.Get("Authorization"))
		payloads = append(payloads, req)
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	tokenFile := filepath.Join(t.TempDir(), "token")
	assert.NilError(t, os.WriteFile(tokenFile, []byte("token-v1\n"), 0600))

	pusher, err := NewRuntimeStatePusher(RuntimePusherConfig{
		TargetURL: srv.URL + runtimeapi.IngestPathForNamespace("west"),
		TokenFile: tokenFile,
	}, "west", "skupper-router-0", "site-west", "uid-west")
	assert.NilError(t, err)

	pusher.UpdateState(qdr.BridgeRuntimeState{
		TcpListeners: map[string]qdr.TcpEndpointRuntime{
			"web": {
				TcpEndpoint:       qdr.TcpEndpoint{Name: "web", Address: "web:80", Host: "0.0.0.0", Port: "80"},
				OperStatus:        "up",
				ConnectionsOpened: 3,
				BytesIn:           256,
			},
		},
	})
	assert.NilError(t, pusher.PushOnce(context.Background()))

	// Simulate kubelet projected token rotation on disk.
	assert.NilError(t, os.WriteFile(tokenFile, []byte("token-v2\n"), 0600))
	assert.NilError(t, pusher.PushOnce(context.Background()))

	mu.Lock()
	defer mu.Unlock()
	assert.Equal(t, len(tokens), 2)
	assert.Equal(t, tokens[0], "Bearer token-v1")
	assert.Equal(t, tokens[1], "Bearer token-v2")
	assert.Equal(t, payloads[0].PodName, "skupper-router-0")
	assert.Equal(t, payloads[0].SiteName, "site-west")
	assert.Equal(t, payloads[0].SiteID, "uid-west")
	assert.Equal(t, payloads[0].TCPListeners["web"].Port, "80")
	assert.Equal(t, payloads[0].TCPListeners["web"].OperStatus, "up")
	assert.Equal(t, payloads[0].TCPListeners["web"].ConnectionsOpened, uint64(3))
}

func TestRuntimeStatePusherHeartbeatRefreshesLiveRouterState(t *testing.T) {
	received := make(chan runtimeapi.RuntimeStatePushRequest, 10)
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req runtimeapi.RuntimeStatePushRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err == nil {
			received <- req
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	tokenFile := filepath.Join(t.TempDir(), "token")
	assert.NilError(t, os.WriteFile(tokenFile, []byte("heartbeat-token"), 0600))

	pusher, err := NewRuntimeStatePusher(RuntimePusherConfig{
		TargetURL:    srv.URL,
		TokenFile:    tokenFile,
		PushInterval: 40 * time.Millisecond,
	}, "west", "pod-1", "site-west", "uid-1")
	assert.NilError(t, err)

	var connCount atomic.Uint64
	pusher.SetStateFetcher(func() (*qdr.BridgeRuntimeState, error) {
		n := connCount.Add(1)
		state := qdr.NewBridgeRuntimeState()
		state.TcpConnectors["db"] = qdr.TcpEndpointRuntime{
			TcpEndpoint:       qdr.TcpEndpoint{Name: "db", Address: "db:5432", Host: "10.0.0.2", Port: "5432"},
			OperStatus:        "up",
			ConnectionsOpened: n,
		}
		return &state, nil
	})

	stopCh := make(chan struct{})
	defer close(stopCh)
	pusher.Start(stopCh)

	// Wait for two heartbeat ticks and verify live counters advanced without any UpdateState call.
	var first, second runtimeapi.RuntimeStatePushRequest
	select {
	case first = <-received:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for first heartbeat push")
	}
	select {
	case second = <-received:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for second heartbeat push")
	}
	assert.Assert(t, second.TCPConnectors["db"].ConnectionsOpened > first.TCPConnectors["db"].ConnectionsOpened)
	assert.Equal(t, second.TCPConnectors["db"].OperStatus, "up")
}
