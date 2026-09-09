package edge

import "context"

type Adapter interface {
	Name() string
	Health(ctx context.Context) error
	Purge(ctx context.Context, paths []string) error
}

type Target struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Authorized bool `json:"authorized"`
}
