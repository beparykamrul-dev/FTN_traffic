package ddns

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/beparykamrul-dev/FTN_traffic/internal/authorization"
	ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
)

type ApprovalVerifier interface { Verify(context.Context, string) error }

type Synchronizer struct {
	Registry *ftndns.Registry
	Authorization *authorization.Manager
	Approval ApprovalVerifier
	Now func() time.Time
}

type SyncResult struct {
	Provider string `json:"provider"`
	Zone string `json:"zone"`
	Record string `json:"record"`
	Applied bool `json:"applied"`
	Error string `json:"error,omitempty"`
}

func (s *Synchronizer) Apply(ctx context.Context, u Update, targets []ProviderTarget) []SyncResult {
	results := make([]SyncResult, 0, len(targets))
	if err := Validate(u); err != nil { return []SyncResult{{Error: err.Error()}} }
	if s.Registry == nil || s.Authorization == nil || s.Approval == nil { return []SyncResult{{Error: "DDNS synchronization dependencies are not configured"}} }
	if err := s.Approval.Verify(ctx, u.ApprovalID); err != nil { return []SyncResult{{Error: err.Error()}} }
	now := time.Now(); if s.Now != nil { now = s.Now() }
	for _, target := range targets {
		if !target.Enabled { continue }
		providerName := strings.TrimSpace(target.Provider)
		zone := strings.TrimSpace(target.Zone); if zone == "" { zone = u.Zone }
		name := strings.TrimSpace(target.RecordName); if name == "" { name = u.Name }
		result := SyncResult{Provider: providerName, Zone: zone, Record: name}
		if providerName == "" || zone == "" || name == "" { result.Error = "invalid provider target"; results = append(results, result); continue }
		if err := s.Authorization.Authorized(providerName, now); err != nil { result.Error = err.Error(); results = append(results, result); continue }
		p, ok := s.Registry.Get(providerName); if !ok || p == nil { result.Error = "provider is not registered"; results = append(results, result); continue }
		if err := p.UpsertRecord(ctx, zone, ftndns.ZoneRecord{Name: name, Type: recordType(u.Address), TTL: u.TTL, Value: u.Address}); err != nil { result.Error = err.Error() } else { result.Applied = true }
		results = append(results, result)
	}
	return results
}

func recordType(address string) string {
	ip := net.ParseIP(strings.TrimSpace(address)); if ip != nil && ip.To4() == nil { return "AAAA" }; return "A"
}

var _ = fmt.Sprintf
