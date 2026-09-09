package dataplane

import "testing"

func TestValidateRouteSetForMutation(t *testing.T) {
	if err := ValidateRouteSetForMutation([]RouteIntent{{Prefix:"192.0.2.0/24",Family:IPv4,Authorized:false}}); err != ErrRouteUnauthorized { t.Fatalf("expected authorization failure, got %v", err) }
	if err := ValidateRouteSetForMutation([]RouteIntent{{Prefix:"192.0.2.0/24",Family:IPv4,Authorized:true,LocalPref:100}}); err != nil { t.Fatal(err) }
}
