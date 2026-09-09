package dataplane

import "net/netip"

type RouterKind string
const ( RouterFRR RouterKind = "frr"; RouterBIRD RouterKind = "bird"; RouterGoBGP RouterKind = "gobgp" )
type AddressFamily string
const ( IPv4 AddressFamily = "ipv4"; IPv6 AddressFamily = "ipv6" )
type RouteIntent struct { Prefix string `json:"prefix"`; Family AddressFamily `json:"family"`; NextHop string `json:"next_hop"`; LocalPref uint32 `json:"local_pref"`; MED uint32 `json:"med"`; Community []string `json:"community"`; Authorized bool `json:"authorized"` }
type RouterStatus struct { ID string `json:"id"`; Kind RouterKind `json:"kind"`; Healthy bool `json:"healthy"`; BGPUp bool `json:"bgp_up"`; IPv4Routes uint64 `json:"ipv4_routes"`; IPv6Routes uint64 `json:"ipv6_routes"` }
type BGPSession struct { ID string `json:"id"`; PeerAddress string `json:"peer_address"`; ASN uint32 `json:"asn"`; RemoteASN uint32 `json:"remote_asn,omitempty"`; Established bool `json:"established"`; IPv4 bool `json:"ipv4"`; IPv6 bool `json:"ipv6"`; RPKIValid bool `json:"rpki_valid"`; PrefixesIn uint64 `json:"prefixes_in"` }
func (s BGPSession) Healthy(requireRPKI bool, maxPrefixes uint64) bool { if s.ID=="" || !s.Established || (!s.IPv4 && !s.IPv6) { return false }; if s.PeerAddress!="" { if _,err:=netip.ParseAddr(s.PeerAddress); err!=nil { return false } }; if requireRPKI && !s.RPKIValid { return false }; return maxPrefixes==0 || s.PrefixesIn<=maxPrefixes }
