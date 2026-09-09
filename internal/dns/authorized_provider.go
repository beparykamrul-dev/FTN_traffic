package dns

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/beparykamrul-dev/FTN_traffic/internal/authorization"
)

type AuthorizedProvider struct {
	Inner   Provider
	Auth    *authorization.Manager
	Now     func() time.Time
}

func (p AuthorizedProvider) Name() string {
	if p.Inner == nil { return "" }
	return p.Inner.Name()
}

func (p AuthorizedProvider) Health(ctx context.Context) error {
	if err := p.authorized(); err != nil { return err }
	return p.Inner.Health(ctx)
}

func (p AuthorizedProvider) ListRecords(ctx context.Context, zone string) ([]ZoneRecord, error) {
	if err := p.authorized(); err != nil { return nil, err }
	return p.Inner.ListRecords(ctx, zone)
}

func (p AuthorizedProvider) UpsertRecord(ctx context.Context, zone string, record ZoneRecord) error {
	if err := p.authorized(); err != nil { return err }
	return p.Inner.UpsertRecord(ctx, zone, record)
}

func (p AuthorizedProvider) DeleteRecord(ctx context.Context, zone, name, recordType string) error {
	if err := p.authorized(); err != nil { return err }
	return p.Inner.DeleteRecord(ctx, zone, name, recordType)
}

func (p AuthorizedProvider) authorized() error {
	if p.Inner == nil { return fmt.Errorf("DNS provider is nil") }
	if p.Auth == nil { return fmt.Errorf("DNS provider authorization manager is nil") }
	now := time.Now()
	if p.Now != nil { now = p.Now() }
	provider := strings.TrimSpace(p.Inner.Name())
	if provider == "" { return fmt.Errorf("DNS provider name is empty") }
	return p.Auth.Authorized(provider, now)
}
