package dataplane

import (
	"context"
	"errors"
)

var ErrBFDRequired = errors.New("BFD session must be up")

type BFDChecker interface {
	CheckBFD(context.Context) error
}

func RequireBFD(ctx context.Context, checker BFDChecker) error {
	if checker == nil {
		return ErrBackendUnavailable
	}
	if err := checker.CheckBFD(ctx); err != nil {
		return err
	}
	return nil
}
