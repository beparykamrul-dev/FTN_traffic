package ddns

import (
    "context"
    "testing"
)

type fakeRoute53Client struct { change Route53Change }
func (f *fakeRoute53Client) ChangeRecord(_ context.Context, c Route53Change) error { f.change = c; return nil }

func TestRoute53ProviderApply(t *testing.T) {
    c := &fakeRoute53Client{}
    p := Route53Provider{Client: c, HostedZoneID: "ZTEST"}
    err := p.Apply(context.Background(), Update{Zone: "example.com", Name: "edge.example.com", Address: "2001:db8::10", TTL: 60, ApprovalID: "approval-1"})
    if err != nil { t.Fatal(err) }
    if c.change.RecordType != "AAAA" || c.change.HostedZoneID != "ZTEST" { t.Fatalf("unexpected change: %+v", c.change) }
}

func TestRoute53ProviderRequiresConfiguration(t *testing.T) {
    p := Route53Provider{}
    if err := p.Apply(context.Background(), Update{Zone: "example.com", Name: "edge.example.com", Address: "192.0.2.10", TTL: 60, ApprovalID: "approval-1"}); err != ErrRoute53NotConfigured { t.Fatalf("got %v", err) }
}
