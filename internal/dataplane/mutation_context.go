package dataplane

import "errors"

var ErrMutationMetadataIncomplete = errors.New("mutation metadata incomplete")

type MutationMetadata struct { Actor string; RequestID string; ApprovalID string }

func (m MutationMetadata) Validate() error {
	if m.Actor == "" || m.RequestID == "" || m.ApprovalID == "" { return ErrMutationMetadataIncomplete }
	return nil
}
