package dataplane

import("context";"strings")

type NFTRuntime struct { Runner CommandRunner; Authorized bool; Approved bool }
func (n NFTRuntime) Health(ctx context.Context) error {if n.Runner==nil{return ErrBackendUnavailable};_,err:=n.Runner.Run(ctx,"nft","list","ruleset");return err}
func (n NFTRuntime) Apply(ctx context.Context,args ...string) error {if !n.Authorized{return ErrUnauthorized};if !n.Approved{return ErrApprovalRequired};if n.Runner==nil{return ErrBackendUnavailable};if len(args)==0{return ErrRuntimeCommandInvalid};for _,a:=range args{if strings.TrimSpace(a)==""{return ErrRuntimeCommandInvalid}};_,err:=n.Runner.Run(ctx,"nft",args...);return err}
