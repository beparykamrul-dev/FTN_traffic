package traffic

import (
	"testing"

	"github.com/beparykamrul-dev/FTN_traffic/internal/latency"
)

func TestSelectorRejectsUnauthorized(t *testing.T) {
	s := Selector{MinAvailability: .99, MaxLoss: 1, MaxP95: 100}
	got, ok := s.Select([]Candidate{{Path: latency.Path{ID: "private", Authorized: false, Healthy: true, Measurement: latency.Measurement{Availability: 1, LossPercent: 0, RTTP95MS: 10}}}}, "")
	if ok || got.Path.ID != "" { t.Fatal("expected no eligible path") }
}

func TestSelectorPrefersLowerScore(t *testing.T) {
	s := Selector{Weights: latency.Weights{Latency: 1}, MinAvailability: .99, MaxLoss: 1, MaxP95: 200}
	c := []Candidate{
		{Path: latency.Path{ID: "slow", Authorized: true, Healthy: true, Measurement: latency.Measurement{Availability: 1, RTTP95MS: 80}}},
		{Path: latency.Path{ID: "fast", Authorized: true, Healthy: true, Measurement: latency.Measurement{Availability: 1, RTTP95MS: 20}}},
	}
	got, ok := s.Select(c, "")
	if !ok || got.Path.ID != "fast" { t.Fatalf("got %q", got.Path.ID) }
}
