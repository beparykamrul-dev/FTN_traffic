package dns

import (
	"context"
	"testing"
	"time"

	"github.com/beparykamrul-dev/FTN_traffic/internal/authorization"
)

type testProvider struct { calls int }
func (p *testProvider) Name() string { return "cloudflare_dns" }
func (p *testProvider) Health(context.Context) error { p.calls++; return nil }
func (p *testProvider) ListRecords(context.Context, string) ([]ZoneRecord, error) { p.calls++; return nil, nil }
func (p *testProvider) UpsertRecord(context.Context, string, ZoneRecord) error { p.calls++; return nil }
func (p *testProvider) DeleteRecord(context.Context, string, string, string) error { p.calls++; return nil }

func TestAuthorizedProviderRejectsMissingAuthorization(t *testing.T) {
	inner := &testProvider{}
	wrapped := AuthorizedProvider{Inner: inner, Auth: authorization.NewManager()}
	if err := wrapped.Health(context.Background()); err == nil { t.Fatal("expected authorization error") }
	if inner.calls != 0 { t.Fatal("provider must not be called without authorization") }
}

func TestAuthorizedProviderAllowsActiveAuthorization(t *testing.T) {
	inner := &testProvider{}
	mgr := authorization.NewManager()
	mgr.Upsert(authorization.ProviderAuth{Provider: "cloudflare_dns", Status: authorization.Connected, ExpiresAt: time.Now().Add(time.Hour)})
	wrapped := AuthorizedProvider{Inner: inner, Auth: mgr}
	if err := wrapped.UpsertRecord(context.Background(), "example.com", ZoneRecord{Name: "www", Type: "A", TTL: 60, Value: "192.0.2.1"}); err != nil { t.Fatal(err) }
	if inner.calls != 1 { t.Fatalf("expected one provider call, got %d", inner.calls) }
}
