package dataplane

import (
 "context"
 "os/exec"
)

type CommandRunner interface { Run(context.Context,string,...string)([]byte,error) }
type LocalCommandRunner struct{}

func (LocalCommandRunner) Run(ctx context.Context,name string,args ...string)([]byte,error) {
 if err:=ValidateRuntimeCommand(name); err!=nil{return nil,err}
 return exec.CommandContext(ctx,name,args...).CombinedOutput()
}
