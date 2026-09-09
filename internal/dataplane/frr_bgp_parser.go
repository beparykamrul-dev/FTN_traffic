package dataplane

import (
	"errors"
	"strconv"
	"strings"
)

var ErrBGPOutputUnrecognized = errors.New("unrecognized BGP session output")

func parseUint(s string) uint64 {
	v, _ := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
	return v
}

func ParseFRRBGPSummary(out string) ([]BGPSession, error) {
	var sessions []BGPSession
	for _, line := range strings.Split(out, "\n") {
		f := strings.Fields(line)
		if len(f) < 5 || strings.EqualFold(f[0], "Neighbor") || strings.HasPrefix(f[0], "BGP") {
			continue
		}
		if !strings.Contains(f[0], ".") && !strings.Contains(f[0], ":") {
			continue
		}
		s := BGPSession{ID: f[0]}
		if n, err := strconv.ParseUint(f[2], 10, 32); err == nil {
			s.RemoteASN = uint32(n)
		} else {
			continue
		}
		last := f[len(f)-1]
		if n, err := strconv.ParseUint(last, 10, 64); err == nil {
			s.PrefixesIn = n
			s.Established = true
		} else {
			state := strings.ToLower(last)
			s.Established = state == "established" || state == "estab"
		}
		s.IPv4 = strings.Contains(f[0], ".")
		s.IPv6 = strings.Contains(f[0], ":")
		sessions = append(sessions, s)
	}
	if len(sessions) == 0 {
		return nil, ErrBGPOutputUnrecognized
	}
	return sessions, nil
}
