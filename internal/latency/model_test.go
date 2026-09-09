package latency

import (
	"testing"
	"time"
)

func TestSelectPrefersHealthyLowLatency(t *testing.T) {
	paths := []Path{
		{ID:"mesh", Class:PathFTNFullMesh, Authorized:true, Healthy:true, Measurement:Measurement{RTTP95MS:12, LossPercent:.2, JitterMS:1, Availability:.9999}},
		{ID:"transit", Class:PathAuthorizedTransit, Authorized:true, Healthy:true, Measurement:Measurement{RTTP95MS:40, LossPercent:.1, JitterMS:2, Availability:.9999}},
	}
	got, ok := Select(paths, "", Weights{Latency:.5, Loss:.2, Jitter:.1, Availability:.2}, .999, 2, 1000, 10)
	if !ok || got.ID != "mesh" { t.Fatalf("selected %+v", got) }
}

func TestSelectHysteresisKeepsCurrent(t *testing.T) {
	now := time.Now()
	paths := []Path{
		{ID:"a", Authorized:true, Healthy:true, Measurement:Measurement{RTTP95MS:20, Availability:1, At:now}},
		{ID:"b", Authorized:true, Healthy:true, Measurement:Measurement{RTTP95MS:19, Availability:1, At:now}},
	}
	got, ok := Select(paths, "a", Weights{Latency:1}, .999, 2, 1000, 10)
	if !ok || got.ID != "a" { t.Fatalf("selected %+v", got) }
}

func TestUnauthorizedNeverSelected(t *testing.T) {
	p := Path{ID:"bad", Authorized:false, Healthy:true, Measurement:Measurement{RTTP95MS:1, Availability:1}}
	if _, ok := Select([]Path{p}, "", Weights{Latency:1}, .999, 2, 1000, 10); ok { t.Fatal("unauthorized path selected") }
}
