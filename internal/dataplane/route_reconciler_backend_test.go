package dataplane

import (
	"context"
	"errors"
	"testing"
)

type routeBackendStub struct { healthErr error; applied int; withdrawn int }
func (r *routeBackendStub) Name() string { return "stub" }
func (r *routeBackendStub) Health(context.Context) error { return r.healthErr }
func (r *routeBackendStub) ApplyRoutes(context.Context, []RouteIntent) error { r.applied++; return nil }
func (r *routeBackendStub) WithdrawRoutes(context.Context, []RouteIntent) error { r.withdrawn++; return nil }
func (r *routeBackendStub) Snapshot(context.Context) (RouterStatus,error) { return RouterStatus{},nil }

func TestRouteReconcilerBackendHealth(t *testing.T) {
	b := &routeBackendStub{healthErr: errors.New("router down")}
	r := RouteReconciler{Router:b, Authorized:true, Approved:true, Policy:&RoutePolicy{MaxPrefixes:10}}
	routes := []RouteIntent{{Prefix:"203.0.113.0/24",Family:IPv4,NextHop:"192.0.2.1",Authorized:true}}
	if err := r.Apply(context.Background(), routes); err == nil { t.Fatal("expected backend health failure") }
	if b.applied != 0 { t.Fatal("route mutation must not execute on unhealthy backend") }
}
