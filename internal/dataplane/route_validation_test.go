package dataplane

import "testing"

func TestValidateRoutes(t *testing.T) {
	good := RouteIntent{Prefix:"203.0.113.0/24", Family:IPv4, NextHop:"192.0.2.1", Authorized:true}
	if err := ValidateRoute(good); err != nil { t.Fatal(err) }
	bad := good; bad.Authorized = false
	if err := ValidateRoute(bad); err != ErrRouteUnauthorized { t.Fatalf("got %v", err) }
	bad = good; bad.Prefix = "not-a-prefix"
	if err := ValidateRoute(bad); err != ErrInvalidPrefix { t.Fatalf("got %v", err) }
	bad = good; bad.NextHop = "2001:db8::1"
	if err := ValidateRoute(bad); err != ErrRouteFamilyMismatch { t.Fatalf("got %v", err) }
}

func TestValidateRoutesRejectsCanonicalDuplicates(t *testing.T) {
	routes := []RouteIntent{
		{Prefix:"203.0.113.1/24", Family:IPv4, NextHop:"192.0.2.1", Authorized:true},
		{Prefix:"203.0.113.0/24", Family:IPv4, NextHop:"192.0.2.1", Authorized:true},
	}
	if err := ValidateRoutes(routes); err != ErrDuplicateRoute { t.Fatalf("got %v", err) }
}

func TestValidateRoutesAcceptsIPv6(t *testing.T) {
	r := RouteIntent{Prefix:"2001:db8:1::/48", Family:IPv6, NextHop:"2001:db8::1", Authorized:true}
	if err := ValidateRoute(r); err != nil { t.Fatal(err) }
}

func TestValidateRouteRejectsControlCharacters(t *testing.T) {
	r := RouteIntent{Prefix:"203.0.113.0/24\n", Family:IPv4, NextHop:"192.0.2.1", Authorized:true}
	if err := ValidateRoute(r); err != ErrRouteControlCharacter { t.Fatalf("got %v", err) }
}
