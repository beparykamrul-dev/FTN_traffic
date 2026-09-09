package dataplane
import "testing"
func TestValidateRoutesRejectsCanonicalDuplicate(t *testing.T){rs:=[]RouteIntent{{Prefix:"203.0.113.1/24",Family:IPv4,Authorized:true},{Prefix:"203.0.113.0/24",Family:IPv4,Authorized:true}};if err:=ValidateRoutes(rs);err!=ErrDuplicateRoute{t.Fatalf("got %v",err)}}
func TestValidateRouteRejectsFamilyMismatch(t *testing.T){r:=RouteIntent{Prefix:"2001:db8::/32",Family:IPv4,Authorized:true};if err:=ValidateRoute(r);err!=ErrRouteFamilyMismatch{t.Fatalf("got %v",err)}}
