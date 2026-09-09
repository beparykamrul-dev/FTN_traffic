package dataplane

import "testing"

func TestValidatePurge(t *testing.T) {
	if err := ValidatePurge(true,true,"/x"); err != nil { t.Fatal(err) }
	if err := ValidatePurge(false,true,"/x"); err != ErrEdgeUnauthorized { t.Fatalf("got %v", err) }
	if err := ValidatePurge(true,false,"/x"); err != ErrPurgeApprovalRequired { t.Fatalf("got %v", err) }
}
