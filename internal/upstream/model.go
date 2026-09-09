package upstream

import "time"

type Role string

const (
	RoleProvider Role = "provider"
	RolePeer     Role = "peer"
	RoleTransit  Role = "transit"
)

type Session struct {
	ID              string    `json:"id"`
	RemoteASN       uint32    `json:"remote_asn"`
	LocalASN        uint32    `json:"local_asn"`
	Role            Role      `json:"role"`
	IPv4            bool      `json:"ipv4"`
	IPv6            bool      `json:"ipv6"`
	Established     bool      `json:"established"`
	Authorized      bool      `json:"authorized"`
	MaxPrefixes     uint32    `json:"max_prefixes"`
	PrefixesIn      uint32    `json:"prefixes_in"`
	PrefixesOut     uint32    `json:"prefixes_out"`
	LastStateChange time.Time `json:"last_state_change"`
}

func (s Session) Eligible() bool {
	return s.Authorized && s.Established && s.LocalASN != 0 && s.RemoteASN != 0
}

// CanProvideUpstream reports whether FTN may expose this session as a
// downstream-provider path. It deliberately does not infer authorization
// from public routing information.
func CanProvideUpstream(s Session) bool {
	return s.Eligible() && (s.IPv4 || s.IPv6) && s.MaxPrefixes > 0
}
