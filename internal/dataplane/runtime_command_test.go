package dataplane

import("context";"testing")

func TestRuntimeCommandRejectsShell(t *testing.T){_,err:=(RuntimeCommand{}).Run(context.Background(),"sh","-c","echo unsafe");if err!=ErrCommandNotAllowed{t.Fatalf("got %v",err)}}
func TestRuntimeCommandRejectsInvalidArgs(t *testing.T){_,err:=(RuntimeCommand{}).Run(context.Background(),"ip","addr\nshow");if err==nil{t.Fatal("expected rejection")}}
func TestRuntimeCommandRejectsNUL(t *testing.T){_,err:=(RuntimeCommand{}).Run(context.Background(),"ip","route\x00show");if err==nil{t.Fatal("expected NUL rejection")}}
