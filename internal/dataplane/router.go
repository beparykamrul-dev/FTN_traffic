package dataplane

import "context"

type RouterAdapter interface {
	Name() string
	Kind() RouterKind
	Health(context.Context) error
	ApplyRoutes(context.Context, []RouteIntent) error
	WithdrawRoutes(context.Context, []RouteIntent) error
	Snapshot(context.Context) (RouterStatus, error)
}

// Mutating methods are contracts only; concrete adapters must enforce
// authorization and approval before touching a production router.
