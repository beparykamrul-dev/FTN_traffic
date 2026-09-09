package dataplane

import "context"

type Reconciler struct { router RouterAdapter }

func NewReconciler(router RouterAdapter) *Reconciler { return &Reconciler{router: router} }

// Reconcile validates the desired route set before delegating to the selected
// local router backend. The router adapter remains responsible for its own
// provider/approval checks before mutation.
func (r *Reconciler) Reconcile(ctx context.Context, desired []RouteIntent) error {
	if r.router == nil { return ErrBackendUnavailable }
	if err := ValidateRoutes(desired); err != nil { return err }
	if err := r.router.Health(ctx); err != nil { return err }
	return r.router.ApplyRoutes(ctx, desired)
}

func (r *Reconciler) Withdraw(ctx context.Context, routes []RouteIntent) error {
	if r.router == nil { return ErrBackendUnavailable }
	if err := ValidateRoutes(routes); err != nil { return err }
	if err := r.router.Health(ctx); err != nil { return err }
	return r.router.WithdrawRoutes(ctx, routes)
}
