package migration

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type State string
const (
	StatePreflight State = "preflight"
	StateHealthGate State = "health_gate"
	StateCanary State = "canary"
	StateObserve State = "observe"
	StateCommit State = "commit"
	StateRollback State = "rollback"
	StateFailed State = "failed"
)

type Request struct { ID, From, To, ApprovalID string; Strategy string }
type Executor interface { Preflight(context.Context, Request) error; Shift(context.Context, Request, string) error; Commit(context.Context, Request) error; Rollback(context.Context, Request) error }
type ApprovalVerifier interface { Verify(context.Context, string) error }
type HealthGate interface { Healthy(context.Context, Request) error }

type Controller struct { mu sync.Mutex; state State; exec Executor; approval ApprovalVerifier; health HealthGate }
func NewController(e Executor, a ApprovalVerifier, h HealthGate) *Controller { return &Controller{state:StatePreflight,exec:e,approval:a,health:h} }
func (c *Controller) State() State { c.mu.Lock(); defer c.mu.Unlock(); return c.state }
func (c *Controller) set(s State) { c.state=s }
func (c *Controller) Run(ctx context.Context, r Request) error {
	if r.ID=="" || r.From=="" || r.To=="" || r.ApprovalID=="" { return errors.New("migration id, source, destination and approval are required") }
	if c.exec==nil || c.approval==nil || c.health==nil { return errors.New("migration dependencies are incomplete") }
	c.mu.Lock(); c.set(StatePreflight); c.mu.Unlock()
	if err:=c.approval.Verify(ctx,r.ApprovalID); err!=nil { c.mu.Lock(); c.set(StateFailed); c.mu.Unlock(); return fmt.Errorf("approval denied: %w",err) }
	if err:=c.exec.Preflight(ctx,r); err!=nil { c.mu.Lock(); c.set(StateFailed); c.mu.Unlock(); return err }
	c.mu.Lock(); c.set(StateHealthGate); c.mu.Unlock()
	if err:=c.health.Healthy(ctx,r); err!=nil { c.mu.Lock(); c.set(StateFailed); c.mu.Unlock(); return fmt.Errorf("health gate failed: %w",err) }
	c.mu.Lock(); c.set(StateCanary); c.mu.Unlock()
	if err:=c.exec.Shift(ctx,r,"canary"); err!=nil { return c.rollback(ctx,r,err) }
	c.mu.Lock(); c.set(StateObserve); c.mu.Unlock()
	if err:=c.health.Healthy(ctx,r); err!=nil { return c.rollback(ctx,r,fmt.Errorf("post-canary health failed: %w",err)) }
	c.mu.Lock(); c.set(StateCommit); c.mu.Unlock()
	if err:=c.exec.Commit(ctx,r); err!=nil { return c.rollback(ctx,r,err) }
	return nil
}
func (c *Controller) rollback(ctx context.Context,r Request,cause error) error { c.mu.Lock(); c.set(StateRollback); c.mu.Unlock(); if err:=c.exec.Rollback(ctx,r); err!=nil { c.mu.Lock(); c.set(StateFailed); c.mu.Unlock(); return fmt.Errorf("%v; rollback failed: %w",cause,err) }; return fmt.Errorf("migration rolled back: %w",cause) }
