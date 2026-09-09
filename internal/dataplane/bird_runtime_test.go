package dataplane

import("context";"errors";"testing")

func TestBIRDRuntimePolicy(t *testing.T){r:=&fakeCommandRunner{};b:=BIRDRuntime{Runner:r,Authorized:true};if err:=b.Apply(context.Background(),"");!errors.Is(err,ErrApprovalRequired){t.Fatalf("got %v",err)};b.Approved=true;if err:=b.Apply(context.Background(),"");!errors.Is(err,ErrRuntimeCommandInvalid){t.Fatalf("got %v",err)}}

func TestBIRDRuntimePropagatesRunnerError(t *testing.T){r:=&fakeCommandRunner{err:errors.New("birdc failed")};b:=BIRDRuntime{Runner:r,Authorized:true,Approved:true};if err:=b.Apply(context.Background(),"show protocols");err==nil||err.Error()!="birdc failed"{t.Fatalf("got %v",err)}}
