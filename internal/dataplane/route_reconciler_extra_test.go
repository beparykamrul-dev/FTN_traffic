package dataplane
import("context";"testing")
type reconRouter struct{err error}
func(r reconRouter)Name()string{return "test"}
func(r reconRouter)Health(context.Context)error{return r.err}
func(r reconRouter)ApplyRoutes(context.Context,[]RouteIntent)error{return nil}
func(r reconRouter)WithdrawRoutes(context.Context,[]RouteIntent)error{return nil}
func(r reconRouter)Snapshot(context.Context)(RouterStatus,error){return RouterStatus{},nil}
func TestRouteReconcilerRequiresApproval(t *testing.T){r:=RouteReconciler{Router:reconRouter{},Authorized:true};err:=r.Apply(context.Background(),[]RouteIntent{{Prefix:"203.0.113.0/24",Family:IPv4,Authorized:true}});if err!=ErrApprovalRequired{t.Fatalf("got %v",err)}}
