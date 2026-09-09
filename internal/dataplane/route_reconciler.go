package dataplane

import "context"

type RouteReconciler struct {
	Router     RouterAdapter
	Authorized bool
	Approved   bool
	Policy     *RoutePolicy
	RPKIValid  bool
}

func (r RouteReconciler) check(ctx context.Context, routes []RouteIntent) error {
	if r.Router == nil {
		return ErrBackendUnavailable
	}
	if !r.Authorized {
		return ErrUnauthorized
	}
	if !r.Approved {
		return ErrApprovalRequired
	}
	if len(routes) == 0 {
		return ErrEmptyRouteBatch
	}
	if r.Policy != nil {
		if err := r.Policy.Validate(RouteBatch{Routes: routes, RPKIValid: r.RPKIValid}); err != nil {
			return err
		}
	} else if err := ValidateRoutes(routes); err != nil {
		return err
	}
	return r.Router.Health(ctx)
}

func (r RouteReconciler) Apply(ctx context.Context, desired []RouteIntent) error {
	if err := r.check(ctx, desired); err != nil {
		return err
	}
	return r.Router.ApplyRoutes(ctx, desired)
}

func (r RouteReconciler) Withdraw(ctx context.Context, routes []RouteIntent) error {
	if err := r.check(ctx, routes); err != nil {
		return err
	}
	return r.Router.WithdrawRoutes(ctx, routes)
}
