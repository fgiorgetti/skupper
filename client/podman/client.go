package podman

import (
	"bytes"
	"context"
	"time"

	"github.com/skupperproject/skupper/api/types"
	corev1 "k8s.io/api/core/v1"
)

const (
	DefaultNetworkDriver = "bridge"
)

func NewClient(ctx context.Context) (*VanClient, error) {
	return nil, nil
}

type VanClient struct {
	PodmanClient *PodmanRestClient
}

func (c *VanClient) RouterCreate(ctx context.Context, options types.SiteConfig) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) RouterInspect(ctx context.Context) (*types.RouterInspectResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) RouterInspectNamespace(ctx context.Context, namespace string) (*types.RouterInspectResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) RouterRemove(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) RouterUpdateVersion(ctx context.Context, hup bool) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) RouterUpdateVersionInNamespace(ctx context.Context, hup bool, namespace string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ConnectorCreateFromFile(ctx context.Context, secretFile string, options types.ConnectorCreateOptions) (*corev1.Secret, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ConnectorCreateSecretFromData(ctx context.Context, secretData []byte, options types.ConnectorCreateOptions) (*corev1.Secret, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ConnectorCreate(ctx context.Context, secret *corev1.Secret, options types.ConnectorCreateOptions) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ConnectorInspect(ctx context.Context, name string) (*types.LinkStatus, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ConnectorList(ctx context.Context) ([]types.LinkStatus, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ConnectorRemove(ctx context.Context, options types.ConnectorRemoveOptions) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ConnectorTokenCreate(ctx context.Context, subject string, namespace string) (*corev1.Secret, bool, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ConnectorTokenCreateFile(ctx context.Context, subject string, secretFile string) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) TokenClaimCreate(ctx context.Context, name string, password []byte, expiry time.Duration, uses int) (*corev1.Secret, bool, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) TokenClaimCreateFile(ctx context.Context, name string, password []byte, expiry time.Duration, uses int, secretFile string) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ServiceInterfaceCreate(ctx context.Context, service *types.ServiceInterface) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ServiceInterfaceInspect(ctx context.Context, address string) (*types.ServiceInterface, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ServiceInterfaceList(ctx context.Context) ([]*types.ServiceInterface, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ServiceInterfaceRemove(ctx context.Context, address string) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ServiceInterfaceUpdate(ctx context.Context, service *types.ServiceInterface) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ServiceInterfaceBind(ctx context.Context, service *types.ServiceInterface, targetType string, targetName string, protocol string, targetPorts map[int]int) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GetHeadlessServiceConfiguration(targetName string, protocol string, address string, ports []int) (*types.ServiceInterface, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) ServiceInterfaceUnbind(ctx context.Context, targetType string, targetName string, address string, deleteIfNoTargets bool) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayBind(ctx context.Context, gatewayName string, endpoint types.GatewayEndpoint) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayUnbind(ctx context.Context, gatewayName string, endpoint types.GatewayEndpoint) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayExpose(ctx context.Context, gatewayName string, gatewayType string, endpoint types.GatewayEndpoint) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayUnexpose(ctx context.Context, gatewayName string, endpoint types.GatewayEndpoint, deleteLast bool) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayForward(ctx context.Context, gatewayName string, endpoint types.GatewayEndpoint) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayUnforward(ctx context.Context, gatewayName string, endpoint types.GatewayEndpoint) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayInit(ctx context.Context, gatewayName string, gatewayType string, configFile string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayDownload(ctx context.Context, gatewayName string, downloadPath string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayExportConfig(ctx context.Context, targetGatewayName string, exportGatewayName string, exportPath string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayGenerateBundle(ctx context.Context, configFile string, bundlePath string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayInspect(ctx context.Context, gatewayName string) (*types.GatewayInspectResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayList(ctx context.Context) ([]*types.GatewayInspectResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GatewayRemove(ctx context.Context, gatewayName string) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) SiteConfigCreate(ctx context.Context, spec types.SiteConfigSpec) (*types.SiteConfig, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) SiteConfigUpdate(ctx context.Context, spec types.SiteConfigSpec) ([]string, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) SiteConfigInspect(ctx context.Context, input *corev1.ConfigMap) (*types.SiteConfig, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) SiteConfigRemove(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) SkupperDump(ctx context.Context, tarName string, version string, kubeConfigPath string, kubeConfigContext string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) SkupperEvents(verbose bool) (*bytes.Buffer, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) SkupperCheckService(service string, verbose bool) (*bytes.Buffer, error) {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GetNamespace() string {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GetVersion(component string, name string) string {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) GetIngressDefault() string {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) RevokeAccess(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (c *VanClient) NetworkStatus() ([]*types.SiteInfo, error) {
	// TODO implement me
	panic("implement me")
}
