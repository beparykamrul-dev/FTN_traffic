package migration

import("context";"crypto/sha256";"encoding/hex";"errors";"fmt";"sync")
type State string
const(StatePreflight State="preflight";StateHealthGate State="health_gate";StateCanary State="canary";StateObserve State="observe";StateCommit State="commit";StateRollback State="rollback";StateFailed State="failed")
var ErrInvalidStrategy=errors.New("invalid migration strategy")
var ErrDuplicateMigration=errors.New("duplicate migration idempotency key")
var ErrMigrationFingerprint=errors.New("migration idempotency key reused with different request")
type Request struct{ID,From,To,ApprovalID string;Strategy string}
type Executor interface{Preflight(context.Context,Request)error;Shift(context.Context,Request,string)error;Commit(context.Context,Request)error;Rollback(context.Context,Request)error}
type ApprovalVerifier interface{Verify(context.Context,string)error};type HealthGate interface{Healthy(context.Context,Request)error};type Auditor interface{Record(context.Context,string,Request,string)error}
type Controller struct{mu sync.Mutex;state State;exec Executor;approval ApprovalVerifier;health HealthGate;audit Auditor;idempotency *Idempotency}
func NewController(e Executor,a ApprovalVerifier,h HealthGate)*Controller{return &Controller{state:StatePreflight,exec:e,approval:a,health:h,idempotency:NewIdempotency()}}
func(c *Controller)WithAudit(a Auditor)*Controller{c.audit=a;return c};func(c *Controller)State()State{c.mu.Lock();defer c.mu.Unlock();return c.state};func(c *Controller)set(s State){c.state=s}
func fingerprint(r Request)string{h:=sha256.Sum256([]byte(r.ID+"\x00"+r.From+"\x00"+r.To+"\x00"+r.ApprovalID+"\x00"+r.Strategy));return hex.EncodeToString(h[:])}
func(c *Controller)stateSet(s State,r Request){c.mu.Lock();c.set(s);c.mu.Unlock();c.idempotency.SetState(r.ID,s)}
func(c *Controller)record(ctx context.Context,event string,r Request,detail string){if c.audit!=nil{_=c.audit.Record(ctx,event,r,detail)}}
func(c *Controller)release(r Request,fp string){c.idempotency.Release(r.ID,fp)}
func(c *Controller)Run(ctx context.Context,r Request)error{
 if r.ID==""||r.From==""||r.To==""||r.ApprovalID==""{return errors.New("migration id, source, destination and approval are required")};if r.From==r.To{return errors.New("migration source and destination must differ")};if err:=ValidateStrategy(r.Strategy);err!=nil{return err};if c.exec==nil||c.approval==nil||c.health==nil||c.idempotency==nil{return errors.New("migration dependencies are incomplete")}
 fp:=fingerprint(r);if !c.idempotency.Claim(r.ID,fp){if !c.idempotency.Matches(r.ID,fp){return ErrMigrationFingerprint};return ErrDuplicateMigration}
 if err:=c.approval.Verify(ctx,r.ApprovalID);err!=nil{c.stateSet(StateFailed,r);c.record(ctx,"migration_denied",r,err.Error());c.release(r,fp);return fmt.Errorf("approval denied: %w",err)}
 c.stateSet(StatePreflight,r);c.record(ctx,"migration_preflight",r,"");if err:=c.exec.Preflight(ctx,r);err!=nil{c.stateSet(StateFailed,r);c.release(r,fp);return err};c.stateSet(StateHealthGate,r)
 if err:=c.health.Healthy(ctx,r);err!=nil{c.stateSet(StateFailed,r);c.release(r,fp);return fmt.Errorf("health gate failed: %w",err)};c.stateSet(StateCanary,r);if err:=c.exec.Shift(ctx,r,r.Strategy);err!=nil{return c.rollback(ctx,r,fp,err)};c.stateSet(StateObserve,r)
 if err:=c.health.Healthy(ctx,r);err!=nil{return c.rollback(ctx,r,fp,fmt.Errorf("post-shift health failed: %w",err))};c.stateSet(StateCommit,r);if err:=c.exec.Commit(ctx,r);err!=nil{return c.rollback(ctx,r,fp,err)};c.record(ctx,"migration_committed",r,"");return nil
}
func(c *Controller)rollback(ctx context.Context,r Request,fp string,cause error)error{c.stateSet(StateRollback,r);c.record(ctx,"migration_rollback",r,cause.Error());if err:=c.exec.Rollback(ctx,r);err!=nil{c.stateSet(StateFailed,r);c.release(r,fp);return fmt.Errorf("%v; rollback failed: %w",cause,err)};c.release(r,fp);return fmt.Errorf("migration rolled back: %w",cause)}
