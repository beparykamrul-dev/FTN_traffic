package dataplane

import "testing"

func TestValidateRoutes(t *testing.T) {
	good := RouteIntent{Prefix:"203.0.113.0/24", Family:IPv4, NextHop:"192.0.2.1", Authorized:true}
	if err := ValidateRoute(good); err != nil { t.Fatal(err) }
	bad := good; bad.Authorized = false
	if err := ValidateRoute(bad); err != ErrRouteUnauthorized { t.Fatalf("got %v", err) }
	bad = good; bad.Prefix = "not-a-prefix"
	if err := ValidateRoute(bad); err != ErrInvalidPrefix { t.Fatalf("got %v", err) }
}
