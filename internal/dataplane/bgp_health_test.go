package dataplane

import "testing"

func TestValidateBGPSessions(t *testing.T) {
	if err := ValidateBGPSessions(nil, false); err != ErrNoEstablishedBGP { t.Fatalf("got %v", err) }
	if err := ValidateBGPSessions([]BGPSession{{Established:false}}, false); err != ErrNoEstablishedBGP { t.Fatalf("got %v", err) }
	if err := ValidateBGPSessions([]BGPSession{{Established:true, RPKIValid:true}}, true); err != nil { t.Fatal(err) }
	if err := ValidateBGPSessions([]BGPSession{{Established:true, RPKIValid:false}}, true); err != ErrNoEstablishedBGP { t.Fatalf("got %v", err) }
}
