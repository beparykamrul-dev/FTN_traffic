package dataplane

import (
	"errors"
	"net/netip"
)

var ErrInvalidBGPSession = errors.New("invalid BGP session")

func ValidateBGPSession(s BGPSession) error {
	if s.ID == "" || s.ASN == 0 || s.ASN > 4294967295 {
		return ErrInvalidBGPSession
	}
	addr, err := netip.ParseAddr(s.PeerAddress)
	if err != nil || !addr.IsGlobalUnicast() {
		return ErrInvalidBGPSession
	}
	return nil
}

func ValidateBGPSessionSet(sessions []BGPSession, requireRPKI bool) error {
	if len(sessions) == 0 {
		return ErrNoEstablishedBGP
	}
	for _, s := range sessions {
		if ValidateBGPSession(s) != nil || !s.Established {
			continue
		}
		if requireRPKI && !s.RPKIValid {
			continue
		}
		return nil
	}
	return ErrNoEstablishedBGP
}
