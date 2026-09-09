package dataplane

import "context"

type RouteMutationGuard interface { Authorize(context.Context,string) error }
type RouteMutationRequest struct { ApprovalID string; Routes []RouteIntent }
type RouteReconciler struct { Router RouterAdapter; Guard RouteMutationGuard }
func(r RouteReconciler) validate(ctx context.Context,x []RouteIntent,approval string)error{if r.Router==nil{return ErrBackendUnavailable};if err:=ValidateRoutes(x);err!=nil{return err};if r.Guard==nil||approval==""{return ErrApprovalRequired};return r.Guard.Authorize(ctx,approval)}
func(r RouteReconciler) Apply(ctx context.Context,desired []RouteIntent)error{if r.Router==nil{return ErrBackendUnavailable};if err:=ValidateRoutes(desired);err!=nil{return err};if err:=r.Router.Health(ctx);err!=nil{return err};return r.Router.ApplyRoutes(ctx,desired)}
func(r RouteReconciler) ApplyApproved(ctx context.Context,q RouteMutationRequest)error{if err:=r.validate(ctx,q.Routes,q.ApprovalID);err!=nil{return err};if err:=r.Router.Health(ctx);err!=nil{return err};return r.Router.ApplyRoutes(ctx,q.Routes)}
func(r RouteReconciler) Withdraw(ctx context.Context,routes []RouteIntent)error{if r.Router==nil{return ErrBackendUnavailable};if err:=ValidateRoutes(routes);err!=nil{return err};if err:=r.Router.Health(ctx);err!=nil{return err};return r.Router.WithdrawRoutes(ctx,routes)}
func(r RouteReconciler) WithdrawApproved(ctx context.Context,q RouteMutationRequest)error{if err:=r.validate(ctx,q.Routes,q.ApprovalID);err!=nil{return err};if err:=r.Router.Health(ctx);err!=nil{return err};return r.Router.WithdrawRoutes(ctx,q.Routes)}
