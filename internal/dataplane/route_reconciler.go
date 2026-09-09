package dataplane

import "context"

type RouteReconciler struct { Router RouterAdapter }

func (r RouteReconciler) Apply(ctx context.Context, desired []RouteIntent) error {
	if r.Router == nil { return ErrBackendUnavailable }
	if err := ValidateRoutes(desired); err != nil { return err }
	if err := r.Router.Health(ctx); err != nil { return err }
	return r.Router.ApplyRoutes(ctx, desired)
}

func (r RouteReconciler) Withdraw(ctx context.Context, routes []RouteIntent) error {
	if r.Router == nil { return ErrBackendUnavailable }
	if err := ValidateRoutes(routes); err != nil { return err }
	if err := r.Router.Health(ctx); err != nil { return err }
	return r.Router.WithdrawRoutes(ctx, routes)
}
