package dataplane

import (
	"errors"
	"net/netip"
	"strings"
)

var (
	ErrInvalidPrefix = errors.New("invalid route prefix")
	ErrInvalidNextHop = errors.New("invalid next-hop")
	ErrRouteUnauthorized = errors.New("route intent unauthorized")
)

func ValidateRoute(r RouteIntent) error {
	if !r.Authorized { return ErrRouteUnauthorized }
	if _, err := netip.ParsePrefix(r.Prefix); err != nil { return ErrInvalidPrefix }
	if r.NextHop != "" {
		if _, err := netip.ParseAddr(r.NextHop); err != nil { return ErrInvalidNextHop }
	}
	if r.Family != IPv4 && r.Family != IPv6 { return errors.New("unsupported address family") }
	if strings.ContainsAny(r.Prefix, "\n\r") || strings.ContainsAny(r.NextHop, "\n\r") { return errors.New("route contains control characters") }
	return nil
}

func ValidateRoutes(routes []RouteIntent) error {
	for _, r := range routes { if err := ValidateRoute(r); err != nil { return err } }
	return nil
}
