package runtimeapi

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	GroupName        = "runtime.skupper.io"
	GroupVersion     = "v1alpha1"
	APIVersion       = GroupName + "/" + GroupVersion
	EndpointKind     = "Endpoint"
	EndpointListKind = "EndpointList"
	EndpointResource = "endpoints"
	EndpointSingular = "endpoint"

	RouterNameAnnotation = "runtime.skupper.io/router-name"
	AdaptorEndpointsPath = "/runtime/v1alpha1/endpoints"
	DefaultAdaptorPort   = 9191
)

type EndpointType string

const (
	EndpointTypeTCPListener  EndpointType = "tcpListener"
	EndpointTypeTCPConnector EndpointType = "tcpConnector"
)

type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Status            EndpointStatus `json:"status,omitempty"`
}

type EndpointStatus struct {
	Type       EndpointType `json:"type"`
	Address    string       `json:"address,omitempty"`
	Host       string       `json:"host,omitempty"`
	Port       string       `json:"port,omitempty"`
	SiteID     string       `json:"siteId,omitempty"`
	SiteName   string       `json:"siteName,omitempty"`
	SslProfile string       `json:"sslProfile,omitempty"`
	Observer   string       `json:"observer,omitempty"`
	ProcessID  string       `json:"processId,omitempty"`
}

type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

func NewEndpointList(items []Endpoint) *EndpointList {
	if items == nil {
		items = []Endpoint{}
	}
	return &EndpointList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: APIVersion,
			Kind:       EndpointListKind,
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
				Name:         EndpointResource,
				SingularName: EndpointSingular,
				Namespaced:   true,
				Group:        GroupName,
				Version:      GroupVersion,
				Kind:         EndpointKind,
				Verbs:        metav1.Verbs{"get", "list"},
				ShortNames:   []string{"rtep"},
			},
		},
	}
}
