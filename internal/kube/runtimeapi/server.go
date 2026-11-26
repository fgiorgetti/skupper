package runtimeapi

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
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
	maxIngestBodyBytes   = 4 << 20
)

type Config struct {
	Disabled    bool
	Address     string
	TLSCertFile string
	TLSKeyFile  string
	ServiceName string
	StateTTL    time.Duration
}

func BoundConfig(flags *flag.FlagSet) (*Config, error) {
	cfg := &Config{}
	err := iflag.BoolVar(flags, &cfg.Disabled, "disable-runtime-api", "SKUPPER_RUNTIME_API_DISABLE", false, "Set to disable the aggregated runtime API server.")
	iflag.StringVar(flags, &cfg.Address, "runtime-api-address", "SKUPPER_RUNTIME_API_ADDRESS", defaultListenAddress, "The HTTPS address for the aggregated runtime API server to listen on.")
	iflag.StringVar(flags, &cfg.TLSCertFile, "runtime-api-tls-cert", "SKUPPER_RUNTIME_API_TLS_CERT", "", "Optional path to TLS certificate for the aggregated runtime API server (self-signed generated if empty).")
	iflag.StringVar(flags, &cfg.TLSKeyFile, "runtime-api-tls-key", "SKUPPER_RUNTIME_API_TLS_KEY", "", "Optional path to TLS private key for the aggregated runtime API server (self-signed generated if empty).")
	iflag.StringVar(flags, &cfg.ServiceName, "runtime-api-service-name", "SKUPPER_RUNTIME_API_SERVICE_NAME", defaultServiceName, "Service name used as SAN when generating a self-signed certificate.")
	ttlDefault, ttlErr := durationEnvVar("SKUPPER_RUNTIME_API_STATE_TTL", DefaultRuntimeStateTTL)
	flags.DurationVar(&cfg.StateTTL, "runtime-api-state-ttl", ttlDefault, "Expiration TTL for pushed kube-adaptor runtime state reports.")
	if err == nil {
		err = ttlErr
	}
	return cfg, err
}

func durationEnvVar(name string, defaultValue time.Duration) (time.Duration, error) {
	if raw, ok := os.LookupEnv(name); ok {
		d, err := time.ParseDuration(raw)
		if err != nil {
			return defaultValue, fmt.Errorf("bad value for %q: %w", name, err)
		}
		return d, nil
	}
	return defaultValue, nil
}

// Server implements the runtime.skupper.io/v1alpha1 Kubernetes Aggregated API and
// the TokenReview-authenticated runtime-state push ingest endpoint.
type Server struct {
	config              Config
	controllerNamespace string
	store               *RuntimeStateStore
	authenticator       IngestAuthenticator
	server              *http.Server
	logger              *slog.Logger
	listenerAddr        string
}

func NewServer(cfg *Config, controllerNamespace string, store *RuntimeStateStore, authenticator IngestAuthenticator) *Server {
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
	if c.StateTTL <= 0 {
		c.StateTTL = DefaultRuntimeStateTTL
	}
	s := &Server{
		config:              c,
		controllerNamespace: controllerNamespace,
		store:               store,
		authenticator:       authenticator,
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

	if s.store != nil {
		s.store.StartPruner(stopCh, 0)
	}

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

	if ns, ok := parseIngestNamespace(path); ok {
		s.handleIngest(w, r, ns)
		return
	}

	prefix := "/apis/" + APIVersion + "/"
	if !strings.HasPrefix(path, prefix) {
		writeStatusError(w, http.StatusNotFound, metav1.StatusReasonNotFound, fmt.Sprintf("the server could not find the requested resource (%s)", r.URL.Path))
		return
	}

	segments := strings.Split(strings.TrimPrefix(path, prefix), "/")
	switch {
	case len(segments) == 1 && segments[0] == AdaptorResource:
		s.handleList(w, r, "")
	case len(segments) == 3 && segments[0] == "namespaces" && segments[2] == AdaptorResource:
		s.handleList(w, r, segments[1])
	case len(segments) == 4 && segments[0] == "namespaces" && segments[2] == AdaptorResource:
		s.handleGet(w, r, segments[1], segments[3])
	default:
		writeStatusError(w, http.StatusNotFound, metav1.StatusReasonNotFound, fmt.Sprintf("the server could not find the requested resource (%s)", r.URL.Path))
	}
}

func parseIngestNamespace(path string) (string, bool) {
	if !strings.HasPrefix(path, IngestPathPrefix) || !strings.HasSuffix(path, IngestPathSuffix) {
		return "", false
	}
	ns := strings.TrimSuffix(strings.TrimPrefix(path, IngestPathPrefix), IngestPathSuffix)
	if ns == "" || strings.Contains(ns, "/") {
		return "", false
	}
	return ns, true
}

func (s *Server) handleIngest(w http.ResponseWriter, r *http.Request, namespace string) {
	if r.Method != http.MethodPost {
		writeStatusError(w, http.StatusMethodNotAllowed, metav1.StatusReasonMethodNotAllowed, "only POST is supported for runtime state ingest")
		return
	}
	if s.authenticator == nil || s.store == nil {
		writeStatusError(w, http.StatusServiceUnavailable, metav1.StatusReasonServiceUnavailable, "runtime state ingest is not configured")
		return
	}

	authRes, err := s.authenticator.AuthenticateAndAuthorize(r.Context(), r.Header.Get("Authorization"), namespace)
	if err != nil {
		var authErr *AuthError
		if errors.As(err, &authErr) {
			writeStatusError(w, authErr.StatusCode, authErr.Reason, authErr.Message)
			return
		}
		writeStatusError(w, http.StatusInternalServerError, metav1.StatusReasonInternalError, err.Error())
		return
	}

	var payload RuntimeStatePushRequest
	dec := json.NewDecoder(io.LimitReader(r.Body, maxIngestBodyBytes))
	if err := dec.Decode(&payload); err != nil {
		writeStatusError(w, http.StatusBadRequest, metav1.StatusReasonBadRequest, fmt.Sprintf("invalid runtime state payload: %v", err))
		return
	}

	podName := authRes.PodName
	if podName == "" {
		podName = strings.TrimSpace(payload.PodName)
	}
	if podName == "" {
		podName = authRes.Username
	}

	s.store.Upsert(namespace, podName, payload)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request, namespace string) {
	if r.Method != http.MethodGet {
		writeStatusError(w, http.StatusMethodNotAllowed, metav1.StatusReasonMethodNotAllowed, "only GET is supported for runtime adaptors")
		return
	}
	if r.URL.Query().Get("watch") == "true" || r.URL.Query().Get("watch") == "1" {
		writeStatusError(w, http.StatusMethodNotAllowed, metav1.StatusReasonMethodNotAllowed, "watch is not supported for runtime adaptors")
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

	if s.store == nil {
		s.writeListResponse(w, r, []Adaptor{})
		return
	}
	items, err := s.store.List(r.Context(), namespace)
	if err != nil {
		s.logger.Error("Failed to list runtime adaptors", slog.String("namespace", namespace), slog.Any("error", err))
		writeStatusError(w, http.StatusInternalServerError, metav1.StatusReasonInternalError, err.Error())
		return
	}

	filtered := filterAdaptors(items, labelSel, fieldSel)
	s.writeListResponse(w, r, filtered)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request, namespace, name string) {
	if r.Method != http.MethodGet {
		writeStatusError(w, http.StatusMethodNotAllowed, metav1.StatusReasonMethodNotAllowed, "only GET is supported for runtime adaptors")
		return
	}
	if s.store == nil {
		writeNotFoundAdaptor(w, name)
		return
	}
	adp, err := s.store.Get(r.Context(), namespace, name)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			writeNotFoundAdaptor(w, name)
			return
		}
		s.logger.Error("Failed to get runtime adaptor",
			slog.String("namespace", namespace),
			slog.String("name", name),
			slog.Any("error", err),
		)
		writeStatusError(w, http.StatusInternalServerError, metav1.StatusReasonInternalError, err.Error())
		return
	}
	if adp.APIVersion == "" {
		adp.APIVersion = APIVersion
	}
	if adp.Kind == "" {
		adp.Kind = AdaptorKind
	}

	if ok, tableVersion := wantsTable(r.Header.Get("Accept")); ok {
		table, err := adaptorsToTable([]Adaptor{*adp}, tableVersion, r.URL.Query().Get("includeObject"), time.Now())
		if err != nil {
			writeStatusError(w, http.StatusInternalServerError, metav1.StatusReasonInternalError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, table)
		return
	}
	writeJSON(w, http.StatusOK, adp)
}

func (s *Server) writeListResponse(w http.ResponseWriter, r *http.Request, items []Adaptor) {
	if ok, tableVersion := wantsTable(r.Header.Get("Accept")); ok {
		table, err := adaptorsToTable(items, tableVersion, r.URL.Query().Get("includeObject"), time.Now())
		if err != nil {
			writeStatusError(w, http.StatusInternalServerError, metav1.StatusReasonInternalError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, table)
		return
	}
	writeJSON(w, http.StatusOK, NewAdaptorList(items))
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

func filterAdaptors(items []Adaptor, labelSel labels.Selector, fieldSel fields.Selector) []Adaptor {
	filtered := make([]Adaptor, 0, len(items))
	for _, adp := range items {
		if !labelSel.Matches(labels.Set(adp.Labels)) {
			continue
		}
		if !fieldSel.Matches(fields.Set{
			"metadata.name":      adp.Name,
			"metadata.namespace": adp.Namespace,
		}) {
			continue
		}
		filtered = append(filtered, adp)
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

func writeNotFoundAdaptor(w http.ResponseWriter, name string) {
	writeStatusError(w, http.StatusNotFound, metav1.StatusReasonNotFound, fmt.Sprintf("%s.%s %q not found", AdaptorResource, GroupName, name))
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
