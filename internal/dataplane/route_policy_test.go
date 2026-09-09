package dataplane

import "testing"

func TestRoutePolicyRequiresRPKI(t *testing.T) {
 p:=RoutePolicy{RequireRPKI:true,MaxPrefixes:10}
 r:=RouteBatch{RPKIValid:false,Routes:[]RouteIntent{{Prefix:"203.0.113.0/24",Family:IPv4,Authorized:true}}}
 if err:=p.Validate(r);err!=ErrRPKIRequired{t.Fatalf("expected RPKI error, got %v",err)}
}

func TestRoutePolicyRejectsZeroAndExcessPrefixes(t *testing.T) {
 r:=RouteBatch{RPKIValid:true,Routes:[]RouteIntent{{Prefix:"203.0.113.0/24",Family:IPv4,Authorized:true}}}
 if err:=(RoutePolicy{}).Validate(r);err!=ErrMaxPrefixExceeded{t.Fatalf("expected zero-limit error, got %v",err)}
 if err:=(RoutePolicy{MaxPrefixes:1}).Validate(RouteBatch{RPKIValid:true,Routes:append(r.Routes,r.Routes[0])});err!=ErrMaxPrefixExceeded{t.Fatalf("expected max-prefix error, got %v",err)}
}

func TestRoutePolicyRejectsEmptyBatch(t *testing.T) {
 if err:=(RoutePolicy{MaxPrefixes:10}).Validate(RouteBatch{RPKIValid:true});err!=ErrEmptyRouteBatch{t.Fatalf("expected empty-batch error, got %v",err)}
}
