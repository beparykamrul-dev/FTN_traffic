package dataplane

import "testing"

func TestNormalizeRouteBatchDeterministic(t *testing.T) {
	got := NormalizeRouteBatch([]RouteIntent{{Family:IPv6, Prefix:"2001:db8::/32", NextHop:"2001:db8::2"},{Family:IPv4, Prefix:"192.0.2.0/24", NextHop:"192.0.2.1", Community:[]string{" 64500:1 "}}})
	if len(got) != 2 || got[0].Family != IPv4 { t.Fatalf("unexpected order: %+v", got) }
	if got[1].Community == nil { return }
}
