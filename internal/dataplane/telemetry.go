package dataplane

import (
	"context"
	"errors"
	"math"
	"regexp"
	"strings"
)

type TelemetryBackend interface {
	Name() string
	Health(context.Context) error
	Publish(context.Context, Metric) error
}

type Metric struct {
	Name   string            `json:"name"`
	Value  float64           `json:"value"`
	Labels map[string]string `json:"labels,omitempty"`
}

var (
	ErrInvalidMetric = errors.New("invalid telemetry metric")
	metricNameRE     = regexp.MustCompile(`^[a-zA-Z_:][a-zA-Z0-9_:]*$`)
	labelNameRE      = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
)

func ValidateMetric(m Metric) error {
	if !metricNameRE.MatchString(strings.TrimSpace(m.Name)) || strings.ContainsAny(m.Name, "\n\r") {
		return ErrInvalidMetric
	}
	if math.IsNaN(m.Value) || math.IsInf(m.Value, 0) {
		return ErrInvalidMetric
	}
	for k, v := range m.Labels {
		if !labelNameRE.MatchString(k) || strings.ContainsAny(k+v, "\n\r") {
			return ErrInvalidMetric
		}
	}
	return nil
}
