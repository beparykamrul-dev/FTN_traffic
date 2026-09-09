package dataplane

import (
	"context"
	"testing"
)

type testTelemetry struct { healthy bool; published int }
func (t *testTelemetry) Name() string { return "test" }
func (t *testTelemetry) Health(context.Context) error { if !t.healthy { return ErrBackendUnavailable }; return nil }
func (t *testTelemetry) Publish(context.Context, Metric) error { t.published++; return nil }

func TestTelemetrySinkValidatesAndHealthChecks(t *testing.T) {
	b := &testTelemetry{healthy:true}
	s := TelemetrySink{Backend:b}
	if err := s.Publish(context.Background(), Metric{Name:"ftn_test", Value:1}); err != nil { t.Fatal(err) }
	if b.published != 1 { t.Fatalf("published=%d", b.published) }
	if err := s.Publish(context.Background(), Metric{Name:"", Value:1}); err != ErrInvalidMetric { t.Fatalf("got %v", err) }
	b.healthy=false
	if err := s.Publish(context.Background(), Metric{Name:"ftn_test", Value:2}); err != ErrBackendUnavailable { t.Fatalf("got %v", err) }
}
