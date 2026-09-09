package dataplane

import "context"

type BIRDRuntime struct { Runner CommandRunner; Authorized bool; Approved bool }
func (b BIRDRuntime) Health(ctx context.Context) error { if b.Runner==nil{return ErrBackendUnavailable}; _,err:=b.Runner.Run(ctx,"birdc","show status"); return err }
func (b BIRDRuntime) SummaryOutput(ctx context.Context) (string,error) { if b.Runner==nil{return "",ErrBackendUnavailable}; out,err:=b.Runner.Run(ctx,"birdc","show protocols"); return string(out),err }
func (b BIRDRuntime) Apply(ctx context.Context, command string) error { if !b.Authorized{return ErrUnauthorized}; if !b.Approved{return ErrApprovalRequired}; if b.Runner==nil{return ErrBackendUnavailable}; if err:=ValidateBGPRuntimeCommand(command);err!=nil{return err}; _,err:=b.Runner.Run(ctx,"birdc",command); return err }
