package dataplane

import (
	"context"
	"testing"
)

type orchestratorBGPStub struct { RouterAdapter }
func (orchestratorBGPStub) SessionSummary(context.Context) ([]BGPSession,error) { return []BGPSession{{ID:"peer",PeerAddress:"192.0.2.2",Established:true,IPv4:true,RPKIValid:true,PrefixesIn:10}},nil }
type orchestratorAuditStub struct { events int }
func (a *orchestratorAuditStub) Record(context.Context, AuditEvent) error { a.events++; return nil }

func TestRouteOrchestratorApply(t *testing.T) {
	r := healthRouterStub{}
	o := RouteOrchestrator{Reconciler:RouteReconciler{Router:r,Authorized:true,Approved:true},Health:HealthGate{Router:r,BGP:orchestratorBGPStub{RouterAdapter:r},RequireRPKI:true,MaxPrefixes:20},Auditor:&orchestratorAuditStub{},Actor:"admin",RequestID:"req-1",ApprovalID:"appr-1"}
	if err := o.Apply(context.Background(), []RouteIntent{{Prefix:"192.0.2.0/24",Family:IPv4,Authorized:true}}); err != nil { t.Fatal(err) }
}

func TestRouteOrchestratorRejectsUnapproved(t *testing.T) {
	r := healthRouterStub{}
	o := RouteOrchestrator{Reconciler:RouteReconciler{Router:r,Authorized:true,Approved:false},Health:HealthGate{Router:r,BGP:orchestratorBGPStub{RouterAdapter:r}},Auditor:&orchestratorAuditStub{},Actor:"admin",RequestID:"req-1",ApprovalID:"appr-1"}
	if err := o.Apply(context.Background(), []RouteIntent{{Prefix:"192.0.2.0/24",Family:IPv4,Authorized:true}}); err == nil { t.Fatal("expected approval failure") }
}
