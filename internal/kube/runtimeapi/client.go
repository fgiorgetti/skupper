package runtimeapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"sort"
	"sync"
	"time"
)

var ErrNotFound = errors.New("endpoint not found")

type EndpointProvider interface {
	List(ctx context.Context, namespace string) ([]Endpoint, error)
	Get(ctx context.Context, namespace, name string) (*Endpoint, error)
}

type SiteNamespaceLister func() []string

type AdaptorProvider struct {
	HTTPClient     *http.Client
	ServiceName    string
	Port           int
	SiteNamespaces SiteNamespaceLister
	BaseURL        func(namespace string) string
	Logger         *slog.Logger
}

func NewAdaptorProvider(siteNamespaces SiteNamespaceLister) *AdaptorProvider {
	return &AdaptorProvider{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		ServiceName:    "skupper-router-local",
		Port:           DefaultAdaptorPort,
		SiteNamespaces: siteNamespaces,
		Logger:         slog.New(slog.Default().Handler()).With(slog.String("component", "kube.runtimeapi.client")),
	}
}

func (p *AdaptorProvider) baseURLFor(namespace string) string {
	if p.BaseURL != nil {
		return p.BaseURL(namespace)
	}
	serviceName := p.ServiceName
	if serviceName == "" {
		serviceName = "skupper-router-local"
	}
	port := p.Port
	if port == 0 {
		port = DefaultAdaptorPort
	}
	return fmt.Sprintf("http://%s.%s.svc:%d", serviceName, namespace, port)
}

func (p *AdaptorProvider) httpClient() *http.Client {
	if p.HTTPClient != nil {
		return p.HTTPClient
	}
	return http.DefaultClient
}

func (p *AdaptorProvider) logger() *slog.Logger {
	if p.Logger != nil {
		return p.Logger
	}
	return slog.Default()
}

func (p *AdaptorProvider) hasSiteInNamespace(namespace string) bool {
	if p.SiteNamespaces == nil {
		return true
	}
	for _, ns := range p.SiteNamespaces() {
		if ns == namespace {
			return true
		}
	}
	return false
}

func (p *AdaptorProvider) List(ctx context.Context, namespace string) ([]Endpoint, error) {
	if namespace != "" {
		if !p.hasSiteInNamespace(namespace) {
			return []Endpoint{}, nil
		}
		items, err := p.fetchNamespaceList(ctx, namespace)
		if err != nil {
			return nil, err
		}
		sortEndpoints(items)
		return items, nil
	}

	if p.SiteNamespaces == nil {
		return []Endpoint{}, nil
	}
	namespaces := p.SiteNamespaces()
	if len(namespaces) == 0 {
		return []Endpoint{}, nil
	}

	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		results = make([]Endpoint, 0)
	)
	for _, ns := range namespaces {
		wg.Add(1)
		go func(targetNS string) {
			defer wg.Done()
			items, err := p.fetchNamespaceList(ctx, targetNS)
			if err != nil {
				p.logger().Warn("Failed to query kube-adaptor runtime state",
					slog.String("namespace", targetNS),
					slog.Any("error", err),
				)
				return
			}
			mu.Lock()
			results = append(results, items...)
			mu.Unlock()
		}(ns)
	}
	wg.Wait()
	sortEndpoints(results)
	return results, nil
}

func (p *AdaptorProvider) Get(ctx context.Context, namespace, name string) (*Endpoint, error) {
	if !p.hasSiteInNamespace(namespace) {
		return nil, ErrNotFound
	}
	endpointURL := fmt.Sprintf("%s%s/%s", p.baseURLFor(namespace), AdaptorEndpointsPath, url.PathEscape(name))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.httpClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query kube-adaptor in namespace %q: %w", namespace, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("kube-adaptor in namespace %q returned status %d", namespace, resp.StatusCode)
	}

	var ep Endpoint
	if err := json.NewDecoder(resp.Body).Decode(&ep); err != nil {
		return nil, fmt.Errorf("failed to decode endpoint from namespace %q: %w", namespace, err)
	}
	if ep.Namespace == "" {
		ep.Namespace = namespace
	}
	ep.APIVersion = APIVersion
	ep.Kind = EndpointKind
	return &ep, nil
}

func (p *AdaptorProvider) fetchNamespaceList(ctx context.Context, namespace string) ([]Endpoint, error) {
	listURL := p.baseURLFor(namespace) + AdaptorEndpointsPath
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, listURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.httpClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query kube-adaptor in namespace %q: %w", namespace, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("kube-adaptor in namespace %q returned status %d", namespace, resp.StatusCode)
	}

	var list EndpointList
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, fmt.Errorf("failed to decode endpoint list from namespace %q: %w", namespace, err)
	}
	for i := range list.Items {
		if list.Items[i].Namespace == "" {
			list.Items[i].Namespace = namespace
		}
		list.Items[i].APIVersion = APIVersion
		list.Items[i].Kind = EndpointKind
	}
	return list.Items, nil
}

func sortEndpoints(items []Endpoint) {
	sort.Slice(items, func(i, j int) bool {
		if items[i].Namespace != items[j].Namespace {
			return items[i].Namespace < items[j].Namespace
		}
		return items[i].Name < items[j].Name
	})
}
