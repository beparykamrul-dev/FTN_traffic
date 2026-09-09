package dataplane
import("context";"testing")
type reconcileRouter struct{healthy bool;applied int}
func(r *reconcileRouter)Name()string{return "test"};func(r *reconcileRouter)Kind()RouterKind{return RouterFRR};func(r *reconcileRouter)Health(context.Context)error{if !r.healthy{return ErrBackendUnavailable};return nil};func(r *reconcileRouter)ApplyRoutes(_ context.Context,x []RouteIntent)error{r.applied=len(x);return nil};func(r *reconcileRouter)WithdrawRoutes(context.Context,[]RouteIntent)error{return nil};func(r *reconcileRouter)Snapshot(context.Context)(RouterStatus,error){return RouterStatus{ID:"test",Kind:RouterFRR,Healthy:r.healthy},nil}
func TestRouteReconciler(t *testing.T){r:=&reconcileRouter{healthy:true};c:=RouteReconciler{Router:r,Authorized:true,Approved:true};if err:=c.Apply(context.Background(),[]RouteIntent{{Prefix:"203.0.113.0/24",Family:IPv4,Authorized:true}});err!=nil{t.Fatal(err)};if r.applied!=1{t.Fatalf("applied=%d",r.applied)}}
func TestRouteReconcilerApproval(t *testing.T){r:=&reconcileRouter{healthy:true};if err:= (RouteReconciler{Router:r,Authorized:true}).Apply(context.Background(),nil);err!=ErrApprovalRequired{t.Fatalf("want approval, got %v",err)}}
