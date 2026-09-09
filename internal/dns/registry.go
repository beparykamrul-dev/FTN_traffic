package dns

import (
	"fmt"
	"sync"
)

type Registry struct {
	mu sync.RWMutex
	items map[string]Provider
}

func NewRegistry() *Registry { return &Registry{items: make(map[string]Provider)} }

func (r *Registry) Register(p Provider) error {
	if p == nil || p.Name() == "" { return fmt.Errorf("invalid DNS provider") }
	r.mu.Lock(); defer r.mu.Unlock()
	if _, exists := r.items[p.Name()]; exists { return fmt.Errorf("DNS provider %q already registered", p.Name()) }
	r.items[p.Name()] = p
	return nil
}

func (r *Registry) Get(name string) (Provider, bool) { r.mu.RLock(); defer r.mu.RUnlock(); p, ok := r.items[name]; return p, ok }
