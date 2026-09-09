package dataplane

import "context"

type GoBGPRuntime struct { Runner CommandRunner; Authorized bool; Approved bool }
func (g GoBGPRuntime) Health(ctx context.Context) error { if g.Runner==nil{return ErrBackendUnavailable}; _,err:=g.Runner.Run(ctx,"gobgp","global","rib"); return err }
func (g GoBGPRuntime) Apply(ctx context.Context, args ...string) error { if !g.Authorized{return ErrUnauthorized}; if !g.Approved{return ErrApprovalRequired}; if g.Runner==nil{return ErrBackendUnavailable}; _,err:=g.Runner.Run(ctx,append([]string{"global"},args...)...); return err }
