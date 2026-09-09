package dataplane

import (
	"context"
	"errors"
	"os/exec"
)

var ErrCommandNotAllowed = errors.New("local command is not allowed")

type LocalCommandRunner struct{}

var allowedLocalCommands = map[string]bool{"ip":true,"tc":true,"nft":true,"vtysh":true,"birdc":true,"gobgp":true}

func (LocalCommandRunner) Run(ctx context.Context, name string, args ...string) ([]byte, error) {
	if !allowedLocalCommands[name] { return nil, ErrCommandNotAllowed }
	return exec.CommandContext(ctx, name, args...).CombinedOutput()
}
