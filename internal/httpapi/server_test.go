package httpapi

import (
	"net/http/httptest"
	"testing"

	ftndns "github.com/beparykamrul-dev/FTN_traffic/internal/dns"
	"github.com/beparykamrul-dev/FTN_traffic/internal/traffic"
)

func TestHandlerHealthAndReady(t *testing.T) {
	h := Handler(traffic.NewStore(), ftndns.NewPortal())
	for _, path := range []string{"/healthz", "/readyz", "/metrics", "/api/v1/dns/portal", "/api/v1/dns/engines", "/api/v1/dns/providers", "/api/v1/dns/nodes"} {
		r := httptest.NewRecorder()
		h.ServeHTTP(r, httptest.NewRequest("GET", path, nil))
		if r.Code != 200 { t.Fatalf("%s returned %d", path, r.Code) }
	}
}
