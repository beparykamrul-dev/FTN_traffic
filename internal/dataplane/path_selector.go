package dataplane

import "context"

type PathHealth struct {
	ID string `json:"id"`
	Healthy bool `json:"healthy"`
	LatencyMS float64 `json:"latency_ms"`
	LossPercent float64 `json:"packet_loss_percent"`
	Availability float64 `json:"availability"`
}

type PathSelector struct { MinAvailability float64; MaxLossPercent float64 }

func (s PathSelector) Select(paths []PathHealth) (PathHealth, bool) {
	var best PathHealth
	found := false
	for _, p := range paths {
		if !p.Healthy || p.Availability < s.MinAvailability || p.LossPercent > s.MaxLossPercent { continue }
		if !found || p.LatencyMS < best.LatencyMS { best, found = p, true }
	}
	return best, found
}

func ProbeHealth(ctx context.Context, p PathHealth) bool {
	select { case <-ctx.Done(): return false; default: return p.Healthy }
}
