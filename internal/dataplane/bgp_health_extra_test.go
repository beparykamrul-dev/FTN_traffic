package dataplane
import "testing"
func TestBGPSessionRejectsInvalidPeer(t *testing.T){s:=BGPSession{ID:"p",PeerAddress:"not-an-ip",Established:true,IPv4:true,RPKIValid:true};if s.Healthy(true,100){t.Fatal("expected invalid peer rejection")}}
func TestBGPSessionAllowsUnlimitedPrefixes(t *testing.T){s:=BGPSession{ID:"p",Established:true,IPv4:true,RPKIValid:true,PrefixesIn:^uint64(0)};if !s.Healthy(true,0){t.Fatal("expected unlimited prefix health")}}
