package dataplane

import "testing"

func TestSelectHealthyBGPSession(t *testing.T) {
	sessions := []BGPSession{
		{ID:"bad", PeerAddress:"192.0.2.2", Established:true, IPv4:true, RPKIValid:false, PrefixesIn:2},
		{ID:"good", PeerAddress:"192.0.2.3", Established:true, IPv4:true, RPKIValid:true, PrefixesIn:5},
	}
	s, err := SelectHealthyBGPSession(sessions, true, 10)
	if err != nil || s.ID != "good" { t.Fatalf("session=%+v err=%v", s, err) }
	if _, err := SelectHealthyBGPSession(sessions, true, 4); err != ErrNoHealthyBGPSession { t.Fatalf("got %v", err) }
}
