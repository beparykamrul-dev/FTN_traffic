package dataplane

import (
	"context"
	"errors"
	"math"
	"regexp"
	"strings"
)

type TelemetryBackend interface { Name() string; Health(context.Context) error; Publish(context.Context, Metric) error }
type Metric struct { Name string `json:"name"`; Value float64 `json:"value"`; Labels map[string]string `json:"labels,omitempty"` }

var ErrInvalidMetric = errors.New("invalid telemetry metric")
var metricNameRE = regexp.MustCompile(`^[a-zA-Z_:][a-zA-Z0-9_:]*$`)
var labelNameRE = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

const (
	maxMetricNameLen = 128
	maxLabelNameLen = 64
	maxLabelValueLen = 256
	maxMetricLabels = 32
)

func ValidateMetric(m Metric) error {
	name := strings.TrimSpace(m.Name)
	if name == "" || len(name) > maxMetricNameLen || !metricNameRE.MatchString(name) || strings.ContainsAny(name, "\n\r\x00") {
		return ErrInvalidMetric
	}
	if math.IsNaN(m.Value) || math.IsInf(m.Value, 0) {
		return ErrInvalidMetric
	}
	if len(m.Labels) > maxMetricLabels {
		return ErrInvalidMetric
	}
	for k, v := range m.Labels {
		if k == "" || len(k) > maxLabelNameLen || !labelNameRE.MatchString(k) || len(v) > maxLabelValueLen || strings.ContainsAny(k+v, "\n\r\x00") {
			return ErrInvalidMetric
		}
	}
	return nil
}
