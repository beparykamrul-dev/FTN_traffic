package dataplane

import (
	"context"
	"errors"
	"testing"
)

type fakeCommandRunner struct { name string; args []string; err error }
func (f *fakeCommandRunner) Run(_ context.Context, name string, args ...string) ([]byte, error) { f.name=name; f.args=args; return nil, f.err }

func TestFRRRuntimePolicy(t *testing.T) {
	r := &fakeCommandRunner{}
	f := FRRRuntime{Runner:r}
	if err:=f.Apply(context.Background(), []string{"router bgp 65000"}); !errors.Is(err,ErrUnauthorized){t.Fatalf("got %v",err)}
	f.Authorized=true
	if err:=f.Apply(context.Background(), []string{"router bgp 65000"}); !errors.Is(err,ErrApprovalRequired){t.Fatalf("got %v",err)}
	f.Approved=true
	if err:=f.Apply(context.Background(), []string{"router bgp 65000"}); err!=nil{t.Fatal(err)}
}

func TestFRRRuntimePropagatesRunnerError(t *testing.T) {
	r := &fakeCommandRunner{err: errors.New("vtysh failed")}
	f := FRRRuntime{Runner:r, Authorized:true, Approved:true}
	if err:=f.Apply(context.Background(), []string{"show bgp summary"}); err==nil || err.Error()!="vtysh failed" { t.Fatalf("got %v",err) }
}
