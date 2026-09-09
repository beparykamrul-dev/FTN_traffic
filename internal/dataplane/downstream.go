package dataplane

import "errors"

type DownstreamProfile struct {
	ID string `json:"id"`
	ASN uint32 `json:"asn"`
	IPv4Transit bool `json:"ipv4_transit"`
	IPv6Transit bool `json:"ipv6_transit"`
	DefaultRoute bool `json:"default_route"`
	FullTable bool `json:"full_table"`
	MaxPrefixes uint32 `json:"max_prefixes"`
	RPKIRequired bool `json:"rpki_required"`
	BFDRequired bool `json:"bfd_required"`
	Authorized bool `json:"authorized"`
}

func (p DownstreamProfile) Validate(localASN uint32) error {
	if !p.Authorized { return ErrUnauthorized }
	if p.ASN == 0 || localASN == 0 || p.ASN == localASN { return errors.New("invalid downstream ASN") }
	if !p.IPv4Transit && !p.IPv6Transit { return errors.New("at least one address family is required") }
	if p.MaxPrefixes == 0 { return errors.New("max-prefix limit is required") }
	return nil
}
