package dataplane

import "context"

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
