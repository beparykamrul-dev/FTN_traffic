package dataplane

import (
	"context"
	"errors"
	"testing"
)

type healthRouterStub struct{ err error }
func (r healthRouterStub) Name() string { return "health-router" }
func (r healthRouterStub) Health(context.Context) error { return r.err }
func (r healthRouterStub) ApplyRoutes(context.Context, []RouteIntent) error { return nil }
func (r healthRouterStub) WithdrawRoutes(context.Context, []RouteIntent) error { return nil }
func (r healthRouterStub) Snapshot(context.Context) (RouterStatus, error) { return RouterStatus{}, nil }

func TestHealthGateEnforcesMaxPrefixes(t *testing.T) {
	b := HealthGateBGPStub{RouterAdapter:healthRouterStub{}, Sessions:[]BGPSession{{ID:"p",PeerAddress:"192.0.2.2",Established:true,IPv4:true,RPKIValid:true,PrefixesIn:20}}}
	g := HealthGate{Router:b.RouterAdapter,BGP:b,RequireRPKI:true,MaxPrefixes:10}
	if err := g.Check(context.Background()); err == nil { t.Fatal("expected max-prefix failure") }
}

func TestHealthGateBackendFailure(t *testing.T) {
	g := HealthGate{Router:healthRouterStub{err:errors.New("down")},BGP:HealthGateBGPStub{RouterAdapter:healthRouterStub{}}}
	if err := g.Check(context.Background()); err == nil { t.Fatal("expected router failure") }
}

func TestHealthGateSelectsUsableSession(t *testing.T) {
	b := HealthGateBGPStub{RouterAdapter:healthRouterStub{}, Sessions:[]BGPSession{
		{ID:"bad",PeerAddress:"192.0.2.2",Established:true,IPv4:true,RPKIValid:false,PrefixesIn:1},
		{ID:"good",PeerAddress:"192.0.2.3",Established:true,IPv4:true,RPKIValid:true,PrefixesIn:5},
	}}
	g := HealthGate{Router:b.RouterAdapter,BGP:b,RequireRPKI:true,MaxPrefixes:10}
	if err := g.Check(context.Background()); err != nil { t.Fatal(err) }
}
