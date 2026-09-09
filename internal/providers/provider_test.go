package providers

import "testing"

func TestSupports(t *testing.T) { p := Provider{Capabilities: []string{"cache", "health"}}; if !Supports(p, "cache") || Supports(p, "bgp") { t.Fatal("capability matching failed") } }
