package upstream

import "errors"

var (
	ErrUnauthorized = errors.New("upstream session is not authorized")
	ErrInvalidASN   = errors.New("local and remote ASN are required")
	ErrNoAddressFamily = errors.New("at least one address family is required")
)

type Policy struct {
	LocalASN          uint32
	RequireRPKI       bool
	RequireMaxPrefix  bool
	RequireBFD        bool
	AllowDefaultRoute bool
	AllowFullTable    bool
}

func ValidateSession(s Session, p Policy) error {
	if !s.Authorized {
		return ErrUnauthorized
	}
	if s.LocalASN == 0 || s.RemoteASN == 0 {
		return ErrInvalidASN
	}
	if !s.IPv4 && !s.IPv6 {
		return ErrNoAddressFamily
	}
	if p.LocalASN != 0 && s.LocalASN != p.LocalASN {
		return ErrInvalidASN
	}
	if p.RequireMaxPrefix && s.MaxPrefixes == 0 {
		return errors.New("max-prefix policy is required")
	}
	return nil
}
