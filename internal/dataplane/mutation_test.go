package dataplane

import "testing"

func TestValidateMutation(t *testing.T) {
	if err := ValidateMutation(true, true, true); err != nil { t.Fatal(err) }
	if err := ValidateMutation(false, true, true); err != ErrUnauthorized { t.Fatalf("got %v", err) }
	if err := ValidateMutation(true, false, true); err != ErrApprovalRequired { t.Fatalf("got %v", err) }
	if err := ValidateMutation(true, true, false); err != ErrBackendUnavailable { t.Fatalf("got %v", err) }
}
