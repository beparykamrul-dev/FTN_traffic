package dataplane

import (
	"context"
	"errors"
	"math"
	"strings"
)

type TelemetryBackend interface {
	Name() string
	Health(context.Context) error
	Publish(context.Context, Metric) error
}

type Metric struct {
	Name string `json:"name"`
	Value float64 `json:"value"`
	Labels map[string]string `json:"labels,omitempty"`
}

var ErrInvalidMetric = errors.New("invalid telemetry metric")

func ValidateMetric(m Metric) error {
	if strings.TrimSpace(m.Name) == "" || strings.ContainsAny(m.Name, "\n\r") {
		return ErrInvalidMetric
	}
	if math.IsNaN(m.Value) || math.IsInf(m.Value, 0) {
		return ErrInvalidMetric
	}
	for k, v := range m.Labels {
		if strings.TrimSpace(k) == "" || strings.ContainsAny(k+v, "\n\r") {
			return ErrInvalidMetric
		}
	}
	return nil
}
