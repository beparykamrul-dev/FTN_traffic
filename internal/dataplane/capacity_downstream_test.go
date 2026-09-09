package dataplane

import "testing"

func TestCapacityHeadroom(t *testing.T) {
	if err := (Capacity{CapacityMbps:1000, UsedMbps:700, ReservePercent:20}).HasHeadroom(); err != nil { t.Fatal(err) }
	if err := (Capacity{CapacityMbps:1000, UsedMbps:850, ReservePercent:20}).HasHeadroom(); err != ErrInsufficientHeadroom { t.Fatalf("got %v", err) }
}

func TestDownstreamProfile(t *testing.T) {
	p := DownstreamProfile{ID:"isp-a", ASN:65010, IPv4Transit:true, MaxPrefixes:1000, Authorized:true}
	if err := p.Validate(65000); err != nil { t.Fatal(err) }
	p.Authorized=false
	if err := p.Validate(65000); err != ErrUnauthorized { t.Fatalf("got %v", err) }
}
