package dataplane

import "errors"

var ErrNoHealthyBGPSession = errors.New("no healthy BGP session")

func SelectHealthyBGPSession(sessions []BGPSession, requireRPKI bool, maxPrefixes uint64) (BGPSession, error) {
	for _, s := range sessions {
		if s.Healthy(requireRPKI, maxPrefixes) { return s, nil }
	}
	return BGPSession{}, ErrNoHealthyBGPSession
}
