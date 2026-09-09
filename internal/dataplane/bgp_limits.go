package dataplane

import "errors"

var ErrBGPFamilyUnavailable = errors.New("required BGP address family unavailable")

func ValidateBGPFamily(s BGPSession, family AddressFamily) error {
	if family == IPv4 && !s.IPv4 { return ErrBGPFamilyUnavailable }
	if family == IPv6 && !s.IPv6 { return ErrBGPFamilyUnavailable }
	return nil
}
