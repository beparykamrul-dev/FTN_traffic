package dataplane

import "context"

type FirewallBackend interface {
	Name() string
	Health(context.Context) error
	Apply(context.Context, []Rule) error
	Remove(context.Context, []Rule) error
}

type Rule struct {
	ID string `json:"id"`
	Family string `json:"family"`
	Direction string `json:"direction"`
	Action string `json:"action"`
	Source string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Interface string `json:"interface,omitempty"`
	Authorized bool `json:"authorized"`
}
