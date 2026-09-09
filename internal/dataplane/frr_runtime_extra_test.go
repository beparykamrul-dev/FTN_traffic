package dataplane

import (
	"context"
	"testing"
)

type frrRunnerStub struct{ out []byte; err error; name string; args []string }
func (r *frrRunnerStub) Run(_ context.Context, name string, args ...string)([]byte,error){r.name=name;r.args=args;return r.out,r.err}
func TestFRRSummaryOutputUsesSafeCommand(t *testing.T){r:=&frrRunnerStub{out:[]byte("ok")}; f:=FRRRuntime{Runner:r}; out,err:=f.SummaryOutput(context.Background()); if err!=nil||string(out)!="ok"||r.name!="vtysh"{t.Fatalf("unexpected: %q %v %#v",out,err,r)} }
func TestFRRRejectsControlCommand(t *testing.T){r:=&frrRunnerStub{}; f:=FRRRuntime{Runner:r,Authorized:true,Approved:true}; if err:=f.Apply(context.Background(),[]string{"show\nbgp"}); err!=ErrBGPRuntimeCommandInvalid{t.Fatalf("got %v",err)} }
