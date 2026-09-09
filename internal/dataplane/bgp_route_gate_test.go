package dataplane

import (
	"context"
	"testing"
)

type bfdGateStub struct { err error }
func (b bfdGateStub) CheckBFD(context.Context) error { return b.err }

func TestBGPRouteGate(t *testing.T) {
	g := BGPRouteGate{RequireRPKI:true, RequireBFD:true, RPKI:RPKIValid, BFD:bfdGateStub{}}
	if err := g.Check(context.Background()); err != nil { t.Fatal(err) }
	g.RPKI = RPKIInvalid
	if err := g.Check(context.Background()); err != ErrRPKIInvalid { t.Fatalf("got %v", err) }
	g.RPKI = RPKIValid; g.BFD = nil
	if err := g.Check(context.Background()); err != ErrBackendUnavailable { t.Fatalf("got %v", err) }
}
