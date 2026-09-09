package dataplane

import (
	"errors"
	"net/netip"
	"strconv"
	"strings"
)

var ErrGoBGPOutputUnrecognized = errors.New("unrecognized GoBGP session output")

func ParseGoBGPNeighborSummary(out string) ([]BGPSession, error) {
	var sessions []BGPSession
	for _, line := range strings.Split(out, "\n") {
		f := strings.Fields(line)
		if len(f) < 2 || strings.EqualFold(f[0], "neighbor") || strings.EqualFold(f[0], "peer") {
			continue
		}
		if _, err := netip.ParseAddr(f[0]); err != nil {
			continue
		}
		s := BGPSession{ID: f[0], IPv4: strings.Contains(f[0], "."), IPv6: strings.Contains(f[0], ":")}
		for _, x := range f[1:] {
			if n, err := strconv.ParseUint(x, 10, 32); err == nil && n > 0 {
				s.RemoteASN = uint32(n)
				break
			}
		}
		l := strings.ToLower(strings.Join(f, " "))
		s.Established = strings.Contains(l, "established") || strings.Contains(l, "estab") || strings.Contains(l, "up")
		sessions = append(sessions, s)
	}
	if len(sessions) == 0 {
		return nil, ErrGoBGPOutputUnrecognized
	}
	return sessions, nil
}
