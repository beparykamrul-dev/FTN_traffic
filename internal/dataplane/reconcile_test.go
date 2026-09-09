package dataplane

import (
	"context"
	"testing"
)

type testRouter struct{ healthy bool; applied int }
func (r *testRouter) Name() string { return "test" }
func (r *testRouter) Kind() RouterKind { return RouterFRR }
func (r *testRouter) Health(context.Context) error { if !r.healthy { return ErrBackendUnavailable }; return nil }
func (r *testRouter) ApplyRoutes(_ context.Context, x []RouteIntent) error { r.applied=len(x); return nil }
func (r *testRouter) WithdrawRoutes(context.Context, []RouteIntent) error { return nil }
func (r *testRouter) Snapshot(context.Context) (RouterStatus,error) { return RouterStatus{ID:"test",Kind:RouterFRR,Healthy:r.healthy},nil }

func TestRouteReconciler(t *testing.T) {
	r := &testRouter{healthy:true}
	c := RouteReconciler{Router:r}
	if err := c.Apply(context.Background(), []RouteIntent{{Prefix:"203.0.113.0/24",Family:IPv4,Authorized:true}}); err != nil { t.Fatal(err) }
	if r.applied != 1 { t.Fatalf("applied=%d",r.applied) }
}
