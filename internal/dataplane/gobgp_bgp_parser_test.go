package dataplane

import "testing"

func TestParseGoBGPNeighborSummary(t *testing.T) {
	out := "Peer AS State\n192.0.2.3 64512 Established\n"
	s, err := ParseGoBGPNeighborSummary(out)
	if err != nil || len(s) != 1 || s[0].RemoteASN != 64512 || !s[0].Established || !s[0].IPv4 {
		t.Fatalf("unexpected parse: %#v %v", s, err)
	}
}

func TestParseGoBGPNeighborSummaryIPv6AndGarbage(t *testing.T) {
	out := "Peer AS State\ninvalid 64513 Established\n2001:db8::3 64513 Idle\n"
	s, err := ParseGoBGPNeighborSummary(out)
	if err != nil || len(s) != 1 || !s[0].IPv6 || s[0].Established || s[0].RemoteASN != 64513 {
		t.Fatalf("unexpected parse: %#v %v", s, err)
	}
}
