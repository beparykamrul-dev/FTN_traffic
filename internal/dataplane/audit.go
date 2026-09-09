package dataplane

import (
	"context"
	"fmt"
)

type AuditEvent struct { Action string; Resource string; Actor string; RequestID string; ApprovalID string; Result string }
type Auditor interface { Record(context.Context, AuditEvent) error }

func AuditMutation(ctx context.Context, a Auditor, e AuditEvent) error {
	if a == nil { return fmt.Errorf("audit backend unavailable") }
	if e.Action == "" || e.Resource == "" || e.RequestID == "" || e.ApprovalID == "" { return fmt.Errorf("audit fields incomplete") }
	return a.Record(ctx, e)
}
