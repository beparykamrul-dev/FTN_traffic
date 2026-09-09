package upstream

import "testing"

func TestValidateSession(t *testing.T) {
	p := Policy{LocalASN: 65001, RequireMaxPrefix: true}
	good := Session{LocalASN: 65001, RemoteASN: 65002, IPv4: true, Authorized: true, MaxPrefixes: 100}
	if err := ValidateSession(good, p); err != nil {
		t.Fatal(err)
	}
	bad := good
	bad.Authorized = false
	if err := ValidateSession(bad, p); err != ErrUnauthorized {
		t.Fatalf("expected authorization failure, got %v", err)
	}
}
