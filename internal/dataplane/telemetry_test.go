package dataplane

import (
 "math"
 "testing"
)

func TestValidateMetric(t *testing.T) {
 if err:=ValidateMetric(Metric{Name:"bgp_session_up",Value:1,Labels:map[string]string{"router":"r1"}});err!=nil{t.Fatal(err)}
 for _,m:=range []Metric{{Value:1},{Name:"bad\nname",Value:1},{Name:"x",Value:math.NaN()},{Name:"x",Value:1,Labels:map[string]string{"":"v"}}}{
  if err:=ValidateMetric(m);err!=ErrInvalidMetric{t.Fatalf("expected metric error, got %v",err)}
 }
}
