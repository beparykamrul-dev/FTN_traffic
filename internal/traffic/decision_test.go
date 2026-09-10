package traffic

import (
	"testing"
	"github.com/beparykamrul-dev/FTN_traffic/internal/latency"
)

func TestDecisionRejectsWhenNoEligiblePath(t *testing.T) {
	s := Selector{MinAvailability:.99, MaxLoss:1, MaxP95:100}
	if _, err := s.Decide([]Candidate{{Path:latency.Path{ID:"x", Authorized:false, Healthy:true}}}, ""); err != ErrNoEligiblePath { t.Fatalf("got %v", err) }
}

func TestDecisionMarksCurrentPath(t *testing.T) {
	s := Selector{Weights:latency.Weights{Latency:1}, MinAvailability:.99, MaxLoss:1, MaxP95:100}
	got, err := s.Decide([]Candidate{{Path:latency.Path{ID:"a",Authorized:true,Healthy:true,Measurement:latency.Measurement{Availability:1,RTTP95MS:10}}}}, "a")
	if err != nil || !got.CurrentKept || got.SelectedID != "a" { t.Fatalf("%+v %v", got, err) }
}
