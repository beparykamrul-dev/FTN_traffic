package dataplane

import "testing"

func TestValidateRouteCommunities(t *testing.T) {
	r := RouteIntent{Prefix:"203.0.113.0/24", Family:IPv4, Community:[]string{"64512:100"}, Authorized:true}
	if err := ValidateRouteCommunities(r); err != nil { t.Fatal(err) }
	r.Community = []string{"bad"}
	if err := ValidateRouteCommunities(r); err != ErrInvalidCommunity { t.Fatalf("got %v", err) }
}

func TestValidateRouteWithRPKI(t *testing.T) {
	r := RouteIntent{Prefix:"203.0.113.0/24", Family:IPv4, Authorized:true}
	if err := ValidateRouteWithRPKI(r, false); err != ErrRPKIRequired { t.Fatalf("got %v", err) }
	if err := ValidateRouteWithRPKI(r, true); err != nil { t.Fatal(err) }
}
