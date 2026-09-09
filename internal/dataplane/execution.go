package dataplane

import (
	"context"
	"errors"
	"os/exec"
)

var ErrCommandNotAllowed = errors.New("local command is not allowed")

type LocalExecutor struct { Allowed map[string]bool }
func (e LocalExecutor) Run(ctx context.Context, name string, args ...string) ([]byte,error) {
	if !e.Allowed[name] { return nil,ErrCommandNotAllowed }
	return exec.CommandContext(ctx,name,args...).CombinedOutput()
}

func DefaultLocalExecutor() LocalExecutor { return LocalExecutor{Allowed:map[string]bool{"ip":true,"tc":true,"nft":true,"vtysh":true,"birdc":true,"gobgp":true}} }
