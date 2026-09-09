package dataplane

import "testing"

func TestParseFRRBGPSummary(t *testing.T) {
	out := "Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd\n192.0.2.1 4 64500 10 10 0 0 0 01:00 12\n"
	s, err := ParseFRRBGPSummary(out)
	if err != nil || len(s) != 1 || s[0].RemoteASN != 64500 || !s[0].Established || s[0].PrefixesIn != 12 || !s[0].IPv4 {
		t.Fatalf("unexpected parse: %#v %v", s, err)
	}
}

func TestParseFRRBGPSummaryIgnoresHeadersAndMalformedPeers(t *testing.T) {
	out := "BGP router identifier 192.0.2.1\nNeighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd\nnot-an-ip x y z\n2001:db8::2 4 64501 1 1 0 0 0 00:01 7\n"
	s, err := ParseFRRBGPSummary(out)
	if err != nil || len(s) != 1 || !s[0].IPv6 || s[0].RemoteASN != 64501 {
		t.Fatalf("unexpected parse: %#v %v", s, err)
	}
}

func TestParseFRRBGPSummaryRejectsEmptyOutput(t *testing.T) {
	if _, err := ParseFRRBGPSummary("garbage"); err != ErrBGPOutputUnrecognized {
		t.Fatalf("expected parser error, got %v", err)
	}
}
