package traffic

import "github.com/beparykamrul-dev/FTN_traffic/internal/latency"

type Decision struct {
	SelectedID string
	SelectedClass latency.PathClass
	Score float64
	CurrentKept bool
	EligibleCount int
}

func (s Selector) Decide(candidates []Candidate, current string) (Decision, error) {
	chosen, ok := s.Select(candidates, current)
	if !ok { return Decision{}, ErrNoEligiblePath }
	kept := current != "" && chosen.Path.ID == current
	return Decision{SelectedID: chosen.Path.ID, SelectedClass: chosen.Path.Class, Score: s.score(chosen), CurrentKept: kept, EligibleCount: 1}, nil
}
