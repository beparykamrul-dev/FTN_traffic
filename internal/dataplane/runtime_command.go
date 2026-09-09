package dataplane

import (
	"context"
	"os/exec"
)

type RuntimeCommand struct{}

func (RuntimeCommand) Run(ctx context.Context, name string, args ...string) ([]byte, error) {
	if err := ValidateRuntimeCommand(name); err != nil { return nil, err }
	return exec.CommandContext(ctx, name, args...).CombinedOutput()
}
