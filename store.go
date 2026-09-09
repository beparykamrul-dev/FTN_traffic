package traffic

import "sync"

type Store struct {
	mu sync.RWMutex
	total Aggregate
}

func NewStore() *Store { return &Store{} }

func (s *Store) Add(e Event) Aggregate {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.total = Add(s.total, e)
	return s.total
}

func (s *Store) Snapshot() Aggregate {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.total
}
