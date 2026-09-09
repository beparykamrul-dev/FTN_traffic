package dataplane

import "errors"

var ErrNoEstablishedBGP = errors.New("no established BGP session")

func ValidateBGPSessions(sessions []BGPSession, requireRPKI bool) error {
	return ValidateBGPSessionsWithLimit(sessions, requireRPKI, 0)
}

func ValidateBGPSessionsWithLimit(sessions []BGPSession, requireRPKI bool, maxPrefixes uint64) error {
	if len(sessions) == 0 { return ErrNoEstablishedBGP }
	if _, err := SelectHealthyBGPSession(sessions, requireRPKI, maxPrefixes); err != nil { return ErrNoEstablishedBGP }
	return nil
}
