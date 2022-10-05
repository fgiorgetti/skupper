package client

import (
	"context"
	"fmt"
	"strconv"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/pkg/kube"
	kubeqdr "github.com/skupperproject/skupper/pkg/kube/qdr"
	"github.com/skupperproject/skupper/pkg/qdr"
)

func (cli *VanClient) getRouterConfig() (*qdr.RouterConfig, error) {
	configmap, err := kube.GetConfigMap(types.TransportConfigMapName, cli.Namespace, cli.KubeClient)
	if errors.IsNotFound(err) {
		return nil, fmt.Errorf("Skupper is not installed in %s", cli.Namespace)
	} else if err != nil {
		return nil, err
	}
	current, err := qdr.GetRouterConfigFromConfigMap(configmap)
	if err != nil {
		return nil, err
	}
	return current, nil
}

func (cli *VanClient) ConnectorList(ctx context.Context) ([]types.LinkStatus, error) {
	var links []types.LinkStatus
	secrets, err := cli.KubeClient.CoreV1().Secrets(cli.Namespace).List(metav1.ListOptions{LabelSelector: "skupper.io/type in (connection-token, token-claim)"})
	if err != nil {
		return links, err
	}
	current, err := cli.getRouterConfig()
	if err != nil {
		return links, err
	}
	edge := current.IsEdge()
	connections, _ := kubeqdr.GetConnections(cli.Namespace, cli.KubeClient, cli.RestConfig)
	for _, s := range secrets.Items {
		links = append(links, qdr.GetLinkStatus(&s, edge, connections))
	}
	return links, nil
}

func (cli *VanClient) getLocalLinkStatus(namespace string, siteNameMap map[string]string) (map[string]*types.LinkStatus, error) {
	mapLinks := make(map[string]*types.LinkStatus)
	secrets, err := cli.KubeClient.CoreV1().Secrets(namespace).List(metav1.ListOptions{LabelSelector: "skupper.io/type in (connection-token, token-claim)"})
	if err != nil {
		return nil, err
	}

	current, err := cli.getRouterConfig()
	if err != nil {
		return nil, err
	}

	edge := current.IsEdge()
	connections, err := kubeqdr.GetConnections(namespace, cli.KubeClient, cli.RestConfig)
	if err != nil {
		return nil, err
	}

	for _, s := range secrets.Items {
		var connectedTo string
		siteId := s.ObjectMeta.Annotations[types.TokenGeneratedBy]
		connectedTo = siteId[:7] + "-" + siteNameMap[siteId]
		linkStatus := qdr.GetLinkStatus(&s, edge, connections)
		mapLinks[connectedTo] = &linkStatus
	}
	return mapLinks, nil
}
