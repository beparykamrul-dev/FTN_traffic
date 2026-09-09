package authorization

import (
	"testing"
	"time"
)

func TestAuthorizationLifecycle(t *testing.T) {
	m := NewManager()
	now := time.Now()
	m.Upsert(ProviderAuth{Provider: "cloudflare", Status: Connected, Method: "api", ExpiresAt: now.Add(time.Hour)})
	if err := m.Authorized("cloudflare", now); err != nil { t.Fatal(err) }
	if err := m.Authorized("akamai", now); err != ErrNotConnected { t.Fatalf("expected not connected, got %v", err) }
}

func TestExpiredAuthorization(t *testing.T) {
	m := NewManager()
	now := time.Now()
	m.Upsert(ProviderAuth{Provider: "fastly", Status: Connected, ExpiresAt: now.Add(-time.Minute)})
	if err := m.Authorized("fastly", now); err != ErrNotConnected { t.Fatalf("expected expired authorization rejection, got %v", err) }
}
