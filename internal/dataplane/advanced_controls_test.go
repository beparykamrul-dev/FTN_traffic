package dataplane

import "testing"

func TestAdvancedControls(t *testing.T) {
	if err:=ValidateRPKI(RPKIValid,true);err!=nil{t.Fatal(err)}
	if err:=ValidateRPKI(RPKIInvalid,true);err!=ErrRPKIInvalid{t.Fatalf("got %v",err)}
	if err:=(VLAN{ID:100,Parent:"eth0",Authorized:true}).Validate();err!=nil{t.Fatal(err)}
	if err:=(VRF{Name:"customer-a",Table:1001,Authorized:true}).Validate();err!=nil{t.Fatal(err)}
	if err:=(QoSPolicy{ID:"gold",RateMbps:100,BurstKB:64,DSCP:[]uint8{46},Authorized:true}).Validate();err!=nil{t.Fatal(err)}
	if err:=(NATPolicy{ID:"nat-a",SourcePrefix:"10.0.0.0/24",EgressInterface:"wan0",Authorized:true}).Validate();err!=nil{t.Fatal(err)}
}

func TestBFDRejectsMixedAddressFamilies(t *testing.T) {
	b:=BFDSession{ID:"b1",Local:"192.0.2.1",Remote:"2001:db8::1",MinRxMS:300,MinTxMS:300,Multiplier:3,Authorized:true}
	if err:=b.Validate();err!=ErrBFDInvalid{t.Fatalf("expected BFD family error, got %v",err)}
}

func TestRPKIRequiredRejectsUnknown(t *testing.T) {
	if err:=ValidateRPKI(RPKIUnknown,true);err!=ErrRPKIInvalid{t.Fatalf("expected RPKI error, got %v",err)}
}

func TestAdvancedControlsRejectUnauthorizedAndInvalidRanges(t *testing.T) {
	if err:=(VLAN{ID:4095,Parent:"eth0",Authorized:true}).Validate();err!=ErrVLANInvalid{t.Fatalf("expected VLAN range error, got %v",err)}
	if err:=(VRF{Name:"customer-a",Table:1001,Authorized:false}).Validate();err!=ErrVRFInvalid{t.Fatalf("expected VRF authorization error, got %v",err)}
}

func TestBFDRejectsZeroTimers(t *testing.T) {
	b:=BFDSession{ID:"b2",Local:"192.0.2.1",Remote:"192.0.2.2",MinRxMS:0,MinTxMS:50,Multiplier:3,Authorized:true}
	if err:=b.Validate();err!=ErrBFDInvalid{t.Fatalf("expected timer error, got %v",err)}
}
