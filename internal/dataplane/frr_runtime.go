package dataplane

import (
 "context"
 "strings"
)

type FRRRuntime struct { Runner CommandRunner; Authorized bool; Approved bool }
func (f FRRRuntime) Health(ctx context.Context) error { if f.Runner==nil{return ErrBackendUnavailable}; _,err:=f.Runner.Run(ctx,"vtysh","-c","show version"); return err }
func (f FRRRuntime) Apply(ctx context.Context, commands []string) error {
 if !f.Authorized{return ErrUnauthorized}; if !f.Approved{return ErrApprovalRequired}; if f.Runner==nil{return ErrBackendUnavailable}
 for _,cmd:=range commands { if strings.TrimSpace(cmd)=="" {return ErrRuntimeCommandInvalid}; if _,err:=f.Runner.Run(ctx,"vtysh","-c",cmd); err!=nil{return err} }
 return nil
}
