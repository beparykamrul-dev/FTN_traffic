package dataplane

import "context"

type TransportBackend interface {
	Name() string
	Health(context.Context) error
	Connect(context.Context, Peer) error
	Disconnect(context.Context, Peer) error
}

type Peer struct {
	ID string `json:"id"`
	Endpoint string `json:"endpoint"`
	IPv4 bool `json:"ipv4"`
	IPv6 bool `json:"ipv6"`
	Authorized bool `json:"authorized"`
}
