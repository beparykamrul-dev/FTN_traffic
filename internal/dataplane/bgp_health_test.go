package dataplane

import "testing"

func TestValidateBGPSessions(t *testing.T) {
	if err := ValidateBGPSessions(nil, false); err != ErrNoEstablishedBGP { t.Fatalf("got %v", err) }
	if err := ValidateBGPSessions([]BGPSession{{Established:false}}, false); err != ErrNoEstablishedBGP { t.Fatalf("got %v", err) }
	if err := ValidateBGPSessions([]BGPSession{{ID:"p", Established:true, IPv4:true, RPKIValid:true}}, true); err != nil { t.Fatal(err) }
	if err := ValidateBGPSessions([]BGPSession{{ID:"p", Established:true, IPv4:true, RPKIValid:false}}, true); err != ErrNoEstablishedBGP { t.Fatalf("got %v", err) }
}

func TestSelectHealthyBGPSessionPrefersRPKIAndThenPrefixCount(t *testing.T) {
	s, err := SelectHealthyBGPSession([]BGPSession{
		{ID:"invalid", Established:true, IPv4:true, RPKIValid:false, PrefixesIn:1},
		{ID:"valid-more", Established:true, IPv4:true, RPKIValid:true, PrefixesIn:20},
		{ID:"valid-less", Established:true, IPv4:true, RPKIValid:true, PrefixesIn:10},
	}, true, 100)
	if err != nil || s.ID != "valid-less" { t.Fatalf("unexpected selection: %+v %v", s, err) }
}

func TestSelectHealthyBGPSessionRejectsMaxPrefix(t *testing.T) {
	_, err := SelectHealthyBGPSession([]BGPSession{{ID:"p", Established:true, IPv4:true, RPKIValid:true, PrefixesIn:101}}, true, 100)
	if err == nil { t.Fatal("expected max-prefix rejection") }
}
