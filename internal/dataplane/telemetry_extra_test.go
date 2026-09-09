package dataplane
import("math";"testing")
func TestValidateMetricRejectsInvalidLabel(t *testing.T){if err:=ValidateMetric(Metric{Name:"ftn_up",Labels:map[string]string{"bad-label":"x"}});err==nil{t.Fatal("expected invalid label")}}
func TestValidateMetricRejectsNonFinite(t *testing.T){for _,v:=range []float64{math.NaN(),math.Inf(1),math.Inf(-1)}{if err:=ValidateMetric(Metric{Name:"ftn_value",Value:v});err==nil{t.Fatal("expected non-finite rejection")}}}
