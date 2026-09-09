package dataplane

import (
	"context"
	"testing"
)

func TestRuntimeCommandRejectsShell(t *testing.T) {
	_, err := (RuntimeCommand{}).Run(context.Background(), "sh", "-c", "echo unsafe")
	if err != ErrCommandNotAllowed { t.Fatalf("got %v", err) }
}
