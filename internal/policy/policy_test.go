package policy

import "testing"

func TestAllowMutation(t *testing.T) { if err := AllowMutation(false); err != ErrApprovalRequired { t.Fatalf("expected approval error, got %v", err) }; if err := AllowMutation(true); err != nil { t.Fatalf("unexpected error: %v", err) } }
