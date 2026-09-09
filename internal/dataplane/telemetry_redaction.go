package dataplane

import "strings"

var sensitiveLabelNames = map[string]struct{}{
	"token":{}, "authorization":{}, "credential":{}, "credential_ref":{}, "secret":{}, "password":{}, "private_key":{}, "customer_ip":{}, "payload":{},
}

func RedactMetric(m Metric) Metric {
	out := Metric{Name:m.Name, Value:m.Value}
	if len(m.Labels) == 0 { return out }
	out.Labels = make(map[string]string, len(m.Labels))
	for k, v := range m.Labels {
		lk := strings.ToLower(strings.TrimSpace(k))
		if _, sensitive := sensitiveLabelNames[lk]; sensitive { continue }
		out.Labels[k] = v
	}
	return out
}
