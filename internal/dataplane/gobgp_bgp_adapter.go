package dataplane

import "context"

type GoBGPBGPAdapter struct { Runtime GoBGPRuntime; ID string }
func (a GoBGPBGPAdapter) Name() string { if a.ID != "" { return a.ID }; return "gobgp" }
func (a GoBGPBGPAdapter) Kind() RouterKind { return RouterGoBGP }
func (a GoBGPBGPAdapter) Health(ctx context.Context) error { return a.Runtime.Health(ctx) }
func (a GoBGPBGPAdapter) SessionSummary(ctx context.Context)([]BGPSession,error){if a.Runtime.Runner==nil{return nil,ErrBackendUnavailable};out,err:=a.Runtime.Runner.Run(ctx,"gobgp","neighbor");if err!=nil{return nil,err};return ParseBGPSummary(RouterGoBGP,string(out))}
func (a GoBGPBGPAdapter) Snapshot(ctx context.Context)(RouterStatus,error){s,err:=a.SessionSummary(ctx);if err!=nil{return RouterStatus{ID:a.Name(),Kind:a.Kind()},err};up:=false;var v4 uint64;for _,x:=range s{if x.Established{up=true};v4+=x.PrefixesIn};return RouterStatus{ID:a.Name(),Kind:a.Kind(),Healthy:true,BGPUp:up,IPv4Routes:v4},nil}
func (a GoBGPBGPAdapter) ApplyRoutes(context.Context,[]RouteIntent)error{return ErrApprovalRequired}
func (a GoBGPBGPAdapter) WithdrawRoutes(context.Context,[]RouteIntent)error{return ErrApprovalRequired}
