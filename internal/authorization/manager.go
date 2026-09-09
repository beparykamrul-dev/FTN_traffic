package authorization

import (
	"errors"
	"sync"
	"time"
)

var ErrNotConnected = errors.New("provider authorization is not active")

type Manager struct {
	mu    sync.RWMutex
	items map[string]ProviderAuth
}

func NewManager() *Manager { return &Manager{items: make(map[string]ProviderAuth)} }

func (m *Manager) Upsert(a ProviderAuth) {
	m.mu.Lock(); defer m.mu.Unlock()
	m.items[a.Provider] = a
}

func (m *Manager) Get(provider string) (ProviderAuth, bool) {
	m.mu.RLock(); defer m.mu.RUnlock()
	a, ok := m.items[provider]
	return a, ok
}

func (m *Manager) Authorized(provider string, now time.Time) error {
	a, ok := m.Get(provider)
	if !ok || !a.Active(now) { return ErrNotConnected }
	return nil
}

func (m *Manager) Snapshot() []ProviderAuth {
	m.mu.RLock(); defer m.mu.RUnlock()
	out := make([]ProviderAuth, 0, len(m.items))
	for _, a := range m.items { out = append(out, a) }
	return out
}
