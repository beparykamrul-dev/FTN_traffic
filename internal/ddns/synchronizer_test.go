package ddns

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/beparykamrul-dev/FTN_traffic/internal/authorization"
	ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
)

type approvalStub struct{ err error }
func (a approvalStub) Verify(context.Context, string) error { return a.err }

type syncProvider struct{ name string; calls int; record ftndns.ZoneRecord }
func (p *syncProvider) Name() string { return p.name }
func (p *syncProvider) Health(context.Context) error { return nil }
func (p *syncProvider) ListRecords(context.Context, string) ([]ftndns.ZoneRecord, error) { return nil, nil }
func (p *syncProvider) UpsertRecord(_ context.Context, _ string, r ftndns.ZoneRecord) error { p.calls++; p.record = r; return nil }
func (p *syncProvider) DeleteRecord(context.Context, string, string, string) error { return nil }

func TestSynchronizerRequiresApprovalVerification(t *testing.T) {
	reg := ftndns.NewRegistry()
	p := &syncProvider{name: "route53_aws"}
	_ = reg.Register(p)
	auth := authorization.NewManager()
	auth.Upsert(authorization.ProviderAuth{Provider: "route53_aws", Status: authorization.Connected})
	s := Synchronizer{Registry: reg, Authorization: auth, Approval: approvalStub{err: errors.New("approval denied")}}
	got := s.Apply(context.Background(), Update{Zone: "example.com", Name: "home", Address: "203.0.113.10", TTL: 60, ApprovalID: "ap-1"}, []ProviderTarget{{Provider: "route53_aws", Enabled: true}})
	if len(got) != 1 || got[0].Applied || got[0].Error == "" || p.calls != 0 { t.Fatalf("unexpected result: %+v calls=%d", got, p.calls) }
}

func TestSynchronizerRoute53StyleAAndAAAA(t *testing.T) {
	reg := ftndns.NewRegistry()
	p := &syncProvider{name: "route53_aws"}
	if err := reg.Register(p); err != nil { t.Fatal(err) }
	auth := authorization.NewManager()
	auth.Upsert(authorization.ProviderAuth{Provider: "route53_aws", Status: authorization.Connected, ExpiresAt: time.Now().Add(time.Hour)})
	s := Synchronizer{Registry: reg, Authorization: auth, Approval: approvalStub{}, Now: time.Now}
	for _, tc := range []struct{ ip, typ string }{{"203.0.113.10", "A"}, {"2001:db8::10", "AAAA"}} {
		got := s.Apply(context.Background(), Update{Zone: "example.com", Name: "home", Address: tc.ip, TTL: 60, ApprovalID: "ap-1"}, []ProviderTarget{{Provider: "route53_aws", Enabled: true}})
		if len(got) != 1 || !got[0].Applied || p.record.Type != tc.typ { t.Fatalf("unexpected result: %+v record=%+v", got, p.record) }
	}
}
