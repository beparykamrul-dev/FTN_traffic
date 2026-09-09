package dataplane

import "context"

type RouteOrchestrator struct {
	Reconciler RouteReconciler
	Health HealthGate
	Auditor Auditor
	Actor string
	RequestID string
	ApprovalID string
}

func (o RouteOrchestrator) Apply(ctx context.Context, routes []RouteIntent) error {
	normalized, err := NormalizeRoutes(routes); if err != nil { return err }
	if err := o.Health.Check(ctx); err != nil { return err }
	o.Reconciler.RPKIValid = o.Reconciler.RPKIValid
	err = o.Reconciler.Apply(ctx, normalized)
	result := "success"; if err != nil { result = "failed" }
	if auditErr := AuditMutation(ctx, o.Auditor, AuditEvent{Action:"route.orchestrate.apply", Resource:"route", Actor:o.Actor, RequestID:o.RequestID, ApprovalID:o.ApprovalID, Result:result}); auditErr != nil && err == nil { return auditErr }
	return err
}

func (o RouteOrchestrator) Withdraw(ctx context.Context, routes []RouteIntent) error {
	normalized, err := NormalizeRoutes(routes); if err != nil { return err }
	if err := o.Health.Check(ctx); err != nil { return err }
	err = o.Reconciler.Withdraw(ctx, normalized)
	result := "success"; if err != nil { result = "failed" }
	if auditErr := AuditMutation(ctx, o.Auditor, AuditEvent{Action:"route.orchestrate.withdraw", Resource:"route", Actor:o.Actor, RequestID:o.RequestID, ApprovalID:o.ApprovalID, Result:result}); auditErr != nil && err == nil { return auditErr }
	return err
}
