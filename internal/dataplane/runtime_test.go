package dataplane

import (
	"context"
	"testing"
)

type runtimeBackend struct { name string; err error }
func (r runtimeBackend) Name() string { return r.name }
func (r runtimeBackend) Health(context.Context) error { return r.err }

func TestRuntimeHealthUnknownBackend(t *testing.T) {
	r := NewLocalRuntime()
	if err := Health(context.Background(), r.Routers, "missing"); err != ErrBackendUnavailable { t.Fatalf("got %v", err) }
}

func TestRuntimeHealthPropagatesBackendError(t *testing.T) {
	r := NewLocalRuntime()
	b := runtimeBackend{name:"r1", err:ErrBGPDown}
	if err := r.Routers.Register(b); err != nil { t.Fatal(err) }
	if err := Health(context.Background(), r.Routers, "r1"); err != ErrBGPDown { t.Fatalf("got %v", err) }
}
