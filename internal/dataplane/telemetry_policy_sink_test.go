package dataplane

import (
	"context"
	"testing"
)

func TestPolicyTelemetrySinkRedactsAndPublishes(t *testing.T) {
	b := &testTelemetry{healthy: true}
	s := PolicyTelemetrySink{Backend: b, Policy: TelemetryPolicy{MaxLabels: 4}}
	m := Metric{Name: "ftn_path_latency_ms", Value: 12, Labels: map[string]string{"pop": "ctg", "token": "secret"}}
	if err := s.Publish(context.Background(), m); err != nil { t.Fatal(err) }
	if b.published != 1 { t.Fatalf("published=%d", b.published) }
	if _, ok := b.last.Labels["token"]; ok { t.Fatal("secret label was forwarded") }
}

func TestPolicyTelemetrySinkRejectsExcessLabels(t *testing.T) {
	b := &testTelemetry{healthy: true}
	s := PolicyTelemetrySink{Backend: b, Policy: TelemetryPolicy{MaxLabels: 1}}
	m := Metric{Name: "ftn_path_latency_ms", Value: 12, Labels: map[string]string{"pop": "ctg", "path": "p1"}}
	if err := s.Publish(context.Background(), m); err != ErrTelemetryPolicy { t.Fatalf("got %v", err) }
	if b.published != 0 { t.Fatal("rejected metric was published") }
}
