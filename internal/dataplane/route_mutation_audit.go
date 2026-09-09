package dataplane

import "context"

type RouteMutationAuditor struct { Auditor Auditor }

func (a RouteMutationAuditor) RecordRoute(ctx context.Context, action, actor, requestID, approvalID, result string) error {
	return AuditMutation(ctx, a.Auditor, AuditEvent{Action: action, Resource: "route", Actor: actor, RequestID: requestID, ApprovalID: approvalID, Result: result})
}
