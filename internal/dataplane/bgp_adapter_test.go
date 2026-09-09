package dataplane
import "testing"
func TestParseBGPSummaryDispatch(t *testing.T){out:="Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd\n192.0.2.9 4 64509 1 1 0 0 0 00:10 3";s,e:=ParseBGPSummary("frr",out);if e!=nil||len(s)!=1||s[0].RemoteASN!=64509{t.Fatalf("dispatch failed: %#v %v",s,e)}}
