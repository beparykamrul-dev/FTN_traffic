package dataplane

func ValidateRouteCommunities(r RouteIntent) error {
	return ValidateCommunities(r.Community)
}

func ValidateRouteWithRPKI(r RouteIntent, rpkiValid bool) error {
	if err := ValidateRoute(r); err != nil { return err }
	if err := ValidateRouteCommunities(r); err != nil { return err }
	if !rpkiValid { return ErrRPKIRequired }
	return nil
}
