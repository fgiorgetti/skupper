package adaptor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
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

	pusher.UpdateBridges(qdr.BridgeConfig{
		TcpListeners: qdr.TcpEndpointMap{
			"web": {Name: "web", Address: "web:80", Host: "0.0.0.0", Port: "80"},
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
}

func TestRuntimeStatePusherHeartbeatLoop(t *testing.T) {
	received := make(chan struct{}, 10)
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received <- struct{}{}
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

	stopCh := make(chan struct{})
	defer close(stopCh)
	pusher.Start(stopCh)

	pusher.UpdateBridges(qdr.BridgeConfig{
		TcpConnectors: qdr.TcpEndpointMap{
			"db": {Name: "db", Address: "db:5432", Host: "10.0.0.2", Port: "5432"},
		},
	})

	// Expect immediate push + at least one periodic heartbeat push.
	for i := 0; i < 2; i++ {
		select {
		case <-received:
		case <-time.After(time.Second):
			t.Fatalf("timed out waiting for push %d", i+1)
		}
	}
}
