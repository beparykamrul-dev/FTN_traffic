package ddns

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net"
    "net/http"
    "net/url"
    "strings"
    "time"
)

type HTTPDoer interface { Do(*http.Request) (*http.Response, error) }

type DuckDNSClient struct { HTTP HTTPDoer; Token string; BaseURL string }

func (c DuckDNSClient) Apply(ctx context.Context, u Update) error {
    if strings.TrimSpace(c.Token) == "" { return fmt.Errorf("duckdns token is not configured") }
    if err := Validate(u); err != nil { return err }
    base := c.BaseURL; if base == "" { base = "https://www.duckdns.org/update" }
    q := url.Values{}
    q.Set("domains", strings.TrimSuffix(strings.TrimSpace(u.Name), "."))
    q.Set("token", c.Token)
    q.Set("verbose", "true")
    if ip := net.ParseIP(strings.TrimSpace(u.Address)); ip != nil && ip.To4() == nil { q.Set("ipv6", u.Address) } else { q.Set("ip", u.Address) }
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, base+"?"+q.Encode(), nil); if err != nil { return err }
    client := c.HTTP; if client == nil { client = http.DefaultClient }
    resp, err := client.Do(req); if err != nil { return err }; defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 { return fmt.Errorf("duckdns update returned HTTP %d", resp.StatusCode) }
    return nil
}

type PorkbunClient struct { HTTP HTTPDoer; APIKey string; SecretAPIKey string; BaseURL string }

func (c PorkbunClient) Apply(ctx context.Context, u Update) error {
    if strings.TrimSpace(c.APIKey) == "" || strings.TrimSpace(c.SecretAPIKey) == "" { return fmt.Errorf("porkbun credentials are not configured") }
    if err := Validate(u); err != nil { return err }
    base := c.BaseURL; if base == "" { base = "https://api.porkbun.com/api/json/v3" }
    body, err := json.Marshal(map[string]any{"secretapikey": c.SecretAPIKey, "apikey": c.APIKey, "name": strings.TrimSuffix(u.Name, "."), "type": recordType(u.Address), "content": u.Address, "ttl": u.TTL})
    if err != nil { return err }
    endpoint := strings.TrimRight(base, "/") + "/dns/create/" + strings.TrimSuffix(strings.TrimSpace(u.Zone), ".")
    req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body)); if err != nil { return err }
    req.Header.Set("Content-Type", "application/json")
    client := c.HTTP; if client == nil { client = http.DefaultClient }
    resp, err := client.Do(req); if err != nil { return err }; defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 { return fmt.Errorf("porkbun update returned HTTP %d", resp.StatusCode) }
    return nil
}

var providerHTTPTimeout = 15 * time.Second
func ProviderHTTPClient() *http.Client { return &http.Client{Timeout: providerHTTPTimeout} }
