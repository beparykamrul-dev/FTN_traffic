package dataplane

import "testing"

func TestEnforceRPKI(t *testing.T) {
	if err := EnforceRPKI(RPKIUnknown, false); err != nil { t.Fatal(err) }
	if err := EnforceRPKI(RPKIValid, true); err != nil { t.Fatal(err) }
	if err := EnforceRPKI(RPKIInvalid, true); err != ErrRPKIInvalid { t.Fatalf("got %v", err) }
	if err := EnforceRPKI(RPKIUnknown, true); err != ErrRPKIUnknown { t.Fatalf("got %v", err) }
}
