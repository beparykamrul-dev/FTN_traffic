package migration

import (
	"fmt"
	"net/http"
	"sync"
)

type Metrics struct { mu sync.RWMutex; states map[State]uint64; last string }
func NewMetrics()*Metrics{return &Metrics{states:map[State]uint64{}}}
func(m *Metrics) Observe(s State){m.mu.Lock();defer m.mu.Unlock();m.states[s]++}
func(m *Metrics) SetLast(id string){m.mu.Lock();defer m.mu.Unlock();m.last=id}
func(m *Metrics) Handler(w http.ResponseWriter,_ *http.Request){m.mu.RLock();defer m.mu.RUnlock();w.Header().Set("Content-Type","text/plain; version=0.0.4");for s,n:=range m.states{fmt.Fprintf(w,"ftn_migration_state_total{state=%q} %d\n",s,n)};if m.last!=""{fmt.Fprintf(w,"ftn_migration_last{migration=%q} 1\n",m.last)}}
