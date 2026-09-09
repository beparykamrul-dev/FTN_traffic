package providers

import "testing"

func TestRegistry(t *testing.T) {
	r := NewRegistry(Provider{Name: "github", Role: "origin"})
	if _, ok := r.Get("missing"); ok { t.Fatal("missing provider unexpectedly found") }
	r.Set(Provider{Name: "cdn", Role: "edge_delivery"})
	p, ok := r.Get("cdn")
	if !ok || p.Role != "edge_delivery" { t.Fatalf("unexpected provider: %+v", p) }
}
