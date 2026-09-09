package dataplane

import "testing"

func TestParseBGPSummaryDispatchesByRouterKind(t *testing.T) {
	frr := "Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd\n192.0.2.9 4 64509 1 1 0 0 0 00:10 3\n"
	sessions, err := ParseBGPSummary(RouterFRR, frr)
	if err != nil || len(sessions) != 1 || sessions[0].RemoteASN != 64509 {
		t.Fatalf("unexpected FRR result: err=%v sessions=%#v", err, sessions)
	}
}

func TestParseBGPSummaryRejectsUnknownRouter(t *testing.T) {
	if _, err := ParseBGPSummary(RouterKind("unknown"), ""); err != ErrBGPOutputUnrecognized {
		t.Fatalf("expected ErrBGPOutputUnrecognized, got %v", err)
	}
}

func TestParseBGPSummaryHandlesEmptyOutput(t *testing.T) {
	if _, err := ParseBGPSummary(RouterFRR, ""); err != ErrBGPOutputUnrecognized {
		t.Fatalf("expected parser error for empty output, got %v", err)
	}
}
