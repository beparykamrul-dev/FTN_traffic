package dataplane

import "testing"

func TestValidateMetricRejectsInvalidNames(t *testing.T) {
	for _, name := range []string{"", "bad-name", "1bad", "bad\nname"} {
		if err := ValidateMetric(Metric{Name:name, Value:1}); err != ErrInvalidMetric { t.Fatalf("%q: got %v", name, err) }
	}
	if err := ValidateMetric(Metric{Name:"ftn_packets_total", Value:1, Labels:map[string]string{"router_id":"r1"}}); err != nil { t.Fatal(err) }
}
