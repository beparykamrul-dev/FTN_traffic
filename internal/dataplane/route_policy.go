package dataplane

import "errors"

var ErrRPKIRequired = errors.New("route requires valid RPKI")
var ErrMaxPrefixExceeded = errors.New("route exceeds max-prefix limit")

type RoutePolicy struct { RequireRPKI bool `json:"require_rpki"`; MaxPrefixes uint64 `json:"max_prefixes"` }

type RouteBatch struct { Routes []RouteIntent `json:"routes"`; RPKIValid bool `json:"rpki_valid"` }

func (p RoutePolicy) Validate(b RouteBatch) error {
	if p.RequireRPKI && !b.RPKIValid { return ErrRPKIRequired }
	if p.MaxPrefixes == 0 || uint64(len(b.Routes)) > p.MaxPrefixes { return ErrMaxPrefixExceeded }
	return ValidateRoutes(b.Routes)
}
