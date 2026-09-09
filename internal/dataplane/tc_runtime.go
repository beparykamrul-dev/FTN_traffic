package dataplane

import "context"

type TCRuntime struct { Runner CommandRunner; Authorized bool; Approved bool }
func (t TCRuntime) Health(ctx context.Context) error { if t.Runner==nil{return ErrBackendUnavailable}; _,err:=t.Runner.Run(ctx,"tc","qdisc","show"); return err }
func (t TCRuntime) Apply(ctx context.Context, args ...string) error { if !t.Authorized{return ErrUnauthorized}; if !t.Approved{return ErrApprovalRequired}; if t.Runner==nil{return ErrBackendUnavailable}; _,err:=t.Runner.Run(ctx,append([]string{"qdisc"},args...)...); return err }
