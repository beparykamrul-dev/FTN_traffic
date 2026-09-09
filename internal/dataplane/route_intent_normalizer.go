package dataplane

import (
	"fmt"
	"net/netip"
	"strings"
)

func NormalizeRouteIntent(r RouteIntent) (RouteIntent, error) {
	if err := ValidateRoute(r); err != nil { return RouteIntent{}, err }
	p, err := netip.ParsePrefix(r.Prefix)
	if err != nil { return RouteIntent{}, err }
	r.Prefix = p.Masked().String()
	if r.NextHop != "" {
		n, err := netip.ParseAddr(r.NextHop); if err != nil { return RouteIntent{}, err }
		r.NextHop = n.String()
	}
	for _, c := range r.Community {
		if strings.TrimSpace(c) == "" { return RouteIntent{}, fmt.Errorf("empty route community") }
	}
	return r, nil
}

func NormalizeRoutes(routes []RouteIntent) ([]RouteIntent, error) {
	if len(routes) == 0 { return nil, ErrEmptyRouteBatch }
	out := make([]RouteIntent, 0, len(routes))
	seen := make(map[string]struct{}, len(routes))
	for _, r := range routes {
		n, err := NormalizeRouteIntent(r); if err != nil { return nil, err }
		key := string(n.Family)+"|"+n.Prefix+"|"+n.NextHop
		if _, ok := seen[key]; ok { return nil, ErrDuplicateRoute }
		seen[key] = struct{}{}
		out = append(out, n)
	}
	return out, nil
}
