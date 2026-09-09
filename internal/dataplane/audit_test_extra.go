package dataplane

import (
	"context"
	"testing"
)

func TestAuditMutationFailsClosed(t *testing.T) {
	if err := AuditMutation(context.Background(), nil, AuditEvent{Action:"route.apply",Resource:"route",RequestID:"r",ApprovalID:"a"}); err == nil { t.Fatal("expected unavailable auditor failure") }
}

func TestAuditMutationRequiresApproval(t *testing.T) {
	a := &orchestratorAuditStub{}
	if err := AuditMutation(context.Background(), a, AuditEvent{Action:"route.apply",Resource:"route",RequestID:"r"}); err == nil { t.Fatal("expected incomplete audit metadata") }
}
