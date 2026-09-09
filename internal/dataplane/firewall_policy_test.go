package dataplane

import "testing"

func TestValidateRules(t *testing.T) {
	r := Rule{ID:"r1", Family:"ipv4", Direction:"in", Action:"accept", Authorized:true}
	if err := ValidateRules([]Rule{r}); err != nil { t.Fatal(err) }
	r.Authorized = false
	if err := ValidateRule(r); err != ErrFirewallUnauthorized { t.Fatalf("got %v", err) }
}
