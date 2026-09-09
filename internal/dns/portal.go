package dns

import (
	"sort"
	"sync"
	"time"
)

type EngineStatus struct {
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Enabled   bool      `json:"enabled"`
	Healthy   bool      `json:"healthy"`
	CheckedAt time.Time `json:"checked_at"`
}

type ProviderStatus struct {
	Name        string    `json:"name"`
	Authorized  bool      `json:"authorized"`
	Healthy     bool      `json:"healthy"`
	LastChecked time.Time `json:"last_checked"`
}

type NodeStatus struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Anycast   bool      `json:"anycast"`
	Healthy   bool      `json:"healthy"`
	LatencyMS uint64    `json:"latency_ms"`
	CheckedAt time.Time `json:"checked_at"`
}

type PortalSnapshot struct {
	Engines   []EngineStatus   `json:"engines"`
	Providers []ProviderStatus `json:"providers"`
	Nodes     []NodeStatus     `json:"nodes"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type Portal struct {
	mu       sync.RWMutex
	engines  map[string]EngineStatus
	providers map[string]ProviderStatus
	nodes    map[string]NodeStatus
}

func NewPortal() *Portal {
	return &Portal{
		engines: make(map[string]EngineStatus),
		providers: make(map[string]ProviderStatus),
		nodes: make(map[string]NodeStatus),
	}
}

func (p *Portal) UpsertEngine(v EngineStatus) { p.mu.Lock(); defer p.mu.Unlock(); p.engines[v.Name] = v }
func (p *Portal) UpsertProvider(v ProviderStatus) { p.mu.Lock(); defer p.mu.Unlock(); p.providers[v.Name] = v }
func (p *Portal) UpsertNode(v NodeStatus) { p.mu.Lock(); defer p.mu.Unlock(); p.nodes[v.Name] = v }

func (p *Portal) Snapshot() PortalSnapshot {
	p.mu.RLock(); defer p.mu.RUnlock()
	s := PortalSnapshot{UpdatedAt: time.Now()}
	for _, v := range p.engines { s.Engines = append(s.Engines, v) }
	for _, v := range p.providers { s.Providers = append(s.Providers, v) }
	for _, v := range p.nodes { s.Nodes = append(s.Nodes, v) }
	sort.Slice(s.Engines, func(i,j int) bool { return s.Engines[i].Name < s.Engines[j].Name })
	sort.Slice(s.Providers, func(i,j int) bool { return s.Providers[i].Name < s.Providers[j].Name })
	sort.Slice(s.Nodes, func(i,j int) bool { return s.Nodes[i].Name < s.Nodes[j].Name })
	return s
}
