package dataplane

import (
	"context"
	"errors"
	"testing"
)

type reconcileRouter struct { healthErr error; applied bool; withdrawn bool }
func (r *reconcileRouter) Name() string { return "test" }
func (r *reconcileRouter) Kind() RouterKind { return RouterFRR }
func (r *reconcileRouter) Health(context.Context) error { return r.healthErr }
func (r *reconcileRouter) ApplyRoutes(context.Context, []RouteIntent) error { r.applied = true; return nil }
func (r *reconcileRouter) WithdrawRoutes(context.Context, []RouteIntent) error { r.withdrawn = true; return nil }
func (r *reconcileRouter) Snapshot(context.Context) ([]RouteIntent, error) { return nil, nil }

func validRoute() RouteIntent { return RouteIntent{Prefix: "203.0.113.0/24", Family: IPv4, NextHop: "192.0.2.1", Authorized: true} }

func TestRouteReconcilerBlocksUnauthorizedAndUnapproved(t *testing.T) {
	r := &reconcileRouter{}
	if err := (RouteReconciler{Router:r, Approved:true}).Apply(context.Background(), []RouteIntent{validRoute()}); err != ErrUnauthorized { t.Fatalf("got %v", err) }
	if err := (RouteReconciler{Router:r, Authorized:true}).Apply(context.Background(), []RouteIntent{validRoute()}); err != ErrApprovalRequired { t.Fatalf("got %v", err) }
	if r.applied { t.Fatal("router was mutated") }
}

func TestRouteReconcilerBlocksInvalidAndUnhealthy(t *testing.T) {
	r := &reconcileRouter{}
	bad := validRoute(); bad.Authorized = false
	c := RouteReconciler{Router:r, Authorized:true, Approved:true}
	if err := c.Apply(context.Background(), []RouteIntent{bad}); !errors.Is(err, ErrRouteUnauthorized) { t.Fatalf("got %v", err) }
	r.healthErr = errors.New("router unhealthy")
	if err := c.Apply(context.Background(), []RouteIntent{validRoute()}); err == nil { t.Fatal("expected health failure") }
	if r.applied { t.Fatal("router was mutated") }
}
