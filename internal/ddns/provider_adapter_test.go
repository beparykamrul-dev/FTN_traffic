package ddns

import (
    "context"
    "testing"

    ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
)

type testProvider struct { calls int; record ftndns.ZoneRecord }
func (p *testProvider) Name() string { return "test" }
func (p *testProvider) Health(context.Context) error { return nil }
func (p *testProvider) ListRecords(context.Context, string) ([]ftndns.ZoneRecord, error) { return nil, nil }
func (p *testProvider) UpsertRecord(_ context.Context, _ string, r ftndns.ZoneRecord) error { p.calls++; p.record = r; return nil }
func (p *testProvider) DeleteRecord(context.Context, string, string, string) error { return nil }

func TestProviderAdapterApplyIPv4(t *testing.T) {
    p := &testProvider{}
    err := (ProviderAdapter{Provider: p}).Apply(context.Background(), Update{Zone:"example.com", Name:"home", Address:"203.0.113.10", TTL:60, ApprovalID:"ap-1"})
    if err != nil || p.calls != 1 || p.record.Type != "A" { t.Fatalf("unexpected result: err=%v calls=%d record=%+v", err, p.calls, p.record) }
}

func TestProviderAdapterApplyIPv6(t *testing.T) {
    p := &testProvider{}
    err := (ProviderAdapter{Provider: p}).Apply(context.Background(), Update{Zone:"example.com", Name:"home", Address:"2001:db8::10", TTL:60, ApprovalID:"ap-1"})
    if err != nil || p.calls != 1 || p.record.Type != "AAAA" { t.Fatalf("unexpected result: err=%v calls=%d record=%+v", err, p.calls, p.record) }
}
