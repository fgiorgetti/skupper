package adaptor

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"

	iflag "github.com/skupperproject/skupper/internal/flag"
	"github.com/skupperproject/skupper/internal/kube/runtimeapi"
	"github.com/skupperproject/skupper/internal/qdr"
)

const (
	defaultControllerNamespace = "skupper"
	defaultPushTimeout         = 5 * time.Second
)

type RuntimeStateFetcher func() (*qdr.BridgeRuntimeState, error)

type RuntimePusherConfig struct {
	Disabled            bool
	TargetURL           string
	ControllerNamespace string
	TokenFile           string
	CAFile              string
	PushInterval        time.Duration
	RequestTimeout      time.Duration
}

func BoundRuntimePusherConfig(flags *flag.FlagSet) (*RuntimePusherConfig, error) {
	cfg := &RuntimePusherConfig{
		RequestTimeout: defaultPushTimeout,
	}
	err := iflag.BoolVar(flags, &cfg.Disabled, "disable-runtime-api-push", "SKUPPER_RUNTIME_API_PUSH_DISABLE", false, "Set to disable pushing runtime state to skupper-controller.")
	iflag.StringVar(flags, &cfg.TargetURL, "runtime-api-url", "SKUPPER_RUNTIME_API_URL", "", "Full HTTPS URL for pushing runtime state to skupper-controller (auto-derived if empty).")
	iflag.StringVar(flags, &cfg.ControllerNamespace, "controller-namespace", "SKUPPER_CONTROLLER_NAMESPACE", defaultControllerNamespace, "Namespace where skupper-controller runs.")
	iflag.StringVar(flags, &cfg.TokenFile, "runtime-api-token-file", "SKUPPER_RUNTIME_API_TOKEN_FILE", runtimeapi.DefaultRuntimeTokenPath, "Path to projected ServiceAccount token for runtime state pushes.")
	iflag.StringVar(flags, &cfg.CAFile, "runtime-api-ca-file", "SKUPPER_RUNTIME_API_CA_FILE", "", "Optional path to CA bundle for verifying skupper-runtime-api TLS certificate.")

	intervalDefault := runtimeapi.DefaultPushInterval
	if raw, ok := os.LookupEnv("SKUPPER_RUNTIME_API_PUSH_INTERVAL"); ok {
		parsed, parseErr := time.ParseDuration(raw)
		if parseErr != nil && err == nil {
			err = fmt.Errorf("bad value for %q: %w", "SKUPPER_RUNTIME_API_PUSH_INTERVAL", parseErr)
		} else if parseErr == nil {
			intervalDefault = parsed
		}
	}
	flags.DurationVar(&cfg.PushInterval, "runtime-api-push-interval", intervalDefault, "Heartbeat interval for pushing runtime state to skupper-controller.")
	return cfg, err
}

// RuntimeStatePusher asynchronously pushes live router bridge state to skupper-controller.
type RuntimeStatePusher struct {
	targetURL    string
	tokenFile    string
	pushInterval time.Duration
	client       *http.Client
	logger       *slog.Logger

	podName  string
	siteName string
	siteID   string

	mu            sync.RWMutex
	fetchState    RuntimeStateFetcher
	tcpListeners  map[string]qdr.TcpEndpointRuntime
	tcpConnectors map[string]qdr.TcpEndpointRuntime
	hasSnapshot   bool

	notifyCh  chan struct{}
	refreshCh chan struct{}
}

func NewRuntimeStatePusher(cfg RuntimePusherConfig, namespace, podName, siteName, siteID string) (*RuntimeStatePusher, error) {
	targetURL := strings.TrimSpace(cfg.TargetURL)
	if targetURL == "" {
		ctrlNS := strings.TrimSpace(cfg.ControllerNamespace)
		if ctrlNS == "" {
			ctrlNS = defaultControllerNamespace
		}
		targetURL = fmt.Sprintf("https://skupper-runtime-api.%s.svc%s", ctrlNS, runtimeapi.IngestPathForNamespace(namespace))
	}
	tokenFile := strings.TrimSpace(cfg.TokenFile)
	if tokenFile == "" {
		tokenFile = runtimeapi.DefaultRuntimeTokenPath
	}
	interval := cfg.PushInterval
	if interval <= 0 {
		interval = runtimeapi.DefaultPushInterval
	}
	timeout := cfg.RequestTimeout
	if timeout <= 0 {
		timeout = defaultPushTimeout
	}

	tlsCfg, err := buildPusherTLSConfig(cfg.CAFile)
	if err != nil {
		return nil, err
	}

	return &RuntimeStatePusher{
		targetURL:    targetURL,
		tokenFile:    tokenFile,
		pushInterval: interval,
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: tlsCfg,
			},
		},
		logger:    slog.New(slog.Default().Handler()).With(slog.String("component", "kube.adaptor.runtimePusher")),
		podName:   podName,
		siteName:  siteName,
		siteID:    siteID,
		notifyCh:  make(chan struct{}, 1),
		refreshCh: make(chan struct{}, 1),
	}, nil
}

func buildPusherTLSConfig(caFile string) (*tls.Config, error) {
	if strings.TrimSpace(caFile) == "" {
		return &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true,
		}, nil
	}
	pemBytes, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read runtime API CA file %q: %w", caFile, err)
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(pemBytes) {
		return nil, fmt.Errorf("failed to parse CA certificate from %q", caFile)
	}
	return &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    pool,
	}, nil
}

// SetStateFetcher configures the live router query function invoked on each heartbeat and refresh.
func (p *RuntimeStatePusher) SetStateFetcher(fetcher RuntimeStateFetcher) {
	p.mu.Lock()
	p.fetchState = fetcher
	p.mu.Unlock()
}

func (p *RuntimeStatePusher) UpdateState(state qdr.BridgeRuntimeState) {
	p.setSnapshot(state)
	// Coalesce rapid updates without blocking ConfigSync reconciliation.
	select {
	case p.notifyCh <- struct{}{}:
	default:
	}
}

func (p *RuntimeStatePusher) TriggerRefresh() {
	select {
	case p.refreshCh <- struct{}{}:
	default:
	}
}

func (p *RuntimeStatePusher) setSnapshot(state qdr.BridgeRuntimeState) {
	listeners := make(map[string]qdr.TcpEndpointRuntime, len(state.TcpListeners))
	for k, v := range state.TcpListeners {
		listeners[k] = v
	}
	connectors := make(map[string]qdr.TcpEndpointRuntime, len(state.TcpConnectors))
	for k, v := range state.TcpConnectors {
		connectors[k] = v
	}

	p.mu.Lock()
	p.tcpListeners = listeners
	p.tcpConnectors = connectors
	p.hasSnapshot = true
	p.mu.Unlock()
}

func (p *RuntimeStatePusher) Start(stopCh <-chan struct{}) {
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			<-stopCh
			cancel()
		}()

		ticker := time.NewTicker(p.pushInterval)
		defer ticker.Stop()

		for {
			select {
			case <-stopCh:
				return
			case <-p.notifyCh:
				p.pushWithBackoff(ctx)
			case <-p.refreshCh:
				p.refreshFromRouter()
				p.pushWithBackoff(ctx)
			case <-ticker.C:
				p.refreshFromRouter()
				p.pushWithBackoff(ctx)
			}
		}
	}()
}

func (p *RuntimeStatePusher) refreshFromRouter() {
	p.mu.RLock()
	fetcher := p.fetchState
	p.mu.RUnlock()
	if fetcher == nil {
		return
	}
	state, err := fetcher()
	if err != nil {
		p.logger.Warn("Failed to refresh live bridge runtime state from router", slog.Any("error", err))
		return
	}
	if state != nil {
		p.setSnapshot(*state)
	}
}

func (p *RuntimeStatePusher) pushWithBackoff(ctx context.Context) {
	if !p.readyToPush() {
		return
	}
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 500 * time.Millisecond
	b.MaxInterval = 5 * time.Second
	b.MaxElapsedTime = 15 * time.Second
	b.Reset()

	err := backoff.Retry(func() error {
		return p.PushOnce(ctx)
	}, backoff.WithContext(b, ctx))
	if err != nil && ctx.Err() == nil {
		p.logger.Warn("Failed to push runtime state to controller",
			slog.String("url", p.targetURL),
			slog.Any("error", err),
		)
	}
}

func (p *RuntimeStatePusher) readyToPush() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.hasSnapshot
}

func (p *RuntimeStatePusher) snapshotPayload() runtimeapi.RuntimeStatePushRequest {
	p.mu.RLock()
	defer p.mu.RUnlock()

	listeners := make(map[string]qdr.TcpEndpointRuntime, len(p.tcpListeners))
	for k, v := range p.tcpListeners {
		listeners[k] = v
	}
	connectors := make(map[string]qdr.TcpEndpointRuntime, len(p.tcpConnectors))
	for k, v := range p.tcpConnectors {
		connectors[k] = v
	}
	return runtimeapi.RuntimeStatePushRequest{
		PodName:       p.podName,
		SiteName:      p.siteName,
		SiteID:        p.siteID,
		TCPListeners:  listeners,
		TCPConnectors: connectors,
	}
}

// PushOnce reads the projected token fresh from disk and posts the current bridge snapshot.
func (p *RuntimeStatePusher) PushOnce(ctx context.Context) error {
	rawToken, err := os.ReadFile(p.tokenFile)
	if err != nil {
		return fmt.Errorf("failed to read runtime token from %q: %w", p.tokenFile, err)
	}
	token := strings.TrimSpace(string(rawToken))
	if token == "" {
		return errors.New("runtime token file is empty")
	}

	body, err := json.Marshal(p.snapshotPayload())
	if err != nil {
		return backoff.Permanent(fmt.Errorf("failed to marshal runtime state payload: %w", err))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.targetURL, bytes.NewReader(body))
	if err != nil {
		return backoff.Permanent(fmt.Errorf("failed to create runtime state request: %w", err))
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
	statusErr := fmt.Errorf("controller returned HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(respBody)))
	if resp.StatusCode >= 400 && resp.StatusCode < 500 && resp.StatusCode != http.StatusTooManyRequests {
		return backoff.Permanent(statusErr)
	}
	return statusErr
}
