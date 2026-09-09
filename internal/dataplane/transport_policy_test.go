package dataplane

import "testing"

func TestValidatePeer(t *testing.T) {
	p := Peer{ID:"p1", Endpoint:"192.0.2.2:51820", IPv4:true, Authorized:true}
	if err := ValidatePeer(p); err != nil { t.Fatal(err) }
	p.Authorized = false
	if err := ValidatePeer(p); err != ErrTransportUnauthorized { t.Fatalf("got %v", err) }
}
