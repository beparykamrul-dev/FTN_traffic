package dataplane

import "fmt"

func ValidateRouteSetForMutation(routes []RouteIntent) error {
	if len(routes) == 0 { return ErrEmptyRouteBatch }
	for _, r := range routes {
		if !r.Authorized { return ErrRouteUnauthorized }
		if r.LocalPref > 4_294_967_000 { return fmt.Errorf("local preference out of safe range") }
		if r.MED > 4_294_967_000 { return fmt.Errorf("MED out of safe range") }
	}
	_, err := NormalizeRoutes(routes)
	return err
}
