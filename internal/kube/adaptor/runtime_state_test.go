package adaptor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"github.com/skupperproject/skupper/internal/kube/runtimeapi"
	"github.com/skupperproject/skupper/internal/qdr"
)

func TestRuntimeStateStoreSetAndQuery(t *testing.T) {
	store := NewRuntimeStateStore("west", "site-west", "site-uid-1")
	assert.Equal(t, len(store.ListEndpoints()), 0)

	store.SetBridges(qdr.BridgeConfig{
		TcpListeners: qdr.TcpEndpointMap{
			"backend:8080": {
				Name:       "backend:8080",
				Address:    "backend:8080",
				Host:       "0.0.0.0",
				Port:       "8080",
				SslProfile: "tls-profile",
			},
		},
		TcpConnectors: qdr.TcpEndpointMap{
			"backend@10.0.0.2": {
				Name:     "backend@10.0.0.2",
				Address:  "backend:8080",
				Host:     "10.0.0.2",
				Port:     "8080",
				Observer: "http1",
			},
		},
	})

	items := store.ListEndpoints()
	assert.Equal(t, len(items), 2)

	assert.Equal(t, items[0].Name, "backend-10.0.0.2")
	assert.Equal(t, items[0].Namespace, "west")
	assert.Equal(t, items[0].Status.Type, runtimeapi.EndpointTypeTCPConnector)
	assert.Equal(t, items[0].Status.SiteID, "site-uid-1")
	assert.Equal(t, items[0].Status.SiteName, "site-west")
	assert.Equal(t, items[0].Status.Observer, "http1")
	assert.Equal(t, items[0].Annotations[runtimeapi.RouterNameAnnotation], "backend@10.0.0.2")
	assert.Equal(t, items[0].Labels[endpointTypeLabel], string(runtimeapi.EndpointTypeTCPConnector))

	assert.Equal(t, items[1].Name, "backend-8080")
	assert.Equal(t, items[1].Status.Type, runtimeapi.EndpointTypeTCPListener)
	assert.Equal(t, items[1].Status.SslProfile, "tls-profile")
	assert.Equal(t, items[1].Annotations[runtimeapi.RouterNameAnnotation], "backend:8080")

	ep, ok := store.GetEndpoint("backend-8080")
	assert.Assert(t, ok)
	assert.Equal(t, ep.Status.Port, "8080")

	_, ok = store.GetEndpoint("missing")
	assert.Assert(t, !ok)

	firstSeen := items[1].CreationTimestamp
	time.Sleep(10 * time.Millisecond)
	store.SetTCPListeners(map[string]qdr.TcpEndpoint{
		"backend:8080": {
			Name:    "backend:8080",
			Address: "backend:8080",
			Host:    "0.0.0.0",
			Port:    "9090",
		},
	})
	updated, ok := store.GetEndpoint("backend-8080")
	assert.Assert(t, ok)
	assert.Equal(t, updated.Status.Port, "9090")
	assert.Equal(t, updated.CreationTimestamp, firstSeen)
}

func TestRuntimeStateStoreNameCollisionDisambiguation(t *testing.T) {
	store := NewRuntimeStateStore("east", "site-east", "site-uid-2")
	store.SetTCPListeners(map[string]qdr.TcpEndpoint{
		"echo": {Name: "echo", Address: "echo:9000", Host: "0.0.0.0", Port: "9000"},
	})
	store.SetTCPConnectors(map[string]qdr.TcpEndpoint{
		"echo": {Name: "echo", Address: "echo:9000", Host: "10.0.0.5", Port: "9000"},
	})

	items := store.ListEndpoints()
	assert.Equal(t, len(items), 2)
	assert.Equal(t, items[0].Name, "echo")
	assert.Equal(t, items[0].Status.Type, runtimeapi.EndpointTypeTCPListener)
	assert.Equal(t, items[1].Name, "echo-connector")
	assert.Equal(t, items[1].Status.Type, runtimeapi.EndpointTypeTCPConnector)
}

func TestRuntimeStateHandler(t *testing.T) {
	store := NewRuntimeStateStore("west", "site-west", "uid-1")
	store.SetTCPListeners(map[string]qdr.TcpEndpoint{
		"web": {Name: "web", Address: "web:80", Host: "0.0.0.0", Port: "80"},
	})
	handler := RuntimeStateHandler(store)

	t.Run("list", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, runtimeapi.AdaptorEndpointsPath, nil)
		handler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)

		var list runtimeapi.EndpointList
		assert.NilError(t, json.Unmarshal(rec.Body.Bytes(), &list))
		assert.Equal(t, list.APIVersion, runtimeapi.APIVersion)
		assert.Equal(t, list.Kind, runtimeapi.EndpointListKind)
		assert.Equal(t, len(list.Items), 1)
		assert.Equal(t, list.Items[0].Name, "web")
	})

	t.Run("get", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, runtimeapi.AdaptorEndpointsPath+"/web", nil)
		handler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusOK)

		var ep runtimeapi.Endpoint
		assert.NilError(t, json.Unmarshal(rec.Body.Bytes(), &ep))
		assert.Equal(t, ep.Name, "web")
		assert.Equal(t, ep.Status.Address, "web:80")
	})

	t.Run("not found", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, runtimeapi.AdaptorEndpointsPath+"/nope", nil)
		handler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusNotFound)
	})

	t.Run("method not allowed", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, runtimeapi.AdaptorEndpointsPath, nil)
		handler.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, http.StatusMethodNotAllowed)
	})
}
