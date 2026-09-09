package dataplane

import "testing"

func TestValidateRuntimeCommand(t *testing.T) {
	for _, name := range []string{"ip","tc","nft","vtysh","birdc","gobgp"} {
		if err := ValidateRuntimeCommand(name); err != nil { t.Fatalf("%s: %v", name, err) }
	}
	if err := ValidateRuntimeCommand("sh"); err != ErrCommandNotAllowed { t.Fatalf("got %v", err) }
}
