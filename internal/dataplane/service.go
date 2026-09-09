package dataplane

import "context"

type Service struct {
	ID string `json:"id"`
	Kind string `json:"kind"`
	Backend string `json:"backend"`
	Enabled bool `json:"enabled"`
	Authorized bool `json:"authorized"`
}

type ServiceRegistry interface {
	Register(Service) error
	Get(string) (Service, bool)
	List() []Service
	Health(context.Context, string) error
}
