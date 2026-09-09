package ddns

import "testing"

func TestProviderProfiles(t *testing.T) {
    expected := map[string]bool{"duckdns": false, "porkbun": false, "caddy_dns": false}
    for _, p := range ProviderProfiles {
        if _, ok := expected[p.Name]; !ok { continue }
        if len(p.CredentialEnv) == 0 || !p.SupportsIPv4 || !p.SupportsIPv6 { t.Fatalf("invalid profile: %+v", p) }
        expected[p.Name] = true
    }
    for name, found := range expected { if !found { t.Fatalf("missing provider profile %s", name) } }
}
