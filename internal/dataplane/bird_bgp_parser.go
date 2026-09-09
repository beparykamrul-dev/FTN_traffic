package dataplane

import (
	"errors"
	"net/netip"
	"strconv"
	"strings"
)

var ErrBIRDOutputUnrecognized = errors.New("unrecognized BIRD session output")

func ParseBIRDBGPSummary(out string) ([]BGPSession, error) {
	var sessions []BGPSession
	for _, line := range strings.Split(out, "\n") {
		f := strings.Fields(line); if len(f) < 2 { continue }
		if strings.EqualFold(f[0], "name") || strings.EqualFold(f[0], "bird") { continue }
		joined := strings.ToLower(strings.Join(f, " ")); if !strings.Contains(joined, "bgp") { continue }
		s := BGPSession{ID: f[0]}
		if p, err := netip.ParseAddr(f[0]); err == nil { s.PeerAddress = p.String(); s.IPv4 = p.Is4(); s.IPv6 = p.Is6() }
		for _, x := range f[1:] { if n, err := strconv.ParseUint(x, 10, 32); err == nil && n > 0 { s.RemoteASN = uint32(n); break } }
		s.Established = strings.Contains(joined, "established") || strings.Contains(joined, " up") || strings.HasSuffix(joined, "up")
		if !s.IPv4 && !s.IPv6 { s.IPv4 = s.Established }
		sessions = append(sessions, s)
	}
	if len(sessions) == 0 { return nil, ErrBIRDOutputUnrecognized }
	return sessions, nil
}
