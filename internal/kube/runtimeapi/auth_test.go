package runtimeapi

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
	authenticationv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	k8stesting "k8s.io/client-go/testing"
)

func newFakeTokenReviewClient(handler func(*authenticationv1.TokenReview) (*authenticationv1.TokenReview, error)) *fake.Clientset {
	client := fake.NewSimpleClientset()
	client.PrependReactor("create", "tokenreviews", func(action k8stesting.Action) (bool, runtime.Object, error) {
		createAction := action.(k8stesting.CreateAction)
		tr := createAction.GetObject().(*authenticationv1.TokenReview)
		res, err := handler(tr)
		return true, res, err
	})
	return client
}

func TestTokenReviewAuthenticator(t *testing.T) {
	controlled := func() []string { return []string{"west", "east"} }

	t.Run("valid token and matching namespace succeeds", func(t *testing.T) {
		var seenAudience string
		var seenToken string
		client := newFakeTokenReviewClient(func(tr *authenticationv1.TokenReview) (*authenticationv1.TokenReview, error) {
			seenToken = tr.Spec.Token
			if len(tr.Spec.Audiences) > 0 {
				seenAudience = tr.Spec.Audiences[0]
			}
			tr.Status = authenticationv1.TokenReviewStatus{
				Authenticated: true,
				User: authenticationv1.UserInfo{
					Username: "system:serviceaccount:west:skupper-router",
					Extra: map[string]authenticationv1.ExtraValue{
						podNameExtraKey: {"skupper-router-abc123"},
					},
				},
			}
			return tr, nil
		})

		auth := NewTokenReviewAuthenticator(client, RuntimeTokenAudience, controlled)
		res, err := auth.AuthenticateAndAuthorize(context.Background(), "Bearer secret-jwt", "west")
		assert.NilError(t, err)
		assert.Equal(t, seenToken, "secret-jwt")
		assert.Equal(t, seenAudience, RuntimeTokenAudience)
		assert.Equal(t, res.Namespace, "west")
		assert.Equal(t, res.PodName, "skupper-router-abc123")
		assert.Equal(t, res.Username, "system:serviceaccount:west:skupper-router")
	})

	t.Run("missing or malformed authorization header returns 401", func(t *testing.T) {
		auth := NewTokenReviewAuthenticator(fake.NewSimpleClientset(), RuntimeTokenAudience, controlled)
		for _, hdr := range []string{"", "Basic abc", "Bearer ", "Bearer"} {
			_, err := auth.AuthenticateAndAuthorize(context.Background(), hdr, "west")
			var authErr *AuthError
			assert.Assert(t, errors.As(err, &authErr))
			assert.Equal(t, authErr.StatusCode, http.StatusUnauthorized)
			assert.Equal(t, authErr.Reason, metav1.StatusReasonUnauthorized)
		}
	})

	t.Run("unauthenticated token returns 401", func(t *testing.T) {
		client := newFakeTokenReviewClient(func(tr *authenticationv1.TokenReview) (*authenticationv1.TokenReview, error) {
			tr.Status = authenticationv1.TokenReviewStatus{
				Authenticated: false,
				Error:         "token expired",
			}
			return tr, nil
		})
		auth := NewTokenReviewAuthenticator(client, RuntimeTokenAudience, controlled)
		_, err := auth.AuthenticateAndAuthorize(context.Background(), "Bearer expired-token", "west")
		var authErr *AuthError
		assert.Assert(t, errors.As(err, &authErr))
		assert.Equal(t, authErr.StatusCode, http.StatusUnauthorized)
	})

	t.Run("mismatched namespace returns 403", func(t *testing.T) {
		client := newFakeTokenReviewClient(func(tr *authenticationv1.TokenReview) (*authenticationv1.TokenReview, error) {
			tr.Status = authenticationv1.TokenReviewStatus{
				Authenticated: true,
				User: authenticationv1.UserInfo{
					Username: "system:serviceaccount:east:skupper-router",
				},
			}
			return tr, nil
		})
		auth := NewTokenReviewAuthenticator(client, RuntimeTokenAudience, controlled)
		_, err := auth.AuthenticateAndAuthorize(context.Background(), "Bearer east-token", "west")
		var authErr *AuthError
		assert.Assert(t, errors.As(err, &authErr))
		assert.Equal(t, authErr.StatusCode, http.StatusForbidden)
		assert.Equal(t, authErr.Reason, metav1.StatusReasonForbidden)
	})

	t.Run("uncontrolled namespace returns 403", func(t *testing.T) {
		client := newFakeTokenReviewClient(func(tr *authenticationv1.TokenReview) (*authenticationv1.TokenReview, error) {
			tr.Status = authenticationv1.TokenReviewStatus{
				Authenticated: true,
				User: authenticationv1.UserInfo{
					Username: "system:serviceaccount:rogue:skupper-router",
				},
			}
			return tr, nil
		})
		auth := NewTokenReviewAuthenticator(client, RuntimeTokenAudience, controlled)
		_, err := auth.AuthenticateAndAuthorize(context.Background(), "Bearer rogue-token", "rogue")
		var authErr *AuthError
		assert.Assert(t, errors.As(err, &authErr))
		assert.Equal(t, authErr.StatusCode, http.StatusForbidden)
	})

	t.Run("non-serviceaccount principal returns 403", func(t *testing.T) {
		client := newFakeTokenReviewClient(func(tr *authenticationv1.TokenReview) (*authenticationv1.TokenReview, error) {
			tr.Status = authenticationv1.TokenReviewStatus{
				Authenticated: true,
				User: authenticationv1.UserInfo{
					Username: "alice@example.com",
				},
			}
			return tr, nil
		})
		auth := NewTokenReviewAuthenticator(client, RuntimeTokenAudience, controlled)
		_, err := auth.AuthenticateAndAuthorize(context.Background(), "Bearer user-token", "west")
		var authErr *AuthError
		assert.Assert(t, errors.As(err, &authErr))
		assert.Equal(t, authErr.StatusCode, http.StatusForbidden)
	})
}
