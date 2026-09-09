package dataplane

import "errors"

var ErrBGPDown = errors.New("bgp session is down")

type BGPHealth struct {
	Established bool `json:"established"`
	RPKIValid   bool `json:"rpki_valid"`
}

func (h BGPHealth) Validate(requireRPKI bool) error {
	if !h.Established {
		return ErrBGPDown
	}
	if requireRPKI && !h.RPKIValid {
		return ErrRPKIInvalid
	}
	return nil
}
