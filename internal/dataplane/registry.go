package dataplane

import (
	"context"
	"errors"
	"sync"
)

var ErrDuplicateBackend = errors.New("dataplane backend already registered")

// Registry keeps backend implementations local to the FTN runtime.
type Registry[T interface{ Name() string; Health(context.Context) error }] struct {
	mu sync.RWMutex
	items map[string]T
}

func NewRegistry[T interface{ Name() string; Health(context.Context) error }]() *Registry[T] {
	return &Registry[T]{items: make(map[string]T)}
}

func (r *Registry[T]) Register(v T) error {
	if v.Name() == "" { return errors.New("backend name is required") }
	r.mu.Lock(); defer r.mu.Unlock()
	if _, ok := r.items[v.Name()]; ok { return ErrDuplicateBackend }
	r.items[v.Name()] = v
	return nil
}

func (r *Registry[T]) Get(name string) (T, bool) {
	r.mu.RLock(); defer r.mu.RUnlock()
	v, ok := r.items[name]
	return v, ok
}

func (r *Registry[T]) Names() []string {
	r.mu.RLock(); defer r.mu.RUnlock()
	out := make([]string, 0, len(r.items))
	for name := range r.items { out = append(out, name) }
	return out
}
