package providers

import "testing"

func TestGlobalCatalog(t *testing.T) {
	seen := map[string]bool{}
	for _, p := range GlobalCatalog {
		if p.Name == "" || seen[p.Name] { t.Fatalf("invalid or duplicate provider: %+v", p) }
		seen[p.Name] = true
		if p.Mode == "" || len(p.Capabilities) == 0 { t.Fatalf("incomplete catalog entry: %+v", p) }
	}
	for _, want := range []string{"cloudflare", "akamai", "aws", "fastly", "bunny", "tencent_cloud", "tiktok", "netflix", "google", "facebook"} {
		if !seen[want] { t.Fatalf("missing %s", want) }
	}
}
