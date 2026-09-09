package dataplane

import "testing"

func TestValidateBGPSession(t *testing.T) {
	good := BGPSession{ID: "p1", PeerAddress: "192.0.2.1", ASN: 64512}
	if err := ValidateBGPSession(good); err != nil { t.Fatal(err) }
	bad := good
	bad.PeerAddress = "not-an-ip"
	if err := ValidateBGPSession(bad); err != ErrInvalidBGPSession { t.Fatalf("got %v", err) }
}

func TestValidateBGPSessionSetRequiresEstablishedRPKI(t *testing.T) {
	s := []BGPSession{{ID:"p1", PeerAddress:"192.0.2.1", ASN:64512, Established:true, RPKIValid:false}}
	if err := ValidateBGPSessionSet(s, true); err != ErrNoEstablishedBGP { t.Fatalf("got %v", err) }
	s[0].RPKIValid = true
	if err := ValidateBGPSessionSet(s, true); err != nil { t.Fatal(err) }
}
