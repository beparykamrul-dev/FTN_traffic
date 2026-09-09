package dataplane

func ValidateBFD(s BFDSession) error {
	if err := s.Validate(); err != nil { return err }
	if s.Local == s.Remote { return ErrBFDInvalid }
	return nil
}
