package dataplane

import "testing"

func TestValidateMetricRejectsInvalidNames(t *testing.T) {
	for _, name := range []string{"", "bad-name", "1bad", "bad\nname"} {
		if err := ValidateMetric(Metric{Name:name, Value:1}); err != ErrInvalidMetric { t.Fatalf("%q: got %v", name, err) }
	}
	if err := ValidateMetric(Metric{Name:"ftn_packets_total", Value:1, Labels:map[string]string{"router_id":"r1"}}); err != nil { t.Fatal(err) }
}

func TestValidateMetricRejectsOversizedLabels(t *testing.T) {
	labels := make(map[string]string, maxMetricLabels+1)
	for i:=0; i<maxMetricLabels+1; i++ { labels["label_"+string(rune('a'+i))] = "v" }
	if err := ValidateMetric(Metric{Name:"ftn_test", Value:1, Labels:labels}); err != ErrInvalidMetric { t.Fatalf("got %v", err) }
	long := "x"
	for len(long) <= maxLabelValueLen { long += "x" }
	if err := ValidateMetric(Metric{Name:"ftn_test", Value:1, Labels:map[string]string{"site":long}}); err != ErrInvalidMetric { t.Fatalf("got %v", err) }
}
