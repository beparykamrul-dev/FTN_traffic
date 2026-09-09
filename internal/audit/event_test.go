package audit

import (
	"testing"
	"time"
)

func TestEventFields(t *testing.T) {
	now := time.Now()
	e := Event{Actor: "system", Action: "provider.refresh", Resource: "cdn", Approved: true, At: now}
	if e.Actor != "system" || !e.Approved || !e.At.Equal(now) { t.Fatalf("unexpected audit event: %+v", e) }
}
