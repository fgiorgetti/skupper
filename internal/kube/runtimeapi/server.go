package runtimeapi

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/skupperproject/skupper/internal/certs"
	iflag "github.com/skupperproject/skupper/internal/flag"
)

const (
	defaultListenAddress = ":8443"
	defaultServiceName   = "skupper-runtime-api"
)

type Config struct {
	Disabled    bool
	Address     string
	TLSCertFile string
	TLSKeyFile  string
	ServiceName string
}

func BoundConfig(flags *flag.FlagSet) (*Config, error) {
	cfg := &Config{}
	err := iflag.BoolVar(flags, &cfg.Disabled, "disable-runtime-api", "SKUPPER_RUNTIME_API_DISABLE", false, "Set to disable the aggregated runtime API server.")
	iflag.StringVar(flags, &cfg.Address, "runtime-api-address", "SKUPPER_RUNTIME_API_ADDRESS", defaultListenAddress, "The HTTPS address for the aggregated runtime API server to listen on.")
	iflag.StringVar(flags, &cfg.TLSCertFile, "runtime-api-tls-cert", "SKUPPER_RUNTIME_API_TLS_CERT", "", "Optional path to TLS certificate for the aggregated runtime API server (self-signed generated if empty).")
	iflag.StringVar(flags, &cfg.TLSKeyFile, "runtime-api-tls-key", "SKUPPER_RUNTIME_API_TLS_KEY", "", "Optional path to TLS private key for the aggregated runtime API server (self-signed generated if empty).")
	iflag.StringVar(flags, &cfg.ServiceName, "runtime-api-service-name", "SKUPPER_RUNTIME_API_SERVICE_NAME", defaultServiceName, "Service name used as SAN when generating a self-signed certificate.")
	return cfg, err
}

// Server implements a lightweight Kubernetes Aggregated API server for runtime.skupper.io/v1alpha1.
// For local evaluation, TLS uses an in-memory self-signed certificate and relies on kube-apiserver front-end authn/authz.
type Server struct {
	config              Config
	controllerNamespace string
	provider            EndpointProvider
	server              *http.Server
	logger              *slog.Logger
	listenerAddr        string
}

func NewServer(cfg *Config, controllerNamespace string, provider EndpointProvider) *Server {
	var c Config
	if cfg != nil {
		c = *cfg
	}
	if c.Address == "" {
		c.Address = defaultListenAddress
	}
	if c.ServiceName == "" {
		c.ServiceName = defaultServiceName
	}
	s := &Server{
		config:              c,
		controllerNamespace: controllerNamespace,
		provider:            provider,
		logger:              slog.New(slog.Default().Handler()).With(slog.String("component", "kube.runtimeapi.server")),
	}
	s.server = &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      http.HandlerFunc(s.serveHTTP),
	}
	return s
}

func (s *Server) Handler() http.Handler {
	return s.server.Handler
}

func (s *Server) Addr() string {
	return s.listenerAddr
}

func (s *Server) Start(stopCh <-chan struct{}) error {
	tlsCert, err := s.loadOrGenerateTLSCert()
	if err != nil {
		return fmt.Errorf("failed to prepare runtime API TLS certificate: %w", err)
	}
	s.server.TLSConfig = &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{tlsCert},
	}

	ln, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		return fmt.Errorf("failed to start runtime API listener on %s: %w", s.config.Address, err)
	}
	s.listenerAddr = ln.Addr().String()
	tlsListener := tls.NewListener(ln, s.server.TLSConfig)

	go func() {
		<-stopCh
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = s.server.Shutdown(ctx)
	}()

	go func() {
		if err := s.server.Serve(tlsListener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("Runtime API server error", slog.Any("error", err))
		}
	}()

	s.logger.Info("Started runtime aggregated API server", slog.String("address", s.listenerAddr))
	return nil
}

func (s *Server) loadOrGenerateTLSCert() (tls.Certificate, error) {
	if s.config.TLSCertFile != "" || s.config.TLSKeyFile != "" {
		return tls.LoadX509KeyPair(s.config.TLSCertFile, s.config.TLSKeyFile)
	}
	hosts := []string{
		s.config.ServiceName,
		"localhost",
		"127.0.0.1",
	}
	if s.controllerNamespace != "" {
		hosts = append(hosts,
			fmt.Sprintf("%s.%s", s.config.ServiceName, s.controllerNamespace),
			fmt.Sprintf("%s.%s.svc", s.config.ServiceName, s.controllerNamespace),
			fmt.Sprintf("%s.%s.svc.cluster.local", s.config.ServiceName, s.controllerNamespace),
		)
	}
	secret, err := certs.GenerateSecret(s.config.ServiceName, s.config.ServiceName, hosts, 0, nil)
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.X509KeyPair(secret.Data["tls.crt"], secret.Data["tls.key"])
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "" {
		path = "/"
	}

	switch path {
	case "/healthz", "/livez", "/readyz":
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
		return
	case "/apis":
		writeJSON(w, http.StatusOK, apiGroupList())
		return
	case "/apis/" + GroupName:
		writeJSON(w, http.StatusOK, apiGroup())
		return
	case "/apis/" + APIVersion:
		writeJSON(w, http.StatusOK, apiResourceList())
		return
	}

	prefix := "/apis/" + APIVersion + "/"
	if !strings.HasPrefix(path, prefix) {
		writeStatusError(w, http.StatusNotFound, metav1.StatusReasonNotFound, fmt.Sprintf("the server could not find the requested resource (%s)", r.URL.Path))
		return
	}

	segments := strings.Split(strings.TrimPrefix(path, prefix), "/")
	switch {
	case len(segments) == 1 && segments[0] == EndpointResource:
		s.handleList(w, r, "")
	case len(segments) == 3 && segments[0] == "namespaces" && segments[2] == EndpointResource:
		s.handleList(w, r, segments[1])
	case len(segments) == 4 && segments[0] == "namespaces" && segments[2] == EndpointResource:
		s.handleGet(w, r, segments[1], segments[3])
	default:
		writeStatusError(w, http.StatusNotFound, metav1.StatusReasonNotFound, fmt.Sprintf("the server could not find the requested resource (%s)", r.URL.Path))
	}
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request, namespace string) {
	if r.Method != http.MethodGet {
		writeStatusError(w, http.StatusMethodNotAllowed, metav1.StatusReasonMethodNotAllowed, "only GET is supported for runtime endpoints")
		return
	}
	if r.URL.Query().Get("watch") == "true" || r.URL.Query().Get("watch") == "1" {
		writeStatusError(w, http.StatusMethodNotAllowed, metav1.StatusReasonMethodNotAllowed, "watch is not supported for runtime endpoints")
		return
	}

	labelSel, err := parseLabelSelector(r.URL.Query().Get("labelSelector"))
	if err != nil {
		writeStatusError(w, http.StatusBadRequest, metav1.StatusReasonBadRequest, fmt.Sprintf("invalid labelSelector: %v", err))
		return
	}
	fieldSel, err := parseFieldSelector(r.URL.Query().Get("fieldSelector"))
	if err != nil {
		writeStatusError(w, http.StatusBadRequest, metav1.StatusReasonBadRequest, fmt.Sprintf("invalid fieldSelector: %v", err))
		return
	}

	if s.provider == nil {
		s.writeListResponse(w, r, []Endpoint{})
		return
	}
	items, err := s.provider.List(r.Context(), namespace)
	if err != nil {
		s.logger.Error("Failed to list runtime endpoints", slog.String("namespace", namespace), slog.Any("error", err))
		writeStatusError(w, http.StatusServiceUnavailable, metav1.StatusReasonServiceUnavailable, err.Error())
		return
	}

	filtered := filterEndpoints(items, labelSel, fieldSel)
	s.writeListResponse(w, r, filtered)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request, namespace, name string) {
	if r.Method != http.MethodGet {
		writeStatusError(w, http.StatusMethodNotAllowed, metav1.StatusReasonMethodNotAllowed, "only GET is supported for runtime endpoints")
		return
	}
	if s.provider == nil {
		writeNotFoundEndpoint(w, name)
		return
	}
	ep, err := s.provider.Get(r.Context(), namespace, name)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			writeNotFoundEndpoint(w, name)
			return
		}
		s.logger.Error("Failed to get runtime endpoint",
			slog.String("namespace", namespace),
			slog.String("name", name),
			slog.Any("error", err),
		)
		writeStatusError(w, http.StatusServiceUnavailable, metav1.StatusReasonServiceUnavailable, err.Error())
		return
	}
	if ep.APIVersion == "" {
		ep.APIVersion = APIVersion
	}
	if ep.Kind == "" {
		ep.Kind = EndpointKind
	}

	if ok, tableVersion := wantsTable(r.Header.Get("Accept")); ok {
		table, err := endpointsToTable([]Endpoint{*ep}, tableVersion, r.URL.Query().Get("includeObject"), time.Now())
		if err != nil {
			writeStatusError(w, http.StatusInternalServerError, metav1.StatusReasonInternalError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, table)
		return
	}
	writeJSON(w, http.StatusOK, ep)
}

func (s *Server) writeListResponse(w http.ResponseWriter, r *http.Request, items []Endpoint) {
	if ok, tableVersion := wantsTable(r.Header.Get("Accept")); ok {
		table, err := endpointsToTable(items, tableVersion, r.URL.Query().Get("includeObject"), time.Now())
		if err != nil {
			writeStatusError(w, http.StatusInternalServerError, metav1.StatusReasonInternalError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, table)
		return
	}
	writeJSON(w, http.StatusOK, NewEndpointList(items))
}

func parseLabelSelector(raw string) (labels.Selector, error) {
	if raw == "" {
		return labels.Everything(), nil
	}
	return labels.Parse(raw)
}

func parseFieldSelector(raw string) (fields.Selector, error) {
	if raw == "" {
		return fields.Everything(), nil
	}
	return fields.ParseSelector(raw)
}

func filterEndpoints(items []Endpoint, labelSel labels.Selector, fieldSel fields.Selector) []Endpoint {
	filtered := make([]Endpoint, 0, len(items))
	for _, ep := range items {
		if !labelSel.Matches(labels.Set(ep.Labels)) {
			continue
		}
		if !fieldSel.Matches(fields.Set{
			"metadata.name":      ep.Name,
			"metadata.namespace": ep.Namespace,
		}) {
			continue
		}
		filtered = append(filtered, ep)
	}
	return filtered
}

func writeJSON(w http.ResponseWriter, statusCode int, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(data)
}

func writeNotFoundEndpoint(w http.ResponseWriter, name string) {
	writeStatusError(w, http.StatusNotFound, metav1.StatusReasonNotFound, fmt.Sprintf("%s.%s %q not found", EndpointResource, GroupName, name))
}

func writeStatusError(w http.ResponseWriter, statusCode int, reason metav1.StatusReason, message string) {
	status := metav1.Status{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Status",
		},
		Status:  metav1.StatusFailure,
		Message: message,
		Reason:  reason,
		Code:    int32(statusCode),
	}
	writeJSON(w, statusCode, status)
}
