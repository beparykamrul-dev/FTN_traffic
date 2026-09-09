package acquisition

import "fmt"

type Request struct {
	Mode                 Mode `json:"mode"`
	ThirdParty           bool `json:"third_party"`
	PublicOrAuthorized   bool `json:"public_or_authorized"`
}

func Plan(r Request) (Source, error) {
	for _, s := range Sources {
		if s.Mode != r.Mode {
			continue
		}
		if s.RequiresAuthorization && !r.PublicOrAuthorized {
			return Source{}, fmt.Errorf("acquisition mode %q requires authorization", r.Mode)
		}
		if r.ThirdParty && !r.PublicOrAuthorized && !s.PublicDistribution {
			return Source{}, fmt.Errorf("third-party traffic requires an authorized source")
		}
		return s, nil
	}
	return Source{}, fmt.Errorf("unsupported acquisition mode %q", r.Mode)
}
