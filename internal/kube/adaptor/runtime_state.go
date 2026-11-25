package adaptor

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/internal/kube/runtimeapi"
	"github.com/skupperproject/skupper/internal/qdr"
)

const endpointTypeLabel = "runtime.skupper.io/type"

type trackedEndpoint struct {
	endpoint  qdr.TcpEndpoint
	firstSeen time.Time
}

// RuntimeStateStore holds in-memory router runtime state (TCP listeners and connectors) for a site.
type RuntimeStateStore struct {
	mu            sync.RWMutex
	namespace     string
	siteName      string
	siteID        string
	tcpListeners  map[string]trackedEndpoint
	tcpConnectors map[string]trackedEndpoint
}

func NewRuntimeStateStore(namespace, siteName, siteID string) *RuntimeStateStore {
	return &RuntimeStateStore{
		namespace:     namespace,
		siteName:      siteName,
		siteID:        siteID,
		tcpListeners:  map[string]trackedEndpoint{},
		tcpConnectors: map[string]trackedEndpoint{},
	}
}

func (s *RuntimeStateStore) SetTCPListeners(listeners map[string]qdr.TcpEndpoint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tcpListeners = replaceTracked(s.tcpListeners, listeners, time.Now().UTC())
}

func (s *RuntimeStateStore) SetTCPConnectors(connectors map[string]qdr.TcpEndpoint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tcpConnectors = replaceTracked(s.tcpConnectors, connectors, time.Now().UTC())
}

func (s *RuntimeStateStore) SetBridges(bridges qdr.BridgeConfig) {
	now := time.Now().UTC()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tcpListeners = replaceTracked(s.tcpListeners, bridges.TcpListeners, now)
	s.tcpConnectors = replaceTracked(s.tcpConnectors, bridges.TcpConnectors, now)
}

func replaceTracked(prev map[string]trackedEndpoint, next map[string]qdr.TcpEndpoint, now time.Time) map[string]trackedEndpoint {
	updated := make(map[string]trackedEndpoint, len(next))
	for key, ep := range next {
		if ep.Name == "" {
			ep.Name = key
		}
		firstSeen := now
		if existing, ok := prev[key]; ok {
			firstSeen = existing.firstSeen
		}
		updated[key] = trackedEndpoint{
			endpoint:  ep,
			firstSeen: firstSeen,
		}
	}
	return updated
}

func (s *RuntimeStateStore) ListEndpoints() []runtimeapi.Endpoint {
	s.mu.RLock()
	defer s.mu.RUnlock()

	usedNames := make(map[string]struct{}, len(s.tcpListeners)+len(s.tcpConnectors))
	items := make([]runtimeapi.Endpoint, 0, len(s.tcpListeners)+len(s.tcpConnectors))

	for _, key := range sortedKeys(s.tcpListeners) {
		tracked := s.tcpListeners[key]
		ep := s.toAPIEndpoint(runtimeapi.EndpointTypeTCPListener, key, tracked, usedNames)
		items = append(items, ep)
	}
	for _, key := range sortedKeys(s.tcpConnectors) {
		tracked := s.tcpConnectors[key]
		ep := s.toAPIEndpoint(runtimeapi.EndpointTypeTCPConnector, key, tracked, usedNames)
		items = append(items, ep)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Name < items[j].Name
	})
	return items
}

func (s *RuntimeStateStore) GetEndpoint(name string) (*runtimeapi.Endpoint, bool) {
	for _, ep := range s.ListEndpoints() {
		if ep.Name == name {
			copyEp := ep
			return &copyEp, true
		}
	}
	return nil, false
}

func (s *RuntimeStateStore) toAPIEndpoint(epType runtimeapi.EndpointType, key string, tracked trackedEndpoint, usedNames map[string]struct{}) runtimeapi.Endpoint {
	rawName := tracked.endpoint.Name
	if rawName == "" {
		rawName = key
	}
	objName := sanitizeEndpointName(rawName)
	if _, exists := usedNames[objName]; exists {
		suffix := "-listener"
		if epType == runtimeapi.EndpointTypeTCPConnector {
			suffix = "-connector"
		}
		objName += suffix
	}
	usedNames[objName] = struct{}{}

	var annotations map[string]string
	if rawName != objName {
		annotations = map[string]string{
			runtimeapi.RouterNameAnnotation: rawName,
		}
	}

	siteID := tracked.endpoint.SiteId
	if siteID == "" {
		siteID = s.siteID
	}

	return runtimeapi.Endpoint{
		TypeMeta: metav1.TypeMeta{
			APIVersion: runtimeapi.APIVersion,
			Kind:       runtimeapi.EndpointKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:              objName,
			Namespace:         s.namespace,
			CreationTimestamp: metav1.NewTime(tracked.firstSeen),
			Labels: map[string]string{
				endpointTypeLabel: string(epType),
			},
			Annotations: annotations,
		},
		Status: runtimeapi.EndpointStatus{
			Type:       epType,
			Address:    tracked.endpoint.Address,
			Host:       tracked.endpoint.Host,
			Port:       tracked.endpoint.Port,
			SiteID:     siteID,
			SiteName:   s.siteName,
			SslProfile: tracked.endpoint.SslProfile,
			Observer:   tracked.endpoint.Observer,
			ProcessID:  tracked.endpoint.ProcessID,
		},
	}
}

func sortedKeys(m map[string]trackedEndpoint) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sanitizeEndpointName(raw string) string {
	lower := strings.ToLower(strings.TrimSpace(raw))
	var b strings.Builder
	b.Grow(len(lower))
	lastDash := false
	for _, r := range lower {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '.' {
			b.WriteRune(r)
			lastDash = false
		} else if !lastDash && b.Len() > 0 {
			b.WriteByte('-')
			lastDash = true
		}
	}
	res := strings.Trim(b.String(), "-.")
	if res == "" {
		return "endpoint"
	}
	return res
}

// RuntimeStateHandler serves in-memory runtime endpoints over HTTP for the controller's aggregated API server.
func RuntimeStateHandler(store *RuntimeStateStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		path := strings.TrimSuffix(r.URL.Path, "/")
		if path == runtimeapi.AdaptorEndpointsPath {
			writeAdaptorJSON(w, http.StatusOK, runtimeapi.NewEndpointList(store.ListEndpoints()))
			return
		}
		prefix := runtimeapi.AdaptorEndpointsPath + "/"
		if strings.HasPrefix(path, prefix) {
			name := strings.TrimPrefix(path, prefix)
			if name == "" || strings.Contains(name, "/") {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			ep, ok := store.GetEndpoint(name)
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			writeAdaptorJSON(w, http.StatusOK, ep)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
}

func writeAdaptorJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(data)
}
