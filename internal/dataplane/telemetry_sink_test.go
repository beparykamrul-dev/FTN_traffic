package dataplane

import (
	"context"
	"testing"
)

type testTelemetry struct { healthy bool; published int; last Metric }
func (t *testTelemetry) Name() string { return "test" }
func (t *testTelemetry) Health(context.Context) error { if !t.healthy { return ErrBackendUnavailable }; return nil }
func (t *testTelemetry) Publish(_ context.Context, m Metric) error { t.published++; t.last = m; return nil }

func TestTelemetrySinkValidatesAndHealthChecks(t *testing.T) {
	b := &testTelemetry{healthy:true}
	s := TelemetrySink{Backend:b}
	if err := s.Publish(context.Background(), Metric{Name:"ftn_test", Value:1}); err != nil { t.Fatal(err) }
	if b.published != 1 { t.Fatalf("published=%d", b.published) }
	if err := s.Publish(context.Background(), Metric{Name:"", Value:1}); err != ErrInvalidMetric { t.Fatalf("got %v", err) }
	b.healthy=false
	if err := s.Publish(context.Background(), Metric{Name:"ftn_test", Value:2}); err != ErrBackendUnavailable { t.Fatalf("got %v", err) }
}

func TestTelemetrySinkRejectsControlCharacters(t *testing.T) {
	b := &testTelemetry{healthy:true}
	s := TelemetrySink{Backend:b}
	if err := s.Publish(context.Background(), Metric{Name:"ftn_test", Value:1, Labels:map[string]string{"site":"pop\n1"}}); err != ErrInvalidMetric { t.Fatalf("got %v", err) }
	if b.published != 0 { t.Fatalf("invalid metric was published: %d", b.published) }
}

func TestTelemetrySinkRejectsNaNAndOversizedLabels(t *testing.T) {
	b := &testTelemetry{healthy:true}
	s := TelemetrySink{Backend:b}
	if err := s.Publish(context.Background(), Metric{Name:"ftn_test", Value:0.0/0.0}); err == nil { t.Fatal("expected NaN rejection") }
	long := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	if err := s.Publish(context.Background(), Metric{Name:"ftn_test", Value:1, Labels:map[string]string{"site":long}}); err != ErrInvalidMetric { t.Fatalf("got %v", err) }
}
