package dataplane

import "net/netip"

func ValidateBFD(s BFDSession) error { if err:=s.Validate();err!=nil{return err}; local,_:=netip.ParseAddr(s.Local);remote,_:=netip.ParseAddr(s.Remote);if local==remote{return ErrBFDInvalid};return nil }
