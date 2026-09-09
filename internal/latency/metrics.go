package latency

import (
 "fmt"
 "net/http"
 "sync"
)

type Metrics struct { mu sync.RWMutex; selected string; paths map[string]Measurement }
func NewMetrics()*Metrics{return &Metrics{paths:map[string]Measurement{}}}
func(m *Metrics) Set(id string, v Measurement){m.mu.Lock();defer m.mu.Unlock();m.paths[id]=v}
func(m *Metrics) SetSelected(id string){m.mu.Lock();defer m.mu.Unlock();m.selected=id}
func(m *Metrics) Handler(w http.ResponseWriter,_ *http.Request){m.mu.RLock();defer m.mu.RUnlock();w.Header().Set("Content-Type","text/plain; version=0.0.4");if m.selected!=""{fmt.Fprintf(w,"ftn_latency_selected{path=%q} 1\n",m.selected)};for id,v:=range m.paths{fmt.Fprintf(w,"ftn_path_rtt_p95_ms{path=%q} %g\nftn_path_rtt_p50_ms{path=%q} %g\nftn_path_jitter_ms{path=%q} %g\nftn_path_loss_percent{path=%q} %g\nftn_path_availability{path=%q} %g\n",id,v.RTTP95MS,id,v.RTTP50MS,id,v.JitterMS,id,v.LossPercent,id,v.Availability)}}
