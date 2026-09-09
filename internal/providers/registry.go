package providers

import "sync"

type Registry struct {
	mu sync.RWMutex
	items map[string]Provider
}

func NewRegistry(items ...Provider) *Registry {
	r := &Registry{items: make(map[string]Provider, len(items))}
	for _, p := range items { r.items[p.Name] = p }
	return r
}

func (r *Registry) Get(name string) (Provider, bool) {
	r.mu.RLock(); defer r.mu.RUnlock()
	p, ok := r.items[name]
	return p, ok
}

func (r *Registry) Set(p Provider) {
	r.mu.Lock(); defer r.mu.Unlock()
	r.items[p.Name] = p
}
