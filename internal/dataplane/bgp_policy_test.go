package dataplane

import "testing"

func TestRoutePolicy(t *testing.T) {
	good := RouteIntent{Prefix:"203.0.113.0/24", Family:IPv4, Authorized:true}
	p := RoutePolicy{RequireRPKI:true, MaxPrefixes:2}
	if err := p.Validate(RouteBatch{Routes:[]RouteIntent{good}, RPKIValid:true}); err != nil { t.Fatal(err) }
	if err := p.Validate(RouteBatch{Routes:[]RouteIntent{good}, RPKIValid:false}); err != ErrRPKIRequired { t.Fatalf("got %v", err) }
}
