package httpapi

import (
	"encoding/json"
	"net/http"

	ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func Handler(store *traffic.Store, dnsPortal *ftndns.Portal) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) { write(w, http.StatusOK, map[string]any{"status": "ok"}) })
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, _ *http.Request) { write(w, http.StatusOK, map[string]any{"ready": true}) })
	mux.Handle("GET /metrics", MetricsHandler(store))
	mux.HandleFunc("GET /api/v1/dns/portal", func(w http.ResponseWriter, _ *http.Request) { write(w, http.StatusOK, dnsPortal.Snapshot()) })
	mux.HandleFunc("GET /api/v1/dns/engines", func(w http.ResponseWriter, _ *http.Request) { write(w, http.StatusOK, dnsPortal.Snapshot().Engines) })
	mux.HandleFunc("GET /api/v1/dns/providers", func(w http.ResponseWriter, _ *http.Request) { write(w, http.StatusOK, dnsPortal.Snapshot().Providers) })
	mux.HandleFunc("GET /api/v1/dns/nodes", func(w http.ResponseWriter, _ *http.Request) { write(w, http.StatusOK, dnsPortal.Snapshot().Nodes) })
	return mux
}

func write(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
