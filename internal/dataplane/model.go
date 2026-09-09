package dataplane

type RouterKind string

const (
	RouterFRR RouterKind = "frr"
	RouterBIRD RouterKind = "bird"
	RouterGoBGP RouterKind = "gobgp"
)

type AddressFamily string
const (
	IPv4 AddressFamily = "ipv4"
	IPv6 AddressFamily = "ipv6"
)

type RouteIntent struct {
	Prefix string `json:"prefix"`
	Family AddressFamily `json:"family"`
	NextHop string `json:"next_hop"`
	LocalPref uint32 `json:"local_pref"`
	MED uint32 `json:"med"`
	Community []string `json:"community"`
	Authorized bool `json:"authorized"`
}

type RouterStatus struct {
	ID string `json:"id"`
	Kind RouterKind `json:"kind"`
	Healthy bool `json:"healthy"`
	BGPUp bool `json:"bgp_up"`
	IPv4Routes uint64 `json:"ipv4_routes"`
	IPv6Routes uint64 `json:"ipv6_routes"`
}
