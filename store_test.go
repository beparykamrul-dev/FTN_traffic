package traffic

import (
	"sync"
	"testing"
)

func TestStoreConcurrentAdd(t *testing.T) {
	s := NewStore()
	const n = 100
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			s.Add(Event{Bytes: 10, Packets: 2})
		}()
	}
	wg.Wait()
	got := s.Snapshot()
	if got.Events != n || got.Bytes != n*10 || got.Packets != n*2 {
		t.Fatalf("unexpected snapshot: %+v", got)
	}
}
