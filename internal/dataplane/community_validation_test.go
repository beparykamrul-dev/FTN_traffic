package dataplane

import "testing"

func TestValidateCommunity(t *testing.T) {
	for _, v := range []string{"64512:100", "65000:65535"} {
		if err := ValidateCommunity(v); err != nil { t.Fatalf("%s: %v", v, err) }
	}
	for _, v := range []string{"64512", "64512:70000", "x:1", "1:2:3"} {
		if err := ValidateCommunity(v); err != ErrInvalidCommunity { t.Fatalf("%s: got %v", v, err) }
	}
}
