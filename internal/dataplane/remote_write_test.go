package dataplane

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteWriteBackendPublish(t *testing.T) {
	var gotAuth string
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	b := RemoteWriteBackend{Endpoint: "https://example.invalid/write", Token: "secret", Client: srv.Client()}
	// httptest uses a self-signed certificate, so the endpoint validator is
	// exercised separately; the transport test uses a custom client URL.
	b.Endpoint = srv.URL
	if err := b.Publish(context.Background(), Metric{Name: "ftn_test_total", Value: 1}); err != nil {
		t.Fatal(err)
	}
	if gotAuth != "Bearer secret" { t.Fatalf("unexpected authorization header: %q", gotAuth) }
}

func TestRemoteWriteBackendRejectsInsecureEndpoint(t *testing.T) {
	b := RemoteWriteBackend{Endpoint: "http://metrics.example/write", Token: "secret"}
	if _, err := b.validate(); !errors.Is(err, ErrRemoteWriteEndpoint) { t.Fatalf("expected endpoint error, got %v", err) }
}

func TestRemoteWriteBackendRejectsCredentialsInURL(t *testing.T) {
	b := RemoteWriteBackend{Endpoint: "https://user:pass@example.com/write", Token: "secret"}
	if _, err := b.validate(); !errors.Is(err, ErrRemoteWriteEndpoint) { t.Fatalf("expected endpoint error, got %v", err) }
}

func TestRemoteWriteBackendRejectsNon2xx(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusUnauthorized) }))
	defer srv.Close()
	b := RemoteWriteBackend{Endpoint: srv.URL, Token: "secret", Client: srv.Client()}
	if err := b.Publish(context.Background(), Metric{Name: "ftn_test_total", Value: 1}); !errors.Is(err, ErrRemoteWriteResponse) { t.Fatalf("expected response error, got %v", err) }
}
