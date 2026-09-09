package dns

import "testing"

func TestValidateTarget(t *testing.T) {
	if err := ValidateTarget(ProbeTarget{Name: "node-1", Address: "127.0.0.1:53"}); err != nil { t.Fatal(err) }
	if err := ValidateTarget(ProbeTarget{Name: "node-1", Address: "127.0.0.1"}); err == nil { t.Fatal("expected host:port validation error") }
	if err := ValidateTarget(ProbeTarget{Address: "127.0.0.1:53"}); err == nil { t.Fatal("expected name validation error") }
}
