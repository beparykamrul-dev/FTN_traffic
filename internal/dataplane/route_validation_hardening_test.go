package dataplane

import "testing"

func TestValidateRouteFamilyAndNextHop(t *testing.T) {
	base := RouteIntent{Prefix: "203.0.113.0/24", Family: IPv4, Authorized: true}
	if err := ValidateRoute(base); err != nil { t.Fatal(err) }
	base.Family = IPv6
	if err := ValidateRoute(base); err != ErrRouteFamilyMismatch { t.Fatalf("expected family mismatch, got %v", err) }
	base.Family = IPv4
	base.NextHop = "2001:db8::1"
	if err := ValidateRoute(base); err != ErrRouteFamilyMismatch { t.Fatalf("expected next-hop family mismatch, got %v", err) }
}

func TestValidateRoutesRejectsDuplicate(t *testing.T) {
	r := RouteIntent{Prefix: "203.0.113.0/24", Family: IPv4, Authorized: true}
	if err := ValidateRoutes([]RouteIntent{r, r}); err == nil { t.Fatal("expected duplicate route rejection") }
}
