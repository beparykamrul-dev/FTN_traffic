package dataplane

import "context"

type BGPAdapter interface {
	RouterAdapter
	SessionSummary(context.Context) ([]BGPSession, error)
}

type BGPSession struct {
	ID string `json:"id"`
	RemoteASN uint32 `json:"remote_asn"`
	Established bool `json:"established"`
	IPv4 bool `json:"ipv4"`
	IPv6 bool `json:"ipv6"`
	RPKIValid bool `json:"rpki_valid"`
	PrefixesIn uint64 `json:"prefixes_in"`
	PrefixesOut uint64 `json:"prefixes_out"`
}
