package dataplane

import "testing"

func TestValidateBFD(t *testing.T) {
	good := BFDSession{ID:"b1", Local:"192.0.2.1", Remote:"192.0.2.2", MinRxMS:50, MinTxMS:50, Multiplier:3, Authorized:true}
	if err := ValidateBFD(good); err != nil { t.Fatal(err) }
	good.Remote = good.Local
	if err := ValidateBFD(good); err != ErrBFDInvalid { t.Fatalf("got %v", err) }
}
