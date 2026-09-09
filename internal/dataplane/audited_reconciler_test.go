package dataplane

import (
	"context"
	"testing"
)

func TestAuditedReconcilerApply(t *testing.T) {
	a := &auditStub{}
	b := &routeBackendStub{}
	r := AuditedReconciler{
		Reconciler: RouteReconciler{Router:b, Authorized:true, Approved:true, Policy:&RoutePolicy{MaxPrefixes:10}},
		Auditor:a, Actor:"admin", RequestID:"req-apply", ApprovalID:"approval-1",
	}
	routes := []RouteIntent{{Prefix:"203.0.113.0/24",Family:IPv4,NextHop:"192.0.2.1",Authorized:true}}
	if err := r.Apply(context.Background(), routes); err != nil { t.Fatal(err) }
	if b.applied != 1 || len(a.events) != 1 || a.events[0].Action != "route.apply" { t.Fatalf("backend=%d audit=%#v",b.applied,a.events) }
}
