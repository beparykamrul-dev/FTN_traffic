package dataplane

func ValidateRouteCommunities(r RouteIntent) error {
	return ValidateCommunities(r.Community)
}

func ValidateRouteWithPolicy(r RouteIntent, policy RoutePolicy) error {
	if err := ValidateRoute(r); err != nil { return err }
	if err := ValidateRouteCommunities(r); err != nil { return err }
	if policy.RequireRPKI {
		return ErrRPKIRequired
	}
	return nil
}
