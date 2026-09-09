package dataplane

import "testing"

func TestBGPHealth(t *testing.T) {
	if err := (BGPHealth{Established:true, RPKIValid:true}).Validate(true); err != nil { t.Fatal(err) }
	if err := (BGPHealth{}).Validate(false); err != ErrBGPDown { t.Fatalf("got %v", err) }
	if err := (BGPHealth{Established:true}).Validate(true); err != ErrRPKIInvalid { t.Fatalf("got %v", err) }
}
