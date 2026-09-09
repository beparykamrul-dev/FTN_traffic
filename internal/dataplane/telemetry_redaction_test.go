package dataplane

import "testing"

func TestRedactMetric(t *testing.T) {
	m := RedactMetric(Metric{Name:"ftn_route_health", Value:1, Labels:map[string]string{"node":"pop1","token":"secret","customer_ip":"192.0.2.4"}})
	if m.Labels["node"] != "pop1" { t.Fatal("safe label was removed") }
	if _, ok := m.Labels["token"]; ok { t.Fatal("token leaked") }
	if _, ok := m.Labels["customer_ip"]; ok { t.Fatal("customer IP leaked") }
}
