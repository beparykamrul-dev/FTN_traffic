package dataplane

import "context"

type BIRDBGPAdapter struct { Runtime BIRDRuntime; ID string }
func (a BIRDBGPAdapter) Name() string { if a.ID != "" { return a.ID }; return "bird" }
func (a BIRDBGPAdapter) Kind() RouterKind { return RouterBIRD }
func (a BIRDBGPAdapter) Health(ctx context.Context) error { return a.Runtime.Health(ctx) }
func (a BIRDBGPAdapter) SessionSummary(ctx context.Context)([]BGPSession,error){if a.Runtime.Runner==nil{return nil,ErrBackendUnavailable};out,err:=a.Runtime.Runner.Run(ctx,"birdc","show protocols");if err!=nil{return nil,err};return ParseBGPSummary(RouterBIRD,string(out))}
func (a BIRDBGPAdapter) Snapshot(ctx context.Context)(RouterStatus,error){s,err:=a.SessionSummary(ctx);if err!=nil{return RouterStatus{ID:a.Name(),Kind:a.Kind()},err};up:=false;var v4 uint64;for _,x:=range s{if x.Established{up=true};v4+=x.PrefixesIn};return RouterStatus{ID:a.Name(),Kind:a.Kind(),Healthy:true,BGPUp:up,IPv4Routes:v4},nil}
func (a BIRDBGPAdapter) ApplyRoutes(context.Context,[]RouteIntent)error{return ErrApprovalRequired}
func (a BIRDBGPAdapter) WithdrawRoutes(context.Context,[]RouteIntent)error{return ErrApprovalRequired}
