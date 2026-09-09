package latency

import("net/http/httptest";"strings";"testing")
func TestMetricsHandler(t *testing.T){m:=NewMetrics();m.Set("mesh",Measurement{RTTP95MS:12,Availability:.9999});m.SetSelected("mesh");r:=httptest.NewRecorder();m.Handler(r,nil);s:=r.Body.String();if !strings.Contains(s,"ftn_path_rtt_p95_ms")||!strings.Contains(s,"mesh"){t.Fatalf("metrics=%s",s)}}
