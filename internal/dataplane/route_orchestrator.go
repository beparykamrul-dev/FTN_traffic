package dataplane

import (
	"context"
	"fmt"
)

type RouteOrchestrator struct {
	Reconciler RouteReconciler
	Health HealthGate
	Auditor Auditor
	Actor string
	RequestID string
	ApprovalID string
}

func (o RouteOrchestrator) validateMutationMetadata() error {
	if o.Auditor == nil || o.Actor == "" || o.RequestID == "" || o.ApprovalID == "" { return fmt.Errorf("route mutation metadata incomplete") }
	return nil
}

func (o RouteOrchestrator) Apply(ctx context.Context, routes []RouteIntent) error {
	if err := o.validateMutationMetadata(); err != nil { return err }
	normalized, err := NormalizeRoutes(routes); if err != nil { return err }
	if err := ValidateRouteSetForMutation(normalized); err != nil { return err }
	if err := o.Health.Check(ctx); err != nil { return err }
	err = o.Reconciler.Apply(ctx, normalized)
	result := "success"; if err != nil { result = "failed" }
	if auditErr := AuditMutation(ctx, o.Auditor, AuditEvent{Action:"route.orchestrate.apply", Resource:"route", Actor:o.Actor, RequestID:o.RequestID, ApprovalID:o.ApprovalID, Result:result}); auditErr != nil && err == nil { return auditErr }
	return err
}

func (o RouteOrchestrator) Withdraw(ctx context.Context, routes []RouteIntent) error {
	if err := o.validateMutationMetadata(); err != nil { return err }
	normalized, err := NormalizeRoutes(routes); if err != nil { return err }
	if err := ValidateRouteSetForMutation(normalized); err != nil { return err }
	if err := o.Health.Check(ctx); err != nil { return err }
	err = o.Reconciler.Withdraw(ctx, normalized)
	result := "success"; if err != nil { result = "failed" }
	if auditErr := AuditMutation(ctx, o.Auditor, AuditEvent{Action:"route.orchestrate.withdraw", Resource:"route", Actor:o.Actor, RequestID:o.RequestID, ApprovalID:o.ApprovalID, Result:result}); auditErr != nil && err == nil { return auditErr }
	return err
}
