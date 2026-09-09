package policy

import "errors"

var ErrApprovalRequired = errors.New("explicit approval required")

func AllowMutation(approved bool) error { if !approved { return ErrApprovalRequired }; return nil }
