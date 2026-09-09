package providers

type Provider struct { Name string `json:"name"`; Role string `json:"role"`; Capabilities []string `json:"capabilities"`; Healthy bool `json:"healthy"` }

func Supports(p Provider, capability string) bool { for _, c := range p.Capabilities { if c == capability { return true } }; return false }
