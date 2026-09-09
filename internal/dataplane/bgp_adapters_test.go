package dataplane

import("context";"testing")

type adapterRunner struct{out []byte;err error;name string;args []string}
func(r *adapterRunner)Run(_ context.Context,name string,args ...string)([]byte,error){r.name=name;r.args=args;return r.out,r.err}
func TestFRRBGPAdapterSummary(t *testing.T){r:=&adapterRunner{out:[]byte("Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd\n192.0.2.9 4 64509 1 1 0 0 0 00:10 3")};a:=FRRBGPAdapter{Runtime:FRRRuntime{Runner:r},ID:"r1"};s,e:=a.SessionSummary(context.Background());if e!=nil||len(s)!=1||s[0].RemoteASN!=64509{t.Fatalf("%v %#v",e,s)};if r.name!="vtysh"{t.Fatalf("command=%s",r.name)}}
func TestBGPAdaptersRequireExplicitMutationPath(t *testing.T){if e:= (FRRBGPAdapter{}).ApplyRoutes(context.Background(),nil);e!=ErrApprovalRequired{t.Fatalf("frr=%v",e)};if e:= (BIRDBGPAdapter{}).ApplyRoutes(context.Background(),nil);e!=ErrApprovalRequired{t.Fatalf("bird=%v",e)};if e:= (GoBGPBGPAdapter{}).ApplyRoutes(context.Background(),nil);e!=ErrApprovalRequired{t.Fatalf("gobgp=%v",e)}}
