package dataplane

import "context"

type BFDSessionChecker struct { Session BFDSession }

func (c BFDSessionChecker) CheckBFD(context.Context) error {
	if err := ValidateBFD(c.Session); err != nil {
		return err
	}
	if !c.Session.Up {
		return ErrBFDRequired
	}
	return nil
}
