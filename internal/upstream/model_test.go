package upstream

import "testing"

func TestSessionEligible(t *testing.T) {
	s := Session{LocalASN: 65001, RemoteASN: 65002, IPv4: true, Established: true, Authorized: true, MaxPrefixes: 100}
	if !s.Eligible() || !CanProvideUpstream(s) {
		t.Fatal("authorized established session should be eligible")
	}
}

func TestSessionFailsClosed(t *testing.T) {
	cases := []Session{
		{LocalASN: 65001, RemoteASN: 65002, IPv4: true, Established: true, Authorized: false, MaxPrefixes: 100},
		{LocalASN: 65001, RemoteASN: 65002, IPv4: true, Established: false, Authorized: true, MaxPrefixes: 100},
		{LocalASN: 0, RemoteASN: 65002, IPv4: true, Established: true, Authorized: true, MaxPrefixes: 100},
		{LocalASN: 65001, RemoteASN: 65002, IPv4: true, Established: true, Authorized: true, MaxPrefixes: 0},
	}
	for _, s := range cases {
		if s.Eligible() || CanProvideUpstream(s) {
			t.Fatalf("session should fail closed: %+v", s)
		}
	}
}
