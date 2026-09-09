package dataplane

import "errors"

var ErrTelemetryPolicy = errors.New("telemetry policy rejected")

type TelemetryPolicy struct { AllowCustomerPayload bool; AllowSecrets bool; MaxLabels int }
func (p TelemetryPolicy) Validate(m Metric) error {
	if p.AllowCustomerPayload || p.AllowSecrets { return ErrTelemetryPolicy }
	if p.MaxLabels > 0 && len(m.Labels) > p.MaxLabels { return ErrTelemetryPolicy }
	return nil
}
