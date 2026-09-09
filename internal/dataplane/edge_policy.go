package dataplane

import "errors"

var ErrEdgeUnauthorized = errors.New("edge mutation unauthorized")
var ErrPurgeApprovalRequired = errors.New("edge purge requires approval")

func ValidatePurge(authorized, approved bool, key string) error {
	if !authorized { return ErrEdgeUnauthorized }
	if !approved { return ErrPurgeApprovalRequired }
	if key == "" { return errors.New("cache key is required") }
	return nil
}
