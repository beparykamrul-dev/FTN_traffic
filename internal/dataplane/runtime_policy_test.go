package dataplane

import (
	"context"
	"testing"
)

type stubRunner struct{ name string; args []string }
func (s *stubRunner) Run(_ context.Context, name string, args ...string) ([]byte,error) { s.name=name; s.args=args; return nil,nil }

func TestFRRRuntimeApprovalGate(t *testing.T) {
	r := &stubRunner{}
	f := FRRRuntime{Runner:r,Authorized:true,Approved:false}
	if err := f.Apply(context.Background(),"configure terminal"); err != ErrApprovalRequired { t.Fatalf("got %v",err) }
	f.Approved=true
	if err := f.Apply(context.Background(),"show bgp summary"); err != nil { t.Fatal(err) }
	if r.name != "vtysh" { t.Fatalf("command=%s",r.name) }
}
