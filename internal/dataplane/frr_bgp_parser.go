package dataplane

import (
	"errors"
	"net/netip"
	"strconv"
	"strings"
)

var ErrBGPOutputUnrecognized = errors.New("unrecognized BGP session output")

func ParseFRRBGPSummary(out string) ([]BGPSession, error) {
	var sessions []BGPSession
	for _, line := range strings.Split(out, "\n") {
		f := strings.Fields(line)
		if len(f) < 5 || strings.EqualFold(f[0], "Neighbor") || strings.HasPrefix(strings.ToLower(f[0]), "bgp") { continue }
		peer, err := netip.ParseAddr(f[0]); if err != nil { continue }
		asn, err := strconv.ParseUint(f[2], 10, 32); if err != nil || asn == 0 { continue }
		s := BGPSession{ID: f[0], PeerAddress: peer.String(), RemoteASN: uint32(asn), IPv4: peer.Is4(), IPv6: peer.Is6()}
		last := f[len(f)-1]
		if n, err := strconv.ParseUint(last, 10, 64); err == nil { s.PrefixesIn, s.Established = n, true } else { state := strings.ToLower(last); s.Established = state == "established" || state == "estab" }
		sessions = append(sessions, s)
	}
	if len(sessions) == 0 { return nil, ErrBGPOutputUnrecognized }
	return sessions, nil
}
