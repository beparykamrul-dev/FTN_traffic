package dataplane
import "testing"
func TestParseFRRBGPSummary(t *testing.T){out:="Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd\n192.0.2.1 4 64500 10 10 0 0 0 01:00 12";s,e:=ParseFRRBGPSummary(out);if e!=nil||len(s)!=1||s[0].RemoteASN!=64500||!s[0].Established||s[0].PrefixesIn!=12{t.Fatalf("unexpected parse: %#v %v",s,e)}}
