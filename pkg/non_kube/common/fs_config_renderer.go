package common

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"syscall"

	"github.com/google/uuid"
	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/pkg/certs"
	"github.com/skupperproject/skupper/pkg/config"
	"github.com/skupperproject/skupper/pkg/non_kube/apis"
	"github.com/skupperproject/skupper/pkg/qdr"
	"github.com/skupperproject/skupper/pkg/version"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ConfigRouterPath       = "config/router"
	CertificatesCaPath     = "certificates/ca"
	CertificatesClientPath = "certificates/client"
	CertificatesServerPath = "certificates/server"
	LoadedSiteStatePath    = "loaded/state"
	RuntimeSiteStatePath   = "runtime/state"
)

var (
	configurationDirectories = []string{
		ConfigRouterPath,
		CertificatesCaPath,
		CertificatesClientPath,
		CertificatesServerPath,
		LoadedSiteStatePath,
		RuntimeSiteStatePath,
	}
)

const (
	DefaultSslProfileBasePath = "${SSL_PROFILE_BASE_PATH}"
)

func getDefaultOutputPath(siteName string) string {
	if IsRunningInContainer() {
		return path.Join("/output", "sites", siteName)
	}
	return path.Join(config.GetDataHome(), "sites", siteName)
}

type FileSystemConfigurationRenderer struct {
	// OutputPath destination directory where configuration will be rendered into
	OutputPath string
	// SslProfileBasePath path where configuration will be read from in runtime
	SslProfileBasePath string
	// Force creation even if directory already exists
	Force        bool
	RouterConfig qdr.RouterConfig
}

// Render simply renders the given site state as configuration files.
func (c *FileSystemConfigurationRenderer) Render(siteState apis.SiteState) error {
	var err error
	// Set the default output path
	if c.OutputPath == "" {
		c.OutputPath = getDefaultOutputPath(siteState.Site.Name)
	}
	if c.SslProfileBasePath == "" {
		c.SslProfileBasePath = DefaultSslProfileBasePath
	}
	// Proceed only if output path does not exist
	outputDir, err := os.Open(c.OutputPath)
	if err == nil {
		if !c.Force {
			return fmt.Errorf("output directory %s already exists", c.OutputPath)
		}
		outputDirStat, err := outputDir.Stat()
		if err != nil {
			return fmt.Errorf("failed to check if output directory exists (%s): %w", c.OutputPath, err)
		}
		if !outputDirStat.IsDir() {
			return fmt.Errorf("output path must be a directory (%s)", c.OutputPath)
		}
	} else {
		var pathErr *os.PathError
		if ok := errors.As(err, &pathErr); ok && !errors.Is(pathErr.Err, syscall.ENOENT) {
			return fmt.Errorf("unable to use output path %s: %v", c.OutputPath, err)
		}
	}
	// Creating internal configuration directories
	for _, dir := range configurationDirectories {
		configDir := path.Join(c.OutputPath, dir)
		err := os.MkdirAll(configDir, 0755)
		if err != nil && !os.IsExist(err) {
			return fmt.Errorf("unable to create configuration directory %s: %v", configDir, err)
		}
	}
	// Creating the router config
	err = c.createRouterConfig(siteState)
	if err != nil {
		return fmt.Errorf("unable to create router config: %v", err)
	}

	// Creating the certificates
	err = c.createTlsCertificates(siteState)
	if err != nil {
		return fmt.Errorf("unable to create tls certificates: %v", err)
	}
	return nil
}

func (c *FileSystemConfigurationRenderer) createRouterConfig(siteState apis.SiteState) error {
	siteId := uuid.New().String()
	c.RouterConfig = qdr.InitialConfig(siteState.Site.Name, siteId, version.Version, !siteState.IsInterior(), 3)

	// Process order
	// - LinkAccess
	// - Listener
	// - Connector

	// LinkAccess
	for name, la := range siteState.LinkAccesses {
		for _, role := range la.Spec.Roles {
			listenerName := fmt.Sprintf("%s-%s", name, role.Role)
			// TODO LinkAccess.Spec.Options[bindIp] ?
			host := getOption(la.Spec.Options, "bindIp", "0.0.0.0")
			c.RouterConfig.AddListener(qdr.Listener{
				Name: listenerName,
				// TODO LinkAccess.spec.roles[].role should be a dropdown/choice on UI
				Role:             qdr.Role(role.Role),
				Host:             host,
				Port:             int32(role.Port),
				SslProfile:       la.Spec.TlsCredentials,
				SaslMechanisms:   "EXTERNAL",
				AuthenticatePeer: true,
				MaxFrameSize:     types.RouterMaxFrameSizeDefault,
				MaxSessionFrames: types.RouterMaxSessionFramesDefault,
			})
		}
		c.RouterConfig.AddSslProfileWithPath(path.Join(c.SslProfileBasePath, "certificates/server"), qdr.SslProfile{
			Name: la.Spec.TlsCredentials,
		})
	}
	// TODO Render inter-router or edge connectors
	// TODO TCP Listeners and Connectors cannot yet handle names (we need the proxy container first - iptables)
	// TCP Listener
	for name, listener := range siteState.Listeners {
		listenerName := fmt.Sprintf("listener-%s", name)
		c.RouterConfig.Bridges.AddTcpListener(qdr.TcpEndpoint{
			Name:       listenerName,
			Host:       listener.Spec.Host,
			Port:       strconv.Itoa(listener.Spec.Port),
			Address:    listener.Spec.RoutingKey,
			SiteId:     siteId,
			SslProfile: listener.Spec.TlsCredentials,
		})
		if listener.Spec.TlsCredentials != "" {
			c.RouterConfig.AddSslProfileWithPath(path.Join(c.SslProfileBasePath, "certificates/server"), qdr.SslProfile{
				Name: listener.Spec.TlsCredentials,
			})
		}
	}
	// TCP Connector
	for name, connector := range siteState.Connectors {
		connectorName := fmt.Sprintf("connector-%s", name)
		c.RouterConfig.Bridges.AddTcpConnector(qdr.TcpEndpoint{
			Name:       connectorName,
			Host:       connector.Spec.Host,
			Port:       strconv.Itoa(connector.Spec.Port),
			Address:    connector.Spec.RoutingKey,
			SiteId:     siteId,
			SslProfile: connector.Spec.TlsCredentials,
		})
		if connector.Spec.TlsCredentials != "" {
			c.RouterConfig.AddSslProfileWithPath(path.Join(c.SslProfileBasePath, "certificates/client"), qdr.SslProfile{
				Name: connector.Spec.TlsCredentials,
			})
		}
	}
	// Log (static for now) TODO use site specific options to configure logging
	c.RouterConfig.SetLogLevel("ROUTER_CORE", "error+")

	// Saving router config
	routerConfigJson, err := qdr.MarshalRouterConfig(c.RouterConfig)
	if err != nil {
		return fmt.Errorf("unable to marshal router config: %v", err)
	}
	routerConfigFileName := path.Join(c.OutputPath, "config/router", "skrouterd.json")
	err = os.WriteFile(routerConfigFileName, []byte(routerConfigJson), 0644)
	if err != nil {
		return fmt.Errorf("unable to write router config file: %v", err)
	}
	return nil
}

func (c *FileSystemConfigurationRenderer) createTlsCertificates(siteState apis.SiteState) error {
	var err error
	writeSecretFiles := func(basePath string, secret corev1.Secret) error {
		baseDir, err := os.Open(basePath)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(basePath, 0755)
				if err != nil {
					return fmt.Errorf("unable to create directory %s: %v", basePath, err)
				}
			}
		} else {
			baseDirStat, err := baseDir.Stat()
			if err != nil {
				return fmt.Errorf("unable to verify directory %s: %v", basePath, err)
			}
			if !baseDirStat.IsDir() {
				return fmt.Errorf("%s is not a directory", basePath)
			}
		}
		for fileName, data := range secret.Data {
			certFileName := path.Join(basePath, fileName)
			if _, err := os.Open(certFileName); err == nil {
				// ignoring existing certificate
				continue
			}
			err = os.WriteFile(certFileName, data, 0640)
			if err != nil {
				return fmt.Errorf("error writing %s: %v", certFileName, err)
			}
		}
		return nil
	}
	// create certificate authorities first
	for name, certificate := range siteState.Certificates {
		if certificate.Spec.Signing == false {
			continue
		}
		secret := certs.GenerateCASecret(name, certificate.Spec.Subject)
		caPath := path.Join(c.OutputPath, "certificates/ca", name)
		err = writeSecretFiles(caPath, secret)
		if err != nil {
			return err
		}
	}
	// generate all other certificates now
	for name, certificate := range siteState.Certificates {
		var purpose string
		var secret corev1.Secret
		var caSecret *corev1.Secret
		if certificate.Spec.Ca != "" {
			caSecret, err = c.loadCASecret(certificate.Spec.Ca)
			if err != nil {
				return fmt.Errorf("unable to load CA secret %s: %v", certificate.Spec.Ca, err)
			}
		}
		if certificate.Spec.Client {
			// TODO Verify if client certificate is generated correctly
			purpose = "client"
			secret = certs.GenerateSecret(name, certificate.Spec.Subject, strings.Join(certificate.Spec.Hosts, ","), caSecret)
			if connectJson := c.connectJson(siteState); connectJson != nil {
				secret.Data["connect.json"] = []byte(*connectJson)
			}
		} else if certificate.Spec.Server {
			purpose = "server"
			secret = certs.GenerateSecret(name, certificate.Spec.Subject, strings.Join(certificate.Spec.Hosts, ","), caSecret)
		} else {
			continue
		}
		certPath := path.Join(c.OutputPath, "certificates", purpose, name)
		err = writeSecretFiles(certPath, secret)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *FileSystemConfigurationRenderer) connectJson(siteState apis.SiteState) *string {
	var host string
	port := 0
	for _, la := range siteState.LinkAccesses {
		for _, role := range la.Spec.Roles {
			if role.Role == "normal" {
				port = role.Port
				// TODO adjust once model is refined
				host = getOption(la.Spec.Options, "bindIp", "127.0.0.1")
			}
		}
		if port > 0 {
			break
		}
	}
	if port == 0 {
		return nil
	}
	// TODO not sure if it will be needed, but just in case we are using 127.0.0.1 as the target host
	content := `
{
    "scheme": "amqps",
    "host": "` + host + `",
    "port": "` + strconv.Itoa(port) + `",
    "tls": {
        "ca": "/etc/messaging/ca.crt",
        "cert": "/etc/messaging/tls.crt",
        "key": "/etc/messaging/tls.key",
        "verify": true
    }
}
`
	return &content
}

func (c *FileSystemConfigurationRenderer) loadCASecret(name string) (*corev1.Secret, error) {
	caPath := path.Join(c.OutputPath, "certificates/ca", name)
	var secret *corev1.Secret
	caDir, err := os.Open(caPath)
	if err != nil {
		return nil, err
	}
	caDirStat, err := caDir.Stat()
	if err != nil {
		return nil, fmt.Errorf("error checking CA dir stats %s: %v", caPath, err)
	}
	if !caDirStat.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", caPath)
	}
	files, err := caDir.ReadDir(0)
	if err != nil {
		return nil, fmt.Errorf("error reading files in %s: %v", caPath, err)
	}
	secret = &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Data: map[string][]byte{},
	}
	for _, file := range files {
		fileName := path.Join(caPath, file.Name())
		fileContent, err := os.ReadFile(fileName)
		if err != nil {
			return nil, fmt.Errorf("error reading %s: %v", fileName, err)
		}
		secret.Data[file.Name()] = fileContent
	}
	return secret, nil
}

func getOption(m map[string]string, key, defaultValue string) string {
	if m != nil {
		if value, ok := m[key]; ok {
			return value
		}
	}
	return defaultValue
}
