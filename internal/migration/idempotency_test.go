package migration

import (
	"sync"
	"testing"
)

func TestIdempotency(t *testing.T) {
	i := NewIdempotency()
	if !i.Claim("x", "fp-1") { t.Fatal("first claim failed") }
	if i.Claim("x", "fp-1") { t.Fatal("duplicate claim accepted") }
	if i.Claim("", "fp-1") { t.Fatal("empty key accepted") }
	if i.Claim("y", "") { t.Fatal("empty fingerprint accepted") }
	if !i.Matches("x", "fp-1") { t.Fatal("fingerprint should match") }
	if i.Matches("x", "fp-2") { t.Fatal("different fingerprint matched") }
}

func TestIdempotencyFailedStateCanRetry(t *testing.T) {
	i := NewIdempotency()
	if !i.Claim("retry", "fp") { t.Fatal("initial claim failed") }
	i.SetState("retry", StateFailed)
	if !i.Claim("retry", "fp") { t.Fatal("failed request was not released for retry") }
}

func TestIdempotencyCompletedStateRemainsSingleUse(t *testing.T) {
	i := NewIdempotency()
	if !i.Claim("done", "fp") { t.Fatal("initial claim failed") }
	i.SetState("done", StateCommit)
	if i.Claim("done", "fp") { t.Fatal("committed request was accepted twice") }
}

func TestIdempotencyConcurrentClaimIsSingleWinner(t *testing.T) {
	i := NewIdempotency()
	const n = 32
	var wg sync.WaitGroup
	results := make(chan bool, n)
	wg.Add(n)
	for j := 0; j < n; j++ {
		go func() {
			defer wg.Done()
			results <- i.Claim("concurrent", "same-fingerprint")
		}()
	}
	wg.Wait()
	close(results)
	wins := 0
	for ok := range results { if ok { wins++ } }
	if wins != 1 { t.Fatalf("expected exactly one winner, got %d", wins) }
}
