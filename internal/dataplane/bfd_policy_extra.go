package dataplane

import "errors"

var ErrBFDPolicy = errors.New("BFD policy rejected")

func ValidateBFDPolicy(s BFDSession) error {
	if err := ValidateBFD(s); err != nil { return err }
	if s.MinRxMS > 60000 || s.MinTxMS > 60000 || s.Multiplier > 255 { return ErrBFDPolicy }
	if !s.Up { return ErrBFDRequired }
	return nil
}
