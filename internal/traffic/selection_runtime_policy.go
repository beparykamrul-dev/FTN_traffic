package traffic

import "errors"

var ErrRuntimeSelectionPolicy = errors.New("runtime traffic selection policy rejected")

type RuntimeSelectionPolicy struct {
	RequireAuthorization bool
	RequireHealthy bool
	RequireIPv6Ready bool
	AllowProviderFailover bool
}

func (p RuntimeSelectionPolicy) Validate(c Candidate) error {
	if p.RequireAuthorization && !c.Path.Authorized { return ErrRuntimeSelectionPolicy }
	if p.RequireHealthy && !c.Path.Healthy { return ErrRuntimeSelectionPolicy }
	if !p.AllowProviderFailover && c.Path.Class == "" { return ErrRuntimeSelectionPolicy }
	return nil
}
