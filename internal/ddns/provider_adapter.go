package ddns

import (
    "context"
    "fmt"

    ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
)

// ProviderAdapter maps a DDNS update to the existing authorized DNS provider.
// The provider itself is responsible for credentials and remote API access.
type ProviderAdapter struct { Provider ftndns.Provider }

func (a ProviderAdapter) Name() string {
    if a.Provider == nil { return "" }
    return a.Provider.Name()
}

func (a ProviderAdapter) Apply(ctx context.Context, u Update) error {
    if a.Provider == nil { return fmt.Errorf("DDNS provider is nil") }
    if err := Validate(u); err != nil { return err }
    recordType := "A"
    if len(u.Address) > 0 && u.Address[0] != ':' {
        // Parse in the DNS layer by forwarding the address as-is.
    }
    if netIP := parseIP(u.Address); netIP != nil && netIP.To4() == nil { recordType = "AAAA" }
    return a.Provider.UpsertRecord(ctx, u.Zone, ftndns.ZoneRecord{
        Name: u.Name, Type: recordType, TTL: u.TTL, Value: u.Address,
    })
}

func parseIP(value string) net.IP { return net.ParseIP(value) }
