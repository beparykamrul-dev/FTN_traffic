package dataplane

import "context"

type FRRBGPAdapter struct { Runtime FRRRuntime; ID string }
func (a FRRBGPAdapter) Name() string { if a.ID != "" { return a.ID }; return "frr" }
func (a FRRBGPAdapter) Kind() RouterKind { return RouterFRR }
func (a FRRBGPAdapter) Health(ctx context.Context) error { return a.Runtime.Health(ctx) }
func (a FRRBGPAdapter) SessionSummary(ctx context.Context) ([]BGPSession,error) { if a.Runtime.Runner==nil{return nil,ErrBackendUnavailable}; out,err:=a.Runtime.Runner.Run(ctx,"vtysh","-c","show bgp summary");if err!=nil{return nil,err};return ParseBGPSummary(RouterFRR,string(out)) }
func (a FRRBGPAdapter) Snapshot(ctx context.Context)(RouterStatus,error){s,err:=a.SessionSummary(ctx);if err!=nil{return RouterStatus{ID:a.Name(),Kind:a.Kind()},err};var v4,v6 uint64;up:=false;for _,x:=range s{if x.Established{up=true};v4+=x.PrefixesIn};return RouterStatus{ID:a.Name(),Kind:a.Kind(),Healthy:true,BGPUp:up,IPv4Routes:v4,IPv6Routes:v6},nil}
func (a FRRBGPAdapter) ApplyRoutes(context.Context,[]RouteIntent) error{return ErrApprovalRequired}
func (a FRRBGPAdapter) WithdrawRoutes(context.Context,[]RouteIntent) error{return ErrApprovalRequired}
