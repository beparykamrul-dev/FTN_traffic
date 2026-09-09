package dataplane

import "errors"

var (
	ErrRPKIRequired     = errors.New("route requires valid RPKI")
	ErrMaxPrefixExceeded = errors.New("route exceeds max-prefix limit")
	ErrEmptyRouteBatch  = errors.New("route batch is empty")
)

type RoutePolicy struct {
	RequireRPKI  bool   `json:"require_rpki"`
	MaxPrefixes  uint64 `json:"max_prefixes"`
}

type RouteBatch struct {
	Routes    []RouteIntent `json:"routes"`
	RPKIValid bool          `json:"rpki_valid"`
}

func (p RoutePolicy) Validate(b RouteBatch) error {
	if len(b.Routes) == 0 {
		return ErrEmptyRouteBatch
	}
	if p.RequireRPKI && !b.RPKIValid {
		return ErrRPKIRequired
	}
	if p.MaxPrefixes == 0 || uint64(len(b.Routes)) > p.MaxPrefixes {
		return ErrMaxPrefixExceeded
	}
	if err := ValidateRoutes(b.Routes); err != nil {
		return err
	}
	for _, r := range b.Routes {
		if err := ValidateRouteCommunities(r); err != nil { return err }
	}
	return nil
}
