package runtimeapi

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/internal/qdr"
)

const (
	GroupName       = "runtime.skupper.io"
	GroupVersion    = "v1alpha1"
	APIVersion      = GroupName + "/" + GroupVersion
	AdaptorKind     = "Adaptor"
	AdaptorListKind = "AdaptorList"
	AdaptorResource = "adaptors"
	AdaptorSingular = "adaptor"

	RouterNameAnnotation = "runtime.skupper.io/router-name"
	AdaptorTypeLabel     = "runtime.skupper.io/type"

	RuntimeTokenAudience    = "skupper-runtime-ingest"
	DefaultRuntimeTokenPath = "/var/run/secrets/skupper/runtime-token/token"
	DefaultRuntimeStateTTL  = 90 * time.Second
	DefaultPushInterval     = 30 * time.Second

	IngestPathPrefix = "/internal/v1alpha1/namespaces/"
	IngestPathSuffix = "/runtime-state"
)

func IngestPathForNamespace(namespace string) string {
	return fmt.Sprintf("%s%s%s", IngestPathPrefix, namespace, IngestPathSuffix)
}

type AdaptorType string

const (
	AdaptorTypeTCPListener  AdaptorType = "tcpListener"
	AdaptorTypeTCPConnector AdaptorType = "tcpConnector"
)

type Adaptor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Status            AdaptorStatus `json:"status,omitempty"`
}

type AdaptorStatus struct {
	Type              AdaptorType `json:"type"`
	OperStatus        string      `json:"operStatus,omitempty"`
	Address           string      `json:"address,omitempty"`
	Host              string      `json:"host,omitempty"`
	Port              string      `json:"port,omitempty"`
	PodName           string      `json:"podName,omitempty"`
	ConnectionsOpened uint64      `json:"connectionsOpened,omitempty"`
	ConnectionsClosed uint64      `json:"connectionsClosed,omitempty"`
	BytesIn           uint64      `json:"bytesIn,omitempty"`
	BytesOut          uint64      `json:"bytesOut,omitempty"`
	SiteID            string      `json:"siteId,omitempty"`
	SiteName          string      `json:"siteName,omitempty"`
	SslProfile        string      `json:"sslProfile,omitempty"`
	Observer          string      `json:"observer,omitempty"`
	ProcessID         string      `json:"processId,omitempty"`
}

type AdaptorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Adaptor `json:"items"`
}

// RuntimeStatePushRequest is pushed by kube-adaptor pods to the controller ingest endpoint.
type RuntimeStatePushRequest struct {
	PodName       string                            `json:"podName,omitempty"`
	SiteName      string                            `json:"siteName,omitempty"`
	SiteID        string                            `json:"siteId,omitempty"`
	TCPListeners  map[string]qdr.TcpEndpointRuntime `json:"tcpListeners,omitempty"`
	TCPConnectors map[string]qdr.TcpEndpointRuntime `json:"tcpConnectors,omitempty"`
	Adaptors      []Adaptor                         `json:"adaptors,omitempty"`
}

func NewAdaptorList(items []Adaptor) *AdaptorList {
	if items == nil {
		items = []Adaptor{}
	}
	return &AdaptorList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: APIVersion,
			Kind:       AdaptorListKind,
		},
		Items: items,
	}
}

func apiGroup() metav1.APIGroup {
	gv := metav1.GroupVersionForDiscovery{
		GroupVersion: APIVersion,
		Version:      GroupVersion,
	}
	return metav1.APIGroup{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "APIGroup",
		},
		Name:             GroupName,
		Versions:         []metav1.GroupVersionForDiscovery{gv},
		PreferredVersion: gv,
	}
}

func apiGroupList() metav1.APIGroupList {
	return metav1.APIGroupList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "APIGroupList",
		},
		Groups: []metav1.APIGroup{apiGroup()},
	}
}

func apiResourceList() metav1.APIResourceList {
	return metav1.APIResourceList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "APIResourceList",
		},
		GroupVersion: APIVersion,
		APIResources: []metav1.APIResource{
			{
				Name:         AdaptorResource,
				SingularName: AdaptorSingular,
				Namespaced:   true,
				Group:        GroupName,
				Version:      GroupVersion,
				Kind:         AdaptorKind,
				Verbs:        metav1.Verbs{"get", "list"},
				ShortNames:   []string{"adp"},
			},
		},
	}
}
