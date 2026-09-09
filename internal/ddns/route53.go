package ddns

import (
    "context"
    "errors"
    "net"
    "strings"
)

var ErrRoute53NotConfigured = errors.New("route53 provider is not configured")

type Route53Change struct {
    HostedZoneID string
    RecordName string
    RecordType string
    TTL uint32
    Address string
}

type Route53Client interface {
    ChangeRecord(ctx context.Context, change Route53Change) error
}

type Route53Provider struct {
    Client Route53Client
    HostedZoneID string
}

func (p Route53Provider) Name() string { return "route53_aws" }

func (p Route53Provider) Apply(ctx context.Context, u Update) error {
    if p.Client == nil || strings.TrimSpace(p.HostedZoneID) == "" { return ErrRoute53NotConfigured }
    if err := Validate(u); err != nil { return err }
    ip := net.ParseIP(strings.TrimSpace(u.Address))
    typ := "A"
    if ip.To4() == nil { typ = "AAAA" }
    return p.Client.ChangeRecord(ctx, Route53Change{HostedZoneID: p.HostedZoneID, RecordName: u.Name, RecordType: typ, TTL: u.TTL, Address: u.Address})
}
