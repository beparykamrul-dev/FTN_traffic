package dataplane

import "errors"

var ErrNoEstablishedBGP = errors.New("no established BGP session")

func ValidateBGPSessions(sessions []BGPSession, requireRPKI bool) error {
	if len(sessions) == 0 { return ErrNoEstablishedBGP }
	for _, s := range sessions {
		if !s.Established { continue }
		if requireRPKI && !s.RPKIValid { continue }
		return nil
	}
	return ErrNoEstablishedBGP
}
