package dataplane

import "testing"

func TestValidateBFDPolicy(t *testing.T) {
	s := BFDSession{ID:"bfd-1",Local:"192.0.2.1",Remote:"192.0.2.2",MinRxMS:300,MinTxMS:300,Multiplier:3,Up:true,Authorized:true}
	if err := ValidateBFDPolicy(s); err != nil { t.Fatal(err) }
	s.Up=false
	if err := ValidateBFDPolicy(s); err != ErrBFDRequired { t.Fatalf("got %v",err) }
}
