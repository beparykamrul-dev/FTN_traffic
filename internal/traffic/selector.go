package traffic

import (
	"errors"
	"sort"

	"github.com/beparykamrul-dev/FTN_traffic/internal/latency"
)

var ErrNoEligiblePath = errors.New("no eligible traffic path")

type Candidate struct {
	Path latency.Path
	CapacityScore float64
	LocalityScore float64
}

type Selector struct {
	Weights latency.Weights
	MinAvailability float64
	MaxLoss float64
	MaxP95 float64
	Hysteresis float64
}

func (s Selector) Select(candidates []Candidate, current string) (Candidate, bool) {
	eligible := make([]Candidate, 0, len(candidates))
	for _, c := range candidates {
		if !latency.Eligible(c.Path, s.MinAvailability, s.MaxLoss, s.MaxP95) { continue }
		eligible = append(eligible, c)
	}
	if len(eligible) == 0 { return Candidate{}, false }
	sort.SliceStable(eligible, func(i, j int) bool {
		ci := s.score(eligible[i])
		cj := s.score(eligible[j])
		if ci == cj { return eligible[i].Path.ID < eligible[j].Path.ID }
		return ci < cj
	})
	best := eligible[0]
	if current != "" {
		for _, c := range eligible {
			if c.Path.ID != current { continue }
			if s.score(best) >= s.score(c)*(1-s.Hysteresis/100) { return c, true }
		}
	}
	return best, true
}

func (s Selector) score(c Candidate) float64 {
	return c.Path.Cost(s.Weights) - c.CapacityScore - c.LocalityScore
}
