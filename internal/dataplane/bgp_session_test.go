package dataplane

import "testing"

func TestBGPSessionHealthy(t *testing.T) {
	s := BGPSession{ID: "peer-1", RemoteASN: 64500, Established: true, IPv4: true, RPKIValid: true, PrefixesIn: 100}
	if !s.Healthy(true, 1000) { t.Fatal("expected healthy session") }
	s.Established = false
	if s.Healthy(false, 1000) { t.Fatal("expected down session") }
}

func TestBGPSessionRejectsPrefixOverflowAndRPKI(t *testing.T) {
	s := BGPSession{ID: "peer-1", RemoteASN: 64500, Established: true, IPv4: true, RPKIValid: false, PrefixesIn: 1001}
	if s.Healthy(true, 2000) { t.Fatal("expected RPKI rejection") }
	s.RPKIValid = true
	if s.Healthy(false, 1000) { t.Fatal("expected max-prefix rejection") }
}
