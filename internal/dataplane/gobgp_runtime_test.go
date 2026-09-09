package dataplane

import("context";"errors";"testing")

func TestGoBGPRuntimePolicy(t *testing.T){r:=&fakeCommandRunner{};g:=GoBGPRuntime{Runner:r,Authorized:true,Approved:true};if err:=g.Apply(context.Background());!errors.Is(err,ErrRuntimeCommandInvalid){t.Fatalf("got %v",err)};if err:=g.Apply(context.Background(),"rib");err!=nil{t.Fatal(err)}}

func TestGoBGPRuntimePropagatesRunnerError(t *testing.T){r:=&fakeCommandRunner{err:errors.New("gobgp failed")};g:=GoBGPRuntime{Runner:r,Authorized:true,Approved:true};if err:=g.Apply(context.Background(),"rib");err==nil||err.Error()!="gobgp failed"{t.Fatalf("got %v",err)}}
