package dataplane

import "context"

type TelemetrySink struct { Backend TelemetryBackend }

func (s TelemetrySink) Publish(ctx context.Context, m Metric) error {
	if s.Backend == nil { return ErrBackendUnavailable }
	m = RedactMetric(m)
	if err := ValidateMetric(m); err != nil { return err }
	if err := s.Backend.Health(ctx); err != nil { return err }
	return s.Backend.Publish(ctx, m)
}
