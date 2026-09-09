package dataplane

import "errors"

var ErrTransportUnauthorized = errors.New("transport peer unauthorized")
var ErrTransportAddressFamily = errors.New("transport peer has no address family")

func ValidatePeer(p Peer) error {
	if !p.Authorized { return ErrTransportUnauthorized }
	if p.ID == "" || p.Endpoint == "" { return errors.New("peer identity and endpoint are required") }
	if !p.IPv4 && !p.IPv6 { return ErrTransportAddressFamily }
	return nil
}
