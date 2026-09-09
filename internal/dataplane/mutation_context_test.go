package dataplane

import "testing"

func TestMutationMetadataValidate(t *testing.T) {
	if err := (MutationMetadata{Actor:"admin",RequestID:"req-1",ApprovalID:"app-1"}).Validate(); err != nil { t.Fatal(err) }
	if err := (MutationMetadata{Actor:"admin",RequestID:"req-1"}).Validate(); err != ErrMutationMetadataIncomplete { t.Fatalf("got %v", err) }
}
