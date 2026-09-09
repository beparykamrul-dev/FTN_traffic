package dataplane

import("context";"testing")
func TestGoBGPSummaryOutput(t *testing.T){r:=&frrRunnerStub{out:[]byte("ok")}; g:=GoBGPRuntime{Runner:r}; out,err:=g.SummaryOutput(context.Background()); if err!=nil||string(out)!="ok"||r.name!="gobgp"{t.Fatalf("unexpected %q %v %#v",out,err,r)} }
func TestGoBGPRejectsControlArgument(t *testing.T){r:=&frrRunnerStub{}; g:=GoBGPRuntime{Runner:r,Authorized:true,Approved:true}; if err:=g.Apply(context.Background(),"neighbor\n");err!=ErrBGPRuntimeCommandInvalid{t.Fatalf("got %v",err)} }
