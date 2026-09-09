package dataplane

import "testing"

func TestNormalizeRouteIntent(t *testing.T) {
	r, err := NormalizeRouteIntent(RouteIntent{Prefix:"192.0.2.7/24", Family:IPv4, NextHop:"192.0.2.2", Authorized:true})
	if err != nil { t.Fatal(err) }
	if r.Prefix != "192.0.2.0/24" || r.NextHop != "192.0.2.2" { t.Fatalf("unexpected normalization: %+v", r) }
}

func TestNormalizeRoutesRejectsDuplicate(t *testing.T) {
	r := RouteIntent{Prefix:"2001:db8:1::/64", Family:IPv6, Authorized:true}
	if _, err := NormalizeRoutes([]RouteIntent{r,r}); err != ErrDuplicateRoute { t.Fatalf("expected duplicate error, got %v", err) }
}
