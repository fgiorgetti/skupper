package runtimeapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	authenticationv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const podNameExtraKey = "authentication.kubernetes.io/pod-name"

type AuthResult struct {
	Username  string
	Namespace string
	PodName   string
}

type AuthError struct {
	StatusCode int
	Reason     metav1.StatusReason
	Message    string
}

func (e *AuthError) Error() string {
	return e.Message
}

type IngestAuthenticator interface {
	AuthenticateAndAuthorize(ctx context.Context, authHeader string, targetNamespace string) (*AuthResult, error)
}

// TokenReviewAuthenticator validates projected ServiceAccount tokens against kube-apiserver TokenReview
// and authorizes namespace-scoped runtime-state pushes.
type TokenReviewAuthenticator struct {
	Client         kubernetes.Interface
	Audience       string
	SiteNamespaces SiteNamespaceLister
}

func NewTokenReviewAuthenticator(client kubernetes.Interface, audience string, siteNamespaces SiteNamespaceLister) *TokenReviewAuthenticator {
	if audience == "" {
		audience = RuntimeTokenAudience
	}
	return &TokenReviewAuthenticator{
		Client:         client,
		Audience:       audience,
		SiteNamespaces: siteNamespaces,
	}
}

func (a *TokenReviewAuthenticator) AuthenticateAndAuthorize(ctx context.Context, authHeader string, targetNamespace string) (*AuthResult, error) {
	token, ok := extractBearerToken(authHeader)
	if !ok {
		return nil, &AuthError{
			StatusCode: http.StatusUnauthorized,
			Reason:     metav1.StatusReasonUnauthorized,
			Message:    "missing or malformed bearer token",
		}
	}
	if a.Client == nil {
		return nil, &AuthError{
			StatusCode: http.StatusServiceUnavailable,
			Reason:     metav1.StatusReasonServiceUnavailable,
			Message:    "token review client not configured",
		}
	}

	audience := a.Audience
	if audience == "" {
		audience = RuntimeTokenAudience
	}
	review := &authenticationv1.TokenReview{
		Spec: authenticationv1.TokenReviewSpec{
			Token:     token,
			Audiences: []string{audience},
		},
	}
	res, err := a.Client.AuthenticationV1().TokenReviews().Create(ctx, review, metav1.CreateOptions{})
	if err != nil {
		return nil, &AuthError{
			StatusCode: http.StatusServiceUnavailable,
			Reason:     metav1.StatusReasonServiceUnavailable,
			Message:    fmt.Sprintf("token review request failed: %v", err),
		}
	}
	if !res.Status.Authenticated || res.Status.Error != "" {
		msg := "token authentication failed"
		if res.Status.Error != "" {
			msg = fmt.Sprintf("token authentication failed: %s", res.Status.Error)
		}
		return nil, &AuthError{
			StatusCode: http.StatusUnauthorized,
			Reason:     metav1.StatusReasonUnauthorized,
			Message:    msg,
		}
	}

	tokenNamespace, ok := parseServiceAccountNamespace(res.Status.User.Username)
	if !ok {
		return nil, &AuthError{
			StatusCode: http.StatusForbidden,
			Reason:     metav1.StatusReasonForbidden,
			Message:    fmt.Sprintf("caller %q is not a Kubernetes service account", res.Status.User.Username),
		}
	}
	if tokenNamespace != targetNamespace {
		return nil, &AuthError{
			StatusCode: http.StatusForbidden,
			Reason:     metav1.StatusReasonForbidden,
			Message:    fmt.Sprintf("service account namespace %q is not authorized to push state for namespace %q", tokenNamespace, targetNamespace),
		}
	}
	if !a.isControlledSiteNamespace(targetNamespace) {
		return nil, &AuthError{
			StatusCode: http.StatusForbidden,
			Reason:     metav1.StatusReasonForbidden,
			Message:    fmt.Sprintf("namespace %q does not have a controlled Skupper site", targetNamespace),
		}
	}

	var podName string
	if vals := res.Status.User.Extra[podNameExtraKey]; len(vals) > 0 {
		podName = strings.TrimSpace(vals[0])
	}

	return &AuthResult{
		Username:  res.Status.User.Username,
		Namespace: tokenNamespace,
		PodName:   podName,
	}, nil
}

func (a *TokenReviewAuthenticator) isControlledSiteNamespace(namespace string) bool {
	if a.SiteNamespaces == nil {
		return true
	}
	for _, ns := range a.SiteNamespaces() {
		if ns == namespace {
			return true
		}
	}
	return false
}

func extractBearerToken(header string) (string, bool) {
	parts := strings.SplitN(strings.TrimSpace(header), " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", false
	}
	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", false
	}
	return token, true
}

func parseServiceAccountNamespace(username string) (string, bool) {
	parts := strings.Split(username, ":")
	if len(parts) != 4 || parts[0] != "system" || parts[1] != "serviceaccount" {
		return "", false
	}
	if parts[2] == "" || parts[3] == "" {
		return "", false
	}
	return parts[2], true
}
