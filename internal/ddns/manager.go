package ddns

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	ErrApprovalRequired = errors.New("DDNS mutation requires explicit approval")
	ErrInvalidUpdate = errors.New("invalid DDNS update")
)

type Manager struct {
	mu sync.RWMutex
	records map[string]Record
}

func NewManager() *Manager { return &Manager{records: make(map[string]Record)} }

func Key(zone, name string) string { return strings.ToLower(strings.TrimSuffix(zone, ".")) + ":" + strings.ToLower(strings.TrimSuffix(name, ".")) }

func (m *Manager) Get(zone, name string) (Record, bool) {
	m.mu.RLock(); defer m.mu.RUnlock()
	r, ok := m.records[Key(zone, name)]
	return r, ok
}

func (m *Manager) Snapshot() []Record {
	m.mu.RLock(); defer m.mu.RUnlock()
	out := make([]Record, 0, len(m.records))
	for _, r := range m.records { out = append(out, r) }
	return out
}

func Validate(u Update) error {
	if strings.TrimSpace(u.Zone) == "" || strings.TrimSpace(u.Name) == "" || net.ParseIP(strings.TrimSpace(u.Address)) == nil {
		return ErrInvalidUpdate
	}
	if u.TTL == 0 || u.TTL > 86400 { return fmt.Errorf("%w: TTL must be 1..86400", ErrInvalidUpdate) }
	if strings.TrimSpace(u.ApprovalID) == "" { return ErrApprovalRequired }
	return nil
}

func (m *Manager) Apply(u Update, now time.Time) (Record, error) {
	if err := Validate(u); err != nil { return Record{}, err }
	if now.IsZero() { now = time.Now() }
	idBytes := sha256.Sum256([]byte(Key(u.Zone, u.Name) + "|" + u.Address))
	r := Record{ID: hex.EncodeToString(idBytes[:8]), Zone: strings.TrimSuffix(u.Zone, "."), Name: strings.TrimSuffix(u.Name, "."), Type: "A", Address: u.Address, TTL: u.TTL, Enabled: true, LastSeenAt: now}
	if net.ParseIP(u.Address).To4() == nil { r.Type = "AAAA" }
	m.mu.Lock(); m.records[Key(u.Zone, u.Name)] = r; m.mu.Unlock()
	return r, nil
}
