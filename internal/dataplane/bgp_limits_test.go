package dataplane

import "testing"

func TestValidateBGPFamily(t *testing.T) {
	s := BGPSession{ID:"p", Established:true, IPv4:true}
	if err := ValidateBGPFamily(s, IPv4); err != nil { t.Fatal(err) }
	if err := ValidateBGPFamily(s, IPv6); err == nil { t.Fatal("expected IPv6 rejection") }
}
