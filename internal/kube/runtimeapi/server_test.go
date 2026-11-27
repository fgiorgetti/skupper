package runtimeapi_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gotest.tools/v3/assert"
	authenticationv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	k8stesting "k8s.io/client-go/testing"

	"github.com/skupperproject/skupper/internal/kube/adaptor"
	"github.com/skupperproject/skupper/internal/kube/runtimeapi"
	"github.com/skupperproject/skupper/internal/qdr"
)

func TestDiscoveryAndHealth(t *testing.T) {
	srv := runtimeapi.NewServer(nil, "skupper", nil, nil)
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
		assert.Equal(t, resList.APIResources[0].Name, runtimeapi.AdaptorResource)
		assert.Equal(t, resList.APIResources[0].Kind, runtimeapi.AdaptorKind)
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

func TestEndToEndPushIngestAndAggregatedReadWithTLS(t *testing.T) {
	controlledNamespaces := func() []string { return []string{"west", "east"} }

	fakeKube := fake.NewSimpleClientset()
	fakeKube.PrependReactor("create", "tokenreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
		tr := action.(k8stesting.CreateAction).GetObject().(*authenticationv1.TokenReview)
		switch tr.Spec.Token {
		case "west-token":
			tr.Status = authenticationv1.TokenReviewStatus{
				Authenticated: true,
				User: authenticationv1.UserInfo{
					Username: "system:serviceaccount:west:skupper-router",
					Extra: map[string]authenticationv1.ExtraValue{
						"authentication.kubernetes.io/pod-name": {"skupper-router-west-0"},
					},
				},
			}
		case "east-token":
			tr.Status = authenticationv1.TokenReviewStatus{
				Authenticated: true,
				User: authenticationv1.UserInfo{
					Username: "system:serviceaccount:east:skupper-router",
					Extra: map[string]authenticationv1.ExtraValue{
						"authentication.kubernetes.io/pod-name": {"skupper-router-east-0"},
					},
				},
			}
		default:
			tr.Status = authenticationv1.TokenReviewStatus{
				Authenticated: false,
				Error:         "invalid token",
			}
		}
		return true, tr, nil
	})

	store := runtimeapi.NewRuntimeStateStore(time.Minute, controlledNamespaces)
	auth := runtimeapi.NewTokenReviewAuthenticator(fakeKube, runtimeapi.RuntimeTokenAudience, controlledNamespaces)
	srv := runtimeapi.NewServer(&runtimeapi.Config{
		Address:  "127.0.0.1:0",
		StateTTL: time.Minute,
	}, "skupper", store, auth)

	stopCh := make(chan struct{})
	defer close(stopCh)
	assert.NilError(t, srv.Start(stopCh))

	baseURL := "https://" + srv.Addr()
	tmpDir := t.TempDir()
	westTokenFile := filepath.Join(tmpDir, "west-token")
	eastTokenFile := filepath.Join(tmpDir, "east-token")
	assert.NilError(t, os.WriteFile(westTokenFile, []byte("west-token\n"), 0600))
	assert.NilError(t, os.WriteFile(eastTokenFile, []byte("east-token\n"), 0600))

	westPusher, err := adaptor.NewRuntimeStatePusher(adaptor.RuntimePusherConfig{
		TargetURL: baseURL + runtimeapi.IngestPathForNamespace("west"),
		TokenFile: westTokenFile,
	}, "west", "skupper-router-west-0", "site-west", "uid-west")
	assert.NilError(t, err)
	westPusher.UpdateState(qdr.BridgeRuntimeState{
		TcpListeners: map[string]qdr.TcpEndpointRuntime{
			"frontend": {
				TcpEndpoint:       qdr.TcpEndpoint{Name: "frontend", Address: "frontend:8080", Host: "0.0.0.0", Port: "8080"},
				OperStatus:        "up",
				ConnectionsOpened: 7,
				ConnectionsClosed: 6,
				BytesIn:           512,
				BytesOut:          1024,
			},
		},
	})
	assert.NilError(t, westPusher.PushOnce(context.Background()))

	eastPusher, err := adaptor.NewRuntimeStatePusher(adaptor.RuntimePusherConfig{
		TargetURL: baseURL + runtimeapi.IngestPathForNamespace("east"),
		TokenFile: eastTokenFile,
	}, "east", "skupper-router-east-0", "site-east", "uid-east")
	assert.NilError(t, err)
	eastPusher.UpdateState(qdr.BridgeRuntimeState{
		TcpConnectors: map[string]qdr.TcpEndpointRuntime{
			"backend": {
				TcpEndpoint: qdr.TcpEndpoint{Name: "backend", Address: "frontend:8080", Host: "10.1.2.3", Port: "8080"},
				OperStatus:  "up",
			},
		},
	})
	assert.NilError(t, eastPusher.PushOnce(context.Background()))

	httpsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	t.Run("cluster-wide table list served strictly from controller memory", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, baseURL+"/apis/"+runtimeapi.APIVersion+"/adaptors", nil)
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
		assert.Equal(t, table.Rows[0].Cells[1], string(runtimeapi.AdaptorTypeTCPConnector))
		assert.Equal(t, table.Rows[0].Cells[2], "up")
		assert.Equal(t, table.Rows[0].Cells[6], "skupper-router-east-0")
		assert.Equal(t, table.Rows[1].Cells[0], "frontend")
		assert.Equal(t, table.Rows[1].Cells[1], string(runtimeapi.AdaptorTypeTCPListener))
		assert.Equal(t, table.Rows[1].Cells[2], "up")
		assert.Equal(t, table.Rows[1].Cells[6], "skupper-router-west-0")
		assert.Equal(t, table.Rows[1].Cells[8], float64(7))
		assert.Equal(t, table.Rows[1].Cells[10], float64(512))

		var rowObj runtimeapi.Adaptor
		assert.NilError(t, json.Unmarshal(table.Rows[0].Object.Raw, &rowObj))
		assert.Equal(t, rowObj.Namespace, "east")
		assert.Equal(t, rowObj.Name, "backend")
		assert.Equal(t, rowObj.Status.PodName, "skupper-router-east-0")
	})

	t.Run("namespaced JSON list and selectors", func(t *testing.T) {
		resp, err := httpsClient.Get(baseURL + "/apis/" + runtimeapi.APIVersion + "/namespaces/west/adaptors?labelSelector=runtime.skupper.io/type=tcpListener&fieldSelector=metadata.name=frontend")
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		var list runtimeapi.AdaptorList
		assert.NilError(t, json.NewDecoder(resp.Body).Decode(&list))
		assert.Equal(t, len(list.Items), 1)
		assert.Equal(t, list.Items[0].Name, "frontend")
		assert.Equal(t, list.Items[0].Namespace, "west")
	})

	t.Run("cross-namespace push rejected with 403", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, baseURL+runtimeapi.IngestPathForNamespace("west"), bytes.NewBufferString(`{}`))
		assert.NilError(t, err)
		req.Header.Set("Authorization", "Bearer east-token")
		resp, err := httpsClient.Do(req)
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusForbidden)
	})

	t.Run("unauthenticated push rejected with 401", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, baseURL+runtimeapi.IngestPathForNamespace("west"), bytes.NewBufferString(`{}`))
		assert.NilError(t, err)
		req.Header.Set("Authorization", "Bearer bad-token")
		resp, err := httpsClient.Do(req)
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusUnauthorized)
	})

	t.Run("namespaced get as table and includeObject=None", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, baseURL+"/apis/"+runtimeapi.APIVersion+"/namespaces/west/adaptors/frontend?includeObject=None", nil)
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

	t.Run("get missing adaptor returns 404 Status", func(t *testing.T) {
		resp, err := httpsClient.Get(baseURL + "/apis/" + runtimeapi.APIVersion + "/namespaces/west/adaptors/does-not-exist")
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusNotFound)

		var status metav1.Status
		assert.NilError(t, json.NewDecoder(resp.Body).Decode(&status))
		assert.Equal(t, status.Reason, metav1.StatusReasonNotFound)
	})

	t.Run("watch and non-GET return 405 Status", func(t *testing.T) {
		resp, err := httpsClient.Get(baseURL + "/apis/" + runtimeapi.APIVersion + "/namespaces/west/adaptors?watch=true")
		assert.NilError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, resp.StatusCode, http.StatusMethodNotAllowed)

		postResp, err := httpsClient.Post(fmt.Sprintf("%s/apis/%s/namespaces/west/adaptors", baseURL, runtimeapi.APIVersion), "application/json", nil)
		assert.NilError(t, err)
		defer postResp.Body.Close()
		assert.Equal(t, postResp.StatusCode, http.StatusMethodNotAllowed)
	})
}
