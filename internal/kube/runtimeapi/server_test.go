package runtimeapi_test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/internal/kube/adaptor"
	"github.com/skupperproject/skupper/internal/kube/runtimeapi"
	"github.com/skupperproject/skupper/internal/qdr"
)

func TestDiscoveryAndHealth(t *testing.T) {
	srv := runtimeapi.NewServer(nil, "skupper", nil)
	handler := srv.Handler()

	for _, path := range []string{"/healthz", "/livez", "/readyz"} {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
		assert.Equal(t, rec.Code, http.StatusOK)
		assert.Equal(t, rec.Body.String(), "ok")
	}

	t.Run("apis", func(t *testing.T) {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/apis", nil))
		assert.Equal(t, rec.Code, http.StatusOK)

		var groupList metav1.APIGroupList
		assert.NilError(t, json.Unmarshal(rec.Body.Bytes(), &groupList))
		assert.Equal(t, len(groupList.Groups), 1)
		assert.Equal(t, groupList.Groups[0].Name, runtimeapi.GroupName)
	})

	t.Run("group", func(t *testing.T) {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/apis/"+runtimeapi.GroupName, nil))
		assert.Equal(t, rec.Code, http.StatusOK)

		var group metav1.APIGroup
		assert.NilError(t, json.Unmarshal(rec.Body.Bytes(), &group))
		assert.Equal(t, group.PreferredVersion.GroupVersion, runtimeapi.APIVersion)
	})

	t.Run("group version resources", func(t *testing.T) {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/apis/"+runtimeapi.APIVersion, nil))
		assert.Equal(t, rec.Code, http.StatusOK)

		var resList metav1.APIResourceList
		assert.NilError(t, json.Unmarshal(rec.Body.Bytes(), &resList))
		assert.Equal(t, resList.GroupVersion, runtimeapi.APIVersion)
		assert.Equal(t, len(resList.APIResources), 1)
		assert.Equal(t, resList.APIResources[0].Name, runtimeapi.EndpointResource)
		assert.Equal(t, resList.APIResources[0].Kind, runtimeapi.EndpointKind)
		assert.Assert(t, resList.APIResources[0].Namespaced)
	})

	t.Run("openapi returns 404 status", func(t *testing.T) {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/openapi/v3", nil))
		assert.Equal(t, rec.Code, http.StatusNotFound)
		var status metav1.Status
		assert.NilError(t, json.Unmarshal(rec.Body.Bytes(), &status))
		assert.Equal(t, status.Reason, metav1.StatusReasonNotFound)
	})
}

func TestEndToEndWithAdaptorProviderAndTLS(t *testing.T) {
	westStore := adaptor.NewRuntimeStateStore("west", "site-west", "uid-west")
	westStore.SetBridges(qdr.BridgeConfig{
		TcpListeners: qdr.TcpEndpointMap{
			"frontend": {
				Name:    "frontend",
				Address: "frontend:8080",
				Host:    "0.0.0.0",
				Port:    "8080",
			},
		},
	})
	eastStore := adaptor.NewRuntimeStateStore("east", "site-east", "uid-east")
	eastStore.SetBridges(qdr.BridgeConfig{
		TcpConnectors: qdr.TcpEndpointMap{
			"backend": {
				Name:    "backend",
				Address: "frontend:8080",
				Host:    "10.1.2.3",
				Port:    "8080",
			},
		},
	})

	westAdaptor := httptest.NewServer(adaptor.RuntimeStateHandler(westStore))
	defer westAdaptor.Close()
	eastAdaptor := httptest.NewServer(adaptor.RuntimeStateHandler(eastStore))
	defer eastAdaptor.Close()

	provider := runtimeapi.NewAdaptorProvider(func() []string {
		return []string{"west", "east", "unreachable"}
	})
	provider.BaseURL = func(ns string) string {
		switch ns {
		case "west":
			return westAdaptor.URL
		case "east":
			return eastAdaptor.URL
		default:
			return "http://127.0.0.1:1"
		}
	}

	srv := runtimeapi.NewServer(&runtimeapi.Config{
		Address: "127.0.0.1:0",
	}, "skupper", provider)

	stopCh := make(chan struct{})
	defer close(stopCh)
	assert.NilError(t, srv.Start(stopCh))

	httpsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	baseURL := "https://" + srv.Addr()

	t.Run("cluster-wide table list aggregates healthy sites", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, baseURL+"/apis/"+runtimeapi.APIVersion+"/endpoints", nil)
		assert.NilError(t, err)
		req.Header.Set("Accept", "application/json;as=Table;v=v1;g=meta.k8s.io,application/json")

		resp, err := httpsClient.Do(req)
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		var table metav1.Table
		assert.NilError(t, json.NewDecoder(resp.Body).Decode(&table))
		assert.Equal(t, table.APIVersion, "meta.k8s.io/v1")
		assert.Equal(t, table.Kind, "Table")
		assert.Equal(t, len(table.Rows), 2)
		assert.Equal(t, table.Rows[0].Cells[0], "backend")
		assert.Equal(t, table.Rows[0].Cells[1], string(runtimeapi.EndpointTypeTCPConnector))
		assert.Equal(t, table.Rows[1].Cells[0], "frontend")
		assert.Equal(t, table.Rows[1].Cells[1], string(runtimeapi.EndpointTypeTCPListener))

		var rowObj runtimeapi.Endpoint
		assert.NilError(t, json.Unmarshal(table.Rows[0].Object.Raw, &rowObj))
		assert.Equal(t, rowObj.Namespace, "east")
		assert.Equal(t, rowObj.Name, "backend")
	})

	t.Run("namespaced JSON list and selectors", func(t *testing.T) {
		resp, err := httpsClient.Get(baseURL + "/apis/" + runtimeapi.APIVersion + "/namespaces/west/endpoints?labelSelector=runtime.skupper.io/type=tcpListener&fieldSelector=metadata.name=frontend")
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		var list runtimeapi.EndpointList
		assert.NilError(t, json.NewDecoder(resp.Body).Decode(&list))
		assert.Equal(t, len(list.Items), 1)
		assert.Equal(t, list.Items[0].Name, "frontend")
		assert.Equal(t, list.Items[0].Namespace, "west")
	})

	t.Run("namespace without site returns empty list", func(t *testing.T) {
		resp, err := httpsClient.Get(baseURL + "/apis/" + runtimeapi.APIVersion + "/namespaces/kube-system/endpoints")
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		var list runtimeapi.EndpointList
		assert.NilError(t, json.NewDecoder(resp.Body).Decode(&list))
		assert.Equal(t, len(list.Items), 0)
	})

	t.Run("namespaced get as table and includeObject=None", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, baseURL+"/apis/"+runtimeapi.APIVersion+"/namespaces/west/endpoints/frontend?includeObject=None", nil)
		assert.NilError(t, err)
		req.Header.Set("Accept", "application/json;as=Table;v=v1;g=meta.k8s.io")

		resp, err := httpsClient.Do(req)
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		var table metav1.Table
		assert.NilError(t, json.NewDecoder(resp.Body).Decode(&table))
		assert.Equal(t, len(table.Rows), 1)
		assert.Equal(t, len(table.Rows[0].Object.Raw), 0)
	})

	t.Run("get missing endpoint returns 404 Status", func(t *testing.T) {
		resp, err := httpsClient.Get(baseURL + "/apis/" + runtimeapi.APIVersion + "/namespaces/west/endpoints/does-not-exist")
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusNotFound)

		var status metav1.Status
		assert.NilError(t, json.NewDecoder(resp.Body).Decode(&status))
		assert.Equal(t, status.Reason, metav1.StatusReasonNotFound)
	})

	t.Run("watch and non-GET return 405 Status", func(t *testing.T) {
		resp, err := httpsClient.Get(baseURL + "/apis/" + runtimeapi.APIVersion + "/namespaces/west/endpoints?watch=true")
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusMethodNotAllowed)

		postResp, err := httpsClient.Post(fmt.Sprintf("%s/apis/%s/namespaces/west/endpoints", baseURL, runtimeapi.APIVersion), "application/json", nil)
		assert.NilError(t, err)
		defer postResp.Body.Close()
		assert.Equal(t, postResp.StatusCode, http.StatusMethodNotAllowed)
	})
}
