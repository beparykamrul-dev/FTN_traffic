package ingest

import (
	"testing"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func TestNormalize(t *testing.T) {
	e, err := Normalize(traffic.Event{Provider: " cdn ", Plane: " edge ", POP: " ctg "})
	if err != nil || e.Provider != "cdn" || e.Plane != "edge" || e.POP != "ctg" { t.Fatalf("unexpected normalized event: %+v, %v", e, err) }
	if _, err := Normalize(traffic.Event{Provider: "cdn"}); err != ErrInvalidEvent { t.Fatalf("expected invalid event, got %v", err) }
}
