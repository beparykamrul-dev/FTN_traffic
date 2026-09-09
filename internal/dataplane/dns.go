package dataplane

import "context"

type DNSBackend interface {
	Name() string
	Health(context.Context) error
	Resolve(context.Context, string, string) ([]string, error)
}

type DNSProfile struct {
	Recursive bool `json:"recursive"`
	Authoritative bool `json:"authoritative"`
	DNSSEC bool `json:"dnssec"`
	PublicRecursion bool `json:"public_recursion"`
}
