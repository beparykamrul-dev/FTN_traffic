package dataplane

import (
	"context"
	"errors"
)

var ErrBackendNameMismatch = errors.New("backend name does not match requested backend")

type LocalRuntime struct {
	Routers *Registry[RouterAdapter]
	Edges *Registry[EdgeBackend]
	DNS *Registry[DNSBackend]
	Firewalls *Registry[FirewallBackend]
	Telemetry *Registry[TelemetryBackend]
	Transports *Registry[TransportBackend]
}

func NewLocalRuntime() *LocalRuntime {
	return &LocalRuntime{
		Routers: NewRegistry[RouterAdapter](), Edges: NewRegistry[EdgeBackend](), DNS: NewRegistry[DNSBackend](),
		Firewalls: NewRegistry[FirewallBackend](), Telemetry: NewRegistry[TelemetryBackend](), Transports: NewRegistry[TransportBackend](),
	}
}

func Health[T interface{ Name() string; Health(context.Context) error }](ctx context.Context, r *Registry[T], name string) error {
	v, ok := r.Get(name); if !ok { return ErrBackendUnavailable }
	return v.Health(ctx)
}
