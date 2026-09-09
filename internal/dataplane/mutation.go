package dataplane

import "errors"

var (
	ErrUnauthorized = errors.New("dataplane mutation unauthorized")
	ErrBackendUnavailable = errors.New("dataplane backend unavailable")
	ErrApprovalRequired = errors.New("dataplane mutation requires approval")
)

func ValidateMutation(authorized, approved, backendHealthy bool) error {
	if !authorized { return ErrUnauthorized }
	if !approved { return ErrApprovalRequired }
	if !backendHealthy { return ErrBackendUnavailable }
	return nil
}
