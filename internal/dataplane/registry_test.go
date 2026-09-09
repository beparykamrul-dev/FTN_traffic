package dataplane

import (
	"context"
	"testing"
)

type registryBackend struct { name string }
func (r registryBackend) Name() string { return r.name }
func (r registryBackend) Health(context.Context) error { return nil }

func TestRegistryRejectsDuplicateAndMissingNames(t *testing.T) {
	r := NewRegistry[registryBackend]()
	if err := r.Register(registryBackend{}); err == nil { t.Fatal("expected missing-name error") }
	if err := r.Register(registryBackend{name:"r1"}); err != nil { t.Fatal(err) }
	if err := r.Register(registryBackend{name:"r1"}); err != ErrDuplicateBackend { t.Fatalf("got %v", err) }
	if _, ok := r.Get("r1"); !ok { t.Fatal("backend not found") }
}
