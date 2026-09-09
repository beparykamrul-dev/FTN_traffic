package dataplane

import "context"

type AuditedReconciler struct { Reconciler RouteReconciler; Auditor Auditor; Actor string; RequestID string; ApprovalID string }

func (r AuditedReconciler) Apply(ctx context.Context, routes []RouteIntent) error {
	err := r.Reconciler.Apply(ctx, routes)
	result := "success"
	if err != nil { result = "failed" }
	if auditErr := AuditMutation(ctx, r.Auditor, AuditEvent{Action:"route.apply", Resource:"route", Actor:r.Actor, RequestID:r.RequestID, ApprovalID:r.ApprovalID, Result:result}); auditErr != nil && err == nil { return auditErr }
	return err
}

func (r AuditedReconciler) Withdraw(ctx context.Context, routes []RouteIntent) error {
	err := r.Reconciler.Withdraw(ctx, routes)
	result := "success"
	if err != nil { result = "failed" }
	if auditErr := AuditMutation(ctx, r.Auditor, AuditEvent{Action:"route.withdraw", Resource:"route", Actor:r.Actor, RequestID:r.RequestID, ApprovalID:r.ApprovalID, Result:result}); auditErr != nil && err == nil { return auditErr }
	return err
}
