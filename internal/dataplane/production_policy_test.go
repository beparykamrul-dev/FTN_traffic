package dataplane

import "testing"

func TestProductionPolicyFailClosed(t *testing.T) {
	if err := (ProductionPolicy{RequireApproval:true, RequireAuthorization:true, FailClosed:true}).Validate(); err != nil { t.Fatal(err) }
	if err := (ProductionPolicy{RequireApproval:true, RequireAuthorization:true, FailClosed:false}).Validate(); err == nil { t.Fatal("expected fail-closed rejection") }
	if err := (ProductionPolicy{RequireApproval:true, RequireAuthorization:true, FailClosed:true, CapturePayload:true}).Validate(); err == nil { t.Fatal("expected payload capture rejection") }
}

func TestValidateResourceID(t *testing.T) {
	for _, id := range []string{"", "x\n", "x\r", "x\x00"} { if err := ValidateResourceID(id); err == nil { t.Fatalf("accepted %q", id) } }
	if err := ValidateResourceID("ftn-router-01"); err != nil { t.Fatal(err) }
}
