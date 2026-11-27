package runtimeapi

import (
	"context"
	"errors"
	"sort"
	"strings"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/internal/qdr"
)

var ErrNotFound = errors.New("adaptor not found")

type AdaptorProvider interface {
	List(ctx context.Context, namespace string) ([]Adaptor, error)
	Get(ctx context.Context, namespace, name string) (*Adaptor, error)
}

type SiteNamespaceLister func() []string

type podReportKey struct {
	namespace string
	podName   string
}

type adaptorKey struct {
	namespace string
	name      string
}

type podReport struct {
	adaptors    []Adaptor
	lastUpdated time.Time
}

// RuntimeStateStore caches pushed router runtime state per (namespace, pod) in memory with TTL expiration.
type RuntimeStateStore struct {
	mu             sync.RWMutex
	ttl            time.Duration
	siteNamespaces SiteNamespaceLister
	reports        map[podReportKey]podReport
	firstSeen      map[adaptorKey]time.Time
}

func NewRuntimeStateStore(ttl time.Duration, siteNamespaces SiteNamespaceLister) *RuntimeStateStore {
	if ttl <= 0 {
		ttl = DefaultRuntimeStateTTL
	}
	return &RuntimeStateStore{
		ttl:            ttl,
		siteNamespaces: siteNamespaces,
		reports:        map[podReportKey]podReport{},
		firstSeen:      map[adaptorKey]time.Time{},
	}
}

func (s *RuntimeStateStore) TTL() time.Duration {
	return s.ttl
}

func (s *RuntimeStateStore) Upsert(namespace, podName string, req RuntimeStatePushRequest) {
	now := time.Now().UTC()
	adaptors := buildAdaptorsFromPush(namespace, podName, req)

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range adaptors {
		ak := adaptorKey{namespace: namespace, name: adaptors[i].Name}
		seen, ok := s.firstSeen[ak]
		if !ok {
			seen = now
			s.firstSeen[ak] = seen
		}
		adaptors[i].CreationTimestamp = metav1.NewTime(seen)
	}

	s.reports[podReportKey{namespace: namespace, podName: podName}] = podReport{
		adaptors:    adaptors,
		lastUpdated: now,
	}
	s.gcFirstSeenLocked()
}

func (s *RuntimeStateStore) Prune() {
	now := time.Now().UTC()
	controlled := s.controlledNamespaceSet()

	s.mu.Lock()
	defer s.mu.Unlock()

	for key, report := range s.reports {
		if s.isExpired(report, now) || (controlled != nil && !controlled[key.namespace]) {
			delete(s.reports, key)
		}
	}
	s.gcFirstSeenLocked()
}

func (s *RuntimeStateStore) StartPruner(stopCh <-chan struct{}, interval time.Duration) {
	if interval <= 0 {
		interval = s.ttl / 3
		if interval < time.Second {
			interval = time.Second
		}
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-stopCh:
				return
			case <-ticker.C:
				s.Prune()
			}
		}
	}()
}

func (s *RuntimeStateStore) List(_ context.Context, namespace string) ([]Adaptor, error) {
	now := time.Now().UTC()
	controlled := s.controlledNamespaceSet()
	if namespace != "" && controlled != nil && !controlled[namespace] {
		return []Adaptor{}, nil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	type stampedAdaptor struct {
		adp         Adaptor
		lastUpdated time.Time
	}
	deduped := map[adaptorKey]stampedAdaptor{}

	for key, report := range s.reports {
		if namespace != "" && key.namespace != namespace {
			continue
		}
		if controlled != nil && !controlled[key.namespace] {
			continue
		}
		if s.isExpired(report, now) {
			continue
		}
		for _, adp := range report.adaptors {
			ak := adaptorKey{namespace: key.namespace, name: adp.Name}
			existing, ok := deduped[ak]
			if !ok || report.lastUpdated.After(existing.lastUpdated) {
				deduped[ak] = stampedAdaptor{
					adp:         cloneAdaptor(adp),
					lastUpdated: report.lastUpdated,
				}
			}
		}
	}

	items := make([]Adaptor, 0, len(deduped))
	for _, entry := range deduped {
		items = append(items, entry.adp)
	}
	sortAdaptors(items)
	return items, nil
}

func (s *RuntimeStateStore) Get(ctx context.Context, namespace, name string) (*Adaptor, error) {
	items, err := s.List(ctx, namespace)
	if err != nil {
		return nil, err
	}
	for i := range items {
		if items[i].Name == name {
			adp := items[i]
			return &adp, nil
		}
	}
	return nil, ErrNotFound
}

func (s *RuntimeStateStore) controlledNamespaceSet() map[string]bool {
	if s.siteNamespaces == nil {
		return nil
	}
	namespaces := s.siteNamespaces()
	set := make(map[string]bool, len(namespaces))
	for _, ns := range namespaces {
		set[ns] = true
	}
	return set
}

func (s *RuntimeStateStore) isExpired(report podReport, now time.Time) bool {
	return s.ttl > 0 && now.Sub(report.lastUpdated) > s.ttl
}

func (s *RuntimeStateStore) gcFirstSeenLocked() {
	active := make(map[adaptorKey]struct{})
	for key, report := range s.reports {
		for _, adp := range report.adaptors {
			active[adaptorKey{namespace: key.namespace, name: adp.Name}] = struct{}{}
		}
	}
	for ak := range s.firstSeen {
		if _, ok := active[ak]; !ok {
			delete(s.firstSeen, ak)
		}
	}
}

func buildAdaptorsFromPush(namespace, podName string, req RuntimeStatePushRequest) []Adaptor {
	if len(req.TCPListeners) == 0 && len(req.TCPConnectors) == 0 && len(req.Adaptors) > 0 {
		out := make([]Adaptor, 0, len(req.Adaptors))
		for _, adp := range req.Adaptors {
			adpCopy := cloneAdaptor(adp)
			adpCopy.APIVersion = APIVersion
			adpCopy.Kind = AdaptorKind
			adpCopy.Namespace = namespace
			if adpCopy.Status.PodName == "" {
				adpCopy.Status.PodName = podName
			}
			if adpCopy.Status.SiteName == "" {
				adpCopy.Status.SiteName = req.SiteName
			}
			if adpCopy.Status.SiteID == "" {
				adpCopy.Status.SiteID = req.SiteID
			}
			if adpCopy.Labels == nil {
				adpCopy.Labels = map[string]string{}
			}
			if adpCopy.Status.Type != "" {
				adpCopy.Labels[AdaptorTypeLabel] = string(adpCopy.Status.Type)
			}
			out = append(out, adpCopy)
		}
		sortAdaptors(out)
		return out
	}

	usedNames := make(map[string]struct{}, len(req.TCPListeners)+len(req.TCPConnectors))
	items := make([]Adaptor, 0, len(req.TCPListeners)+len(req.TCPConnectors))

	for _, key := range sortedEndpointKeys(req.TCPListeners) {
		items = append(items, toAPIAdaptor(namespace, podName, req.SiteName, req.SiteID, AdaptorTypeTCPListener, key, req.TCPListeners[key], usedNames))
	}
	for _, key := range sortedEndpointKeys(req.TCPConnectors) {
		items = append(items, toAPIAdaptor(namespace, podName, req.SiteName, req.SiteID, AdaptorTypeTCPConnector, key, req.TCPConnectors[key], usedNames))
	}
	sortAdaptors(items)
	return items
}

func toAPIAdaptor(namespace, podName, siteName, defaultSiteID string, adpType AdaptorType, key string, ep qdr.TcpEndpointRuntime, usedNames map[string]struct{}) Adaptor {
	rawName := ep.Name
	if rawName == "" {
		rawName = key
	}
	objName := sanitizeAdaptorName(rawName)
	if _, exists := usedNames[objName]; exists {
		suffix := "-listener"
		if adpType == AdaptorTypeTCPConnector {
			suffix = "-connector"
		}
		objName += suffix
	}
	usedNames[objName] = struct{}{}

	var annotations map[string]string
	if rawName != objName {
		annotations = map[string]string{
			RouterNameAnnotation: rawName,
		}
	}

	siteID := ep.SiteId
	if siteID == "" {
		siteID = defaultSiteID
	}

	return Adaptor{
		TypeMeta: metav1.TypeMeta{
			APIVersion: APIVersion,
			Kind:       AdaptorKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        objName,
			Namespace:   namespace,
			Labels:      map[string]string{AdaptorTypeLabel: string(adpType)},
			Annotations: annotations,
		},
		Status: AdaptorStatus{
			Type:              adpType,
			OperStatus:        ep.OperStatus,
			Address:           ep.Address,
			Host:              ep.Host,
			Port:              ep.Port,
			PodName:           podName,
			ConnectionsOpened: ep.ConnectionsOpened,
			ConnectionsClosed: ep.ConnectionsClosed,
			BytesIn:           ep.BytesIn,
			BytesOut:          ep.BytesOut,
			SiteID:            siteID,
			SiteName:          siteName,
			SslProfile:        ep.SslProfile,
			Observer:          ep.Observer,
			ProcessID:         ep.ProcessID,
		},
	}
}

func sortedEndpointKeys(m map[string]qdr.TcpEndpointRuntime) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sanitizeAdaptorName(raw string) string {
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
		return "adaptor"
	}
	return res
}

func cloneAdaptor(adp Adaptor) Adaptor {
	out := adp
	if adp.Labels != nil {
		out.Labels = make(map[string]string, len(adp.Labels))
		for k, v := range adp.Labels {
			out.Labels[k] = v
		}
	}
	if adp.Annotations != nil {
		out.Annotations = make(map[string]string, len(adp.Annotations))
		for k, v := range adp.Annotations {
			out.Annotations[k] = v
		}
	}
	return out
}

func sortAdaptors(items []Adaptor) {
	sort.Slice(items, func(i, j int) bool {
		if items[i].Namespace != items[j].Namespace {
			return items[i].Namespace < items[j].Namespace
		}
		return items[i].Name < items[j].Name
	})
}
