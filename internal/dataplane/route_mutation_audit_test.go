package dataplane

import (
	"context"
	"testing"
)

type auditStub struct { events []AuditEvent }
func (a *auditStub) Record(_ context.Context, e AuditEvent) error { a.events = append(a.events, e); return nil }

func TestRouteMutationAuditor(t *testing.T) {
	a := &auditStub{}
	r := RouteMutationAuditor{Auditor:a}
	if err := r.RecordRoute(context.Background(), "route.apply", "admin", "req-1", "app-1", "success"); err != nil { t.Fatal(err) }
	if len(a.events) != 1 || a.events[0].Resource != "route" { t.Fatalf("unexpected events: %#v", a.events) }
}
