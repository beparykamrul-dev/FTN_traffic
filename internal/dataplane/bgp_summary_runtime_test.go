package dataplane

import (
	"context"
	"testing"
)

type summarySourceStub struct{ out string; err error }
func (s summarySourceStub) SummaryOutput(context.Context) (string,error) { return s.out,s.err }

func TestCollectBGPSessions(t *testing.T) {
	src := summarySourceStub{out: "Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd\n203.0.113.2 4 64500 1 1 0 0 0 00:01 3"}
	got, err := CollectBGPSessions(context.Background(), RouterFRR, src, false, 10)
	if err != nil || len(got) != 1 || !got[0].Established { t.Fatalf("unexpected result: %#v %v", got, err) }
}
