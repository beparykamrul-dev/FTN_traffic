package routing

import "context"

type Adapter interface {
	Name() string
	Health(ctx context.Context) error
	Advertise(ctx context.Context, prefixes []string) error
	Withdraw(ctx context.Context, prefixes []string) error
}

type Policy struct {
	LocalPref uint32 `json:"local_pref"`
	Communities []string `json:"communities"`
	Authorized bool `json:"authorized"`
}
