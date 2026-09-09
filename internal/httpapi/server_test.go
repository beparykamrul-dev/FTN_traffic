package httpapi

import (
	"net/http/httptest"
	"testing"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func TestHandlerHealthAndReady(t *testing.T) {
	h := Handler(traffic.NewStore())
	for _, path := range []string{"/healthz", "/readyz", "/metrics"} {
		r := httptest.NewRecorder()
		h.ServeHTTP(r, httptest.NewRequest("GET", path, nil))
		if r.Code != 200 { t.Fatalf("%s returned %d", path, r.Code) }
	}
}
