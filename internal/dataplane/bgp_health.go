package dataplane

import (
	"errors"
	"fmt"
)

var ErrNoEstablishedBGP = errors.New("no established BGP session")

// SelectHealthyBGPSession returns the best session satisfying the health policy.
// It only selects; it never changes routing state.
func SelectHealthyBGPSession(sessions []BGPSession, requireRPKI bool, maxPrefixes uint64) (BGPSession, error) {
	var best BGPSession
	found := false
	for _, s := range sessions {
		if !s.Healthy(requireRPKI, maxPrefixes) {
			continue
		}
		if !found || betterBGPSession(s, best, requireRPKI) {
			best = s
			found = true
		}
	}
	if !found {
		return BGPSession{}, fmt.Errorf("%w: no session satisfies health policy", ErrNoEstablishedBGP)
	}
	return best, nil
}

func betterBGPSession(a, b BGPSession, requireRPKI bool) bool {
	if requireRPKI && a.RPKIValid != b.RPKIValid {
		return a.RPKIValid
	}
	if a.PrefixesIn != b.PrefixesIn {
		return a.PrefixesIn < b.PrefixesIn
	}
	return a.ID < b.ID
}

func ValidateBGPSessions(sessions []BGPSession, requireRPKI bool) error {
	return ValidateBGPSessionsWithLimit(sessions, requireRPKI, 0)
}

func ValidateBGPSessionsWithLimit(sessions []BGPSession, requireRPKI bool, maxPrefixes uint64) error {
	if len(sessions) == 0 {
		return ErrNoEstablishedBGP
	}
	if _, err := SelectHealthyBGPSession(sessions, requireRPKI, maxPrefixes); err != nil {
		return ErrNoEstablishedBGP
	}
	return nil
}
