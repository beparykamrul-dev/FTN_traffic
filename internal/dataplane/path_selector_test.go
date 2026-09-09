package dataplane

import "testing"

func TestPathSelector(t *testing.T) {
	s := PathSelector{MinAvailability:0.999, MaxLossPercent:2}
	p, ok := s.Select([]PathHealth{{ID:"bad",Healthy:true,LatencyMS:5,LossPercent:3,Availability:1},{ID:"good",Healthy:true,LatencyMS:10,LossPercent:0.2,Availability:1}})
	if !ok || p.ID != "good" { t.Fatalf("unexpected selection: %+v %v", p, ok) }
}
