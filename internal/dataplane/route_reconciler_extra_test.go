package dataplane

import (
	"context"
	"testing"
)

type reconRouter struct{ err error }
func (r reconRouter) Name() string { return "test" }
func (r reconRouter) Health(context.Context) error { return r.err }
func (r reconRouter) ApplyRoutes(context.Context, []RouteIntent) error { return nil }
func (r reconRouter) WithdrawRoutes(context.Context, []RouteIntent) error { return nil }
func (r reconRouter) Snapshot(context.Context) (RouterStatus, error) { return RouterStatus{}, nil }

func validReconRoute() []RouteIntent {
	return []RouteIntent{{Prefix: "203.0.113.0/24", Family: IPv4, Authorized: true}}
}

func TestRouteReconcilerRequiresApproval(t *testing.T) {
	r := RouteReconciler{Router: reconRouter{}, Authorized: true}
	if err := r.Apply(context.Background(), validReconRoute()); err != ErrApprovalRequired { t.Fatalf("got %v", err) }
}

func TestRouteReconcilerGatesOnBFD(t *testing.T) {
	bfd := &BFDSession{ID:"b1", Local:"192.0.2.1", Remote:"192.0.2.2", MinRxMS:50, MinTxMS:50, Multiplier:3, Authorized:true, Up:false}
	r := RouteReconciler{Router: reconRouter{}, Authorized:true, Approved:true, BFD:bfd}
	if err := r.Apply(context.Background(), validReconRoute()); err != ErrBFDInvalid { t.Fatalf("got %v", err) }
	bfd.Up = true
	if err := r.Apply(context.Background(), validReconRoute()); err != nil { t.Fatalf("expected BFD-gated apply to pass: %v", err) }
}

func TestRouteReconcilerPropagatesRouterHealthFailure(t *testing.T) {
	r := RouteReconciler{Router: reconRouter{err: context.DeadlineExceeded}, Authorized:true, Approved:true}
	if err := r.Apply(context.Background(), validReconRoute()); err != context.DeadlineExceeded { t.Fatalf("got %v", err) }
}
