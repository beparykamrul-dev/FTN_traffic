package dataplane

import("context";"testing")
func TestBIRDSummaryOutput(t *testing.T){r:=&frrRunnerStub{out:[]byte("ok")}; b:=BIRDRuntime{Runner:r}; out,err:=b.SummaryOutput(context.Background()); if err!=nil||string(out)!="ok"||r.name!="birdc"{t.Fatalf("unexpected %q %v %#v",out,err,r)} }
func TestBIRDRejectsControlCommand(t *testing.T){r:=&frrRunnerStub{}; b:=BIRDRuntime{Runner:r,Authorized:true,Approved:true}; if err:=b.Apply(context.Background(),"show\nprotocols");err!=ErrBGPRuntimeCommandInvalid{t.Fatalf("got %v",err)} }
