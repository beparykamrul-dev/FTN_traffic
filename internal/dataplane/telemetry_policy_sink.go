package dataplane

import "context"

// PolicyTelemetrySink applies redaction and an explicit telemetry policy
// before forwarding metrics to an external observability backend.
type PolicyTelemetrySink struct {
	Backend TelemetryBackend
	Policy  TelemetryPolicy
}

func (s PolicyTelemetrySink) Publish(ctx context.Context, m Metric) error {
	if s.Backend == nil {
		return ErrBackendUnavailable
	}
	m = RedactMetric(m)
	if err := s.Policy.Validate(m); err != nil {
		return err
	}
	if err := ValidateMetric(m); err != nil {
		return err
	}
	if err := s.Backend.Health(ctx); err != nil {
		return err
	}
	return s.Backend.Publish(ctx, m)
}
