package dataplane

import "errors"

var ErrRPKIUnknown = errors.New("RPKI state is unknown")

func EnforceRPKI(state RPKIState, required bool) error {
	if !required { return nil }
	switch state {
	case RPKIValid:
		return nil
	case RPKIInvalid:
		return ErrRPKIInvalid
	default:
		return ErrRPKIUnknown
	}
}
