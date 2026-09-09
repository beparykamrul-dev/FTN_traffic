package dataplane

import "testing"

func TestParseBGPSummaryDispatchesByRouterKind(t *testing.T) {
	frr := "Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd\n192.0.2.9 4 64509 1 1 0 0 0 00:10 3\n"
	sessions, err := ParseBGPSummary(RouterFRR, frr)
	if err != nil || len(sessions) != 1 || sessions[0].RemoteASN != 64509 { t.Fatalf("unexpected FRR result: err=%v sessions=%#v", err, sessions) }
}

func TestParseBGPSummaryDispatchesAllSupportedRouters(t *testing.T) {
	cases := []struct{name string; kind RouterKind; output string}{
		{"bird", RouterBIRD, "Name Proto Table State\npeer1 BGP master up AS64510\n"},
		{"gobgp", RouterGoBGP, "Peer AS State\n192.0.2.3 64511 Established\n"},
	}
	for _, tc := range cases { t.Run(tc.name, func(t *testing.T) { s, err := ParseBGPSummary(tc.kind, tc.output); if err != nil || len(s) != 1 { t.Fatalf("err=%v sessions=%#v", err, s) } }) }
}

func TestParseBGPSummaryRejectsUnknownRouter(t *testing.T) {
	if _, err := ParseBGPSummary(RouterKind("unknown"), ""); err != ErrBGPOutputUnrecognized { t.Fatalf("expected ErrBGPOutputUnrecognized, got %v", err) }
}

func TestParseBGPSummaryHandlesEmptyOutput(t *testing.T) {
	if _, err := ParseBGPSummary(RouterFRR, ""); err != ErrBGPOutputUnrecognized { t.Fatalf("expected parser error for empty output, got %v", err) }
}
