package podman

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	runtime "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/skupperproject/skupper/pkg/config"
)

const (
	DEFAULT_BASE_PATH = "/v4.0.0"
)

var (
	formats = strfmt.NewFormats()
)

type PodmanRestClient struct {
	RestClient *runtime.Runtime
}

func NewPodmanClient(endpoint, basePath string) (*PodmanRestClient, error) {
	var err error
	var sockFile bool

	if endpoint == "" {
		endpoint = fmt.Sprintf("%s/podman/podman.sock", config.GetRuntimeDir())
	}

	var u = &url.URL{
		Scheme: "http",
	}
	if strings.HasPrefix(endpoint, "/") {
		u.Host = "unix"
		sockFile = true
	} else {
		host := endpoint
		match, _ := regexp.Match(`http[s]*://`, []byte(host))
		if !match {
			host = "http://" + host
		}
		u, err = url.Parse(host)
		if err != nil {
			return nil, err
		}
	}

	hostPort := u.Hostname()
	if u.Port() != "" {
		hostPort = net.JoinHostPort(u.Hostname(), u.Port())
	}
	if basePath == "" {
		basePath = DEFAULT_BASE_PATH
	}
	c := runtime.New(hostPort, basePath, []string{u.Scheme})
	if u.Scheme == "https" {
		ct := c.Transport.(*http.Transport)
		if ct.TLSClientConfig != nil {
			ct.TLSClientConfig = &tls.Config{}
		}
		ct.TLSClientConfig.InsecureSkipVerify = true
	}
	if sockFile {
		_, err := os.Stat(endpoint)
		if err != nil {
			return nil, fmt.Errorf("invalid sock file: %v", err)
		}
		ct := c.Transport.(*http.Transport)
		ct.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial("unix", endpoint)
		}
	}

	cli := &PodmanRestClient{
		RestClient: c,
	}
	return cli, nil
}

// boolTrue returns a true bool pointer (for false, just use new(bool))
func boolTrue() *bool {
	b := true
	return &b
}
