package migration

import "testing"
func TestIdempotency(t *testing.T){i:=NewIdempotency();if !i.Claim("x"){t.Fatal("first claim failed")};if i.Claim("x"){t.Fatal("duplicate claim accepted")};if i.Claim(""){t.Fatal("empty key accepted")}}
