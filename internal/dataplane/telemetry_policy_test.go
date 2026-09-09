package dataplane

import "testing"

func TestTelemetryPolicy(t *testing.T) {
	p := TelemetryPolicy{MaxLabels:1}
	if err := p.Validate(Metric{Name:"ftn_up",Value:1,Labels:map[string]string{"service":"traffic"}}); err != nil { t.Fatal(err) }
	if err := p.Validate(Metric{Name:"ftn_up",Value:1,Labels:map[string]string{"a":"b","c":"d"}}); err == nil { t.Fatal("expected cardinality rejection") }
	if err := (TelemetryPolicy{AllowSecrets:true}).Validate(Metric{Name:"x"}); err == nil { t.Fatal("expected secret policy rejection") }
}
