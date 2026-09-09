package latency

import "time"

type PathClass string

const (
	PathFTNFullMesh PathClass = "ftn_full_mesh"
	PathAuthorizedIX PathClass = "authorized_ix"
	PathAuthorizedTransit PathClass = "authorized_transit"
	PathAuthorizedGlobalEdge PathClass = "authorized_global_edge"
	PathAuthorizedOrigin PathClass = "authorized_origin_host"
)

type Measurement struct {
	RTTP50MS float64 `json:"rtt_p50_ms"`
	RTTP95MS float64 `json:"rtt_p95_ms"`
	JitterMS float64 `json:"jitter_ms"`
	LossPercent float64 `json:"packet_loss_percent"`
	Availability float64 `json:"availability"`
	At time.Time `json:"at"`
}

type Path struct {
	ID string `json:"id"`
	Class PathClass `json:"class"`
	Authorized bool `json:"authorized"`
	Healthy bool `json:"healthy"`
	Measurement Measurement `json:"measurement"`
}

type Weights struct { Latency, Loss, Jitter, Availability float64 }

func (p Path) Cost(w Weights) float64 {
	m := p.Measurement
	availabilityPenalty := (1 - m.Availability) * 1000
	return w.Latency*m.RTTP95MS + w.Loss*m.LossPercent + w.Jitter*m.JitterMS + w.Availability*availabilityPenalty
}

func Eligible(p Path, minAvailability, maxLoss, maxP95 float64) bool {
	return p.Authorized && p.Healthy && p.Measurement.Availability >= minAvailability && p.Measurement.LossPercent <= maxLoss && p.Measurement.RTTP95MS <= maxP95
}

func Select(paths []Path, current string, w Weights, minAvailability, maxLoss, maxP95, hysteresis float64) (Path, bool) {
	var best Path
	bestCost := 0.0
	found := false
	for _, p := range paths {
		if !Eligible(p, minAvailability, maxLoss, maxP95) { continue }
		c := p.Cost(w)
		if !found || c < bestCost { best, bestCost, found = p, c, true }
	}
	if !found { return Path{}, false }
	if current == "" || best.ID == current { return best, true }
	for _, p := range paths {
		if p.ID == current && Eligible(p, minAvailability, maxLoss, maxP95) {
			cc := p.Cost(w)
			if bestCost >= cc*(1-hysteresis/100) { return p, true }
		}
	}
	return best, true
}
