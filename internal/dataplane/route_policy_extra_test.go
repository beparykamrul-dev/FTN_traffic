package dataplane

import "testing"

func TestRoutePolicyRejectsUnauthorizedRoute(t *testing.T) {
	p := RoutePolicy{RequireRPKI:true, MaxPrefixes:10}
	err := p.Validate(RouteBatch{RPKIValid:true, Routes:[]RouteIntent{{Prefix:"203.0.113.0/24", Family:IPv4, Authorized:false}}})
	if err != ErrRouteUnauthorized { t.Fatalf("got %v", err) }
}
