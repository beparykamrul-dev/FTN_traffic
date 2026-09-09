package ddns

import (
    "context"
    "fmt"
    "net"

    ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
)

// ProviderAdapter maps a DDNS update to the existing authorized DNS provider.
// Provider credentials and remote API access remain inside the provider adapter.
type ProviderAdapter struct { Provider ftndns.Provider }

func (a ProviderAdapter) Name() string {
    if a.Provider == nil { return "" }
    return a.Provider.Name()
}

func (a ProviderAdapter) Apply(ctx context.Context, u Update) error {
    if a.Provider == nil { return fmt.Errorf("DDNS provider is nil") }
    if err := Validate(u); err != nil { return err }
    recordType := "A"
    if ip := net.ParseIP(u.Address); ip != nil && ip.To4() == nil { recordType = "AAAA" }
    return a.Provider.UpsertRecord(ctx, u.Zone, ftndns.ZoneRecord{
        Name: u.Name, Type: recordType, TTL: u.TTL, Value: u.Address,
    })
}
