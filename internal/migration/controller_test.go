package migration

import (
	"context"
	"errors"
	"testing"
)

type fakeExec struct{ shifted, committed, rolled bool; strategy string; failShift bool }
func (f *fakeExec) Preflight(context.Context,Request) error{return nil}
func (f *fakeExec) Shift(_ context.Context,_ Request,s string) error{f.shifted=true;f.strategy=s;if f.failShift{return errors.New("shift")};return nil}
func (f *fakeExec) Commit(context.Context,Request) error{f.committed=true;return nil}
func (f *fakeExec) Rollback(context.Context,Request) error{f.rolled=true;return nil}
type fakeApproval struct{ ok bool }
func(f fakeApproval) Verify(context.Context,string) error{if !f.ok{return errors.New("denied")};return nil}
type fakeHealth struct{ ok bool }
func(f fakeHealth) Healthy(context.Context,Request) error{if !f.ok{return errors.New("unhealthy")};return nil}

func req(id string) Request{return Request{ID:id,From:"old",To:"new",ApprovalID:"ap-1",Strategy:"canary"}}
func TestMigrationCommit(t *testing.T){e:=&fakeExec{};c:=NewController(e,fakeApproval{true},fakeHealth{true});if err:=c.Run(context.Background(),req("1"));err!=nil{t.Fatal(err)};if !e.committed||e.strategy!="canary"||c.State()!=StateCommit{t.Fatalf("state=%s exec=%+v",c.State(),e)}}
func TestMigrationFailsClosedWithoutApproval(t *testing.T){e:=&fakeExec{};c:=NewController(e,fakeApproval{false},fakeHealth{true});if err:=c.Run(context.Background(),req("1"));err==nil{t.Fatal("expected denial")};if e.shifted{t.Fatal("shift occurred without approval")};if c.State()!=StateFailed{t.Fatalf("state=%s",c.State())}}
func TestMigrationRollsBackOnShiftFailure(t *testing.T){e:=&fakeExec{failShift:true};c:=NewController(e,fakeApproval{true},fakeHealth{true});if err:=c.Run(context.Background(),req("1"));err==nil{t.Fatal("expected rollback error")};if !e.rolled||c.State()!=StateRollback{t.Fatalf("state=%s rolled=%v",c.State(),e.rolled)}}
func TestMigrationRejectsInvalidStrategy(t *testing.T){c:=NewController(&fakeExec{},fakeApproval{true},fakeHealth{true});r:=req("bad");r.Strategy="invalid";if err:=c.Run(context.Background(),r);err!=ErrInvalidStrategy{t.Fatalf("got %v",err)}}
func TestMigrationRejectsDuplicate(t *testing.T){e:=&fakeExec{};c:=NewController(e,fakeApproval{true},fakeHealth{true});if err:=c.Run(context.Background(),req("dup"));err!=nil{t.Fatal(err)};if err:=c.Run(context.Background(),req("dup"));err!=ErrDuplicateMigration{t.Fatalf("got %v",err)}}
