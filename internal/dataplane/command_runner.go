package dataplane

import (
	"context"
	"fmt"
	"os/exec"
)

type CommandRunner interface { Run(context.Context, string, ...string) ([]byte, error) }

type LocalCommandRunner struct{}

var allowedCommands = map[string]struct{}{"ip":{},"tc":{},"nft":{},"vtysh":{},"birdc":{},"gobgp":{}}

func (LocalCommandRunner) Run(ctx context.Context, name string, args ...string) ([]byte, error) {
	if _, ok := allowedCommands[name]; !ok { return nil, fmt.Errorf("command not allowed: %s", name) }
	return exec.CommandContext(ctx, name, args...).CombinedOutput()
}
