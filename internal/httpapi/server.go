package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func Handler(store *traffic.Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) { write(w, http.StatusOK, map[string]any{"status": "ok"}) })
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, _ *http.Request) { write(w, http.StatusOK, map[string]any{"ready": true}) })
	mux.Handle("GET /metrics", MetricsHandler(store))
	return mux
}

func write(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
