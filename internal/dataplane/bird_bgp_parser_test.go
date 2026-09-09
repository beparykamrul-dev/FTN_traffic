package dataplane

import "testing"

func TestParseBIRDBGPSummary(t *testing.T) {
	out := "Name Proto Table State Since Info\npeer1 BGP --- up 2026-09-09 Established AS64510\n"
	s, err := ParseBIRDBGPSummary(out)
	if err != nil || len(s) != 1 || s[0].ID != "peer1" || !s[0].Established || s[0].RemoteASN != 64510 {
		t.Fatalf("unexpected parse: %#v %v", s, err)
	}
}

func TestParseBIRDBGPSummaryIgnoresNonBGP(t *testing.T) {
	out := "Name Proto Table State Since Info\nstatic Static master up\npeer2 BGP master down 2026-09-09 Active AS64511\n"
	s, err := ParseBIRDBGPSummary(out)
	if err != nil || len(s) != 1 || s[0].ID != "peer2" || s[0].Established || s[0].RemoteASN != 64511 {
		t.Fatalf("unexpected parse: %#v %v", s, err)
	}
}
