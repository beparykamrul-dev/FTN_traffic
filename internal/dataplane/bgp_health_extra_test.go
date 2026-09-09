package dataplane

import "testing"

func TestBGPSessionRejectsInvalidPeer(t *testing.T) {
	s := BGPSession{ID: "p", PeerAddress: "not-an-ip", Established: true, IPv4: true, RPKIValid: true}
	if s.Healthy(true, 100) { t.Fatal("expected invalid peer rejection") }
}

func TestBGPSessionAllowsUnlimitedPrefixes(t *testing.T) {
	s := BGPSession{ID: "p", Established: true, IPv4: true, RPKIValid: true, PrefixesIn: ^uint64(0)}
	if !s.Healthy(true, 0) { t.Fatal("expected unlimited prefix health") }
}

func TestValidateBGPSessionsWithLimit(t *testing.T) {
	s := BGPSession{ID: "p1", PeerAddress: "203.0.113.2", Established: true, IPv4: true, RPKIValid: true, PrefixesIn: 5}
	if err := ValidateBGPSessionsWithLimit([]BGPSession{s}, true, 10); err != nil { t.Fatalf("expected healthy session: %v", err) }
	if err := ValidateBGPSessionsWithLimit([]BGPSession{s}, true, 4); err == nil { t.Fatal("expected max-prefix rejection") }
}

func TestValidateBGPSessionsWithLimitNoSessions(t *testing.T) {
	if err := ValidateBGPSessionsWithLimit(nil, false, 0); err != ErrNoEstablishedBGP { t.Fatalf("unexpected error: %v", err) }
}
