package httpapi

import (
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func TestMetricsHandler(t *testing.T) {
	s := traffic.NewStore()
	s.Add(traffic.Event{Provider: "cdn", Plane: "edge", POP: "ctg", Bytes: 42, Packets: 3})
	r := httptest.NewRecorder()
	MetricsHandler(s)(r, httptest.NewRequest("GET", "/metrics", nil))
	body := r.Body.String()
	for _, want := range []string{"ftn_traffic_events_total 1", "ftn_traffic_bytes_total 42", "ftn_traffic_packets_total 3"} {
		if !strings.Contains(body, want) { t.Fatalf("missing %q in %q", want, body) }
	}
}
