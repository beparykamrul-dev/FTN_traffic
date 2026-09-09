package ddns

import (
	"errors"
	"testing"
	"time"
)

func TestApplyRequiresApproval(t *testing.T) {
	_, err := NewManager().Apply(Update{Zone:"example.com", Name:"home", Address:"203.0.113.10", TTL:60}, time.Now())
	if !errors.Is(err, ErrApprovalRequired) { t.Fatalf("expected approval error, got %v", err) }
}

func TestApplyAAndAAAA(t *testing.T) {
	m := NewManager()
	r, err := m.Apply(Update{Zone:"example.com", Name:"home", Address:"203.0.113.10", TTL:60, ApprovalID:"ap-1"}, time.Now())
	if err != nil || r.Type != "A" { t.Fatalf("A update failed: %+v %v", r, err) }
	r, err = m.Apply(Update{Zone:"example.com", Name:"v6", Address:"2001:db8::10", TTL:120, ApprovalID:"ap-2"}, time.Now())
	if err != nil || r.Type != "AAAA" { t.Fatalf("AAAA update failed: %+v %v", r, err) }
}
