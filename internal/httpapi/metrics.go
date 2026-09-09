package httpapi

import (
	"fmt"
	"net/http"

	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func MetricsHandler(store *traffic.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		s := store.Snapshot()
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "ftn_traffic_events_total %d\nftn_traffic_bytes_total %d\nftn_traffic_packets_total %d\n", s.Events, s.Bytes, s.Packets)
	}
}
