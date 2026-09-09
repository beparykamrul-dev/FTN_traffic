package dataplane

import "testing"

func TestValidateAuditEvent(t *testing.T) {
	good := AuditEvent{Action:"route.apply",Resource:"router-1",RequestID:"req-1",ApprovalID:"app-1"}
	if err := ValidateAuditEvent(good); err != nil { t.Fatal(err) }
	bad := good; bad.RequestID = "req\n1"
	if err := ValidateAuditEvent(bad); err == nil { t.Fatal("expected audit rejection") }
}
