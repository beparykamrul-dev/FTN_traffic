package dataplane

import (
	"context"
	"testing"
)

func TestLocalCommandRunnerAllowlist(t *testing.T) {
	r := LocalCommandRunner{}
	if _, err := r.Run(context.Background(), "sh", "-c", "echo unsafe"); err != ErrCommandNotAllowed { t.Fatalf("got %v", err) }
}
