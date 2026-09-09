package migration

import "testing"

func TestValidateStrategy(t *testing.T) {
	for _, s := range []string{"ecmp","weighted","canary","failover","blue_green"} { if err:=ValidateStrategy(s); err!=nil { t.Fatal(err) } }
	if err:=ValidateStrategy("shell"); err==nil { t.Fatal("expected rejection") }
}
