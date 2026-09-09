package dataplane

import("context";"testing")
func TestLocalExecutorAllowlist(t *testing.T){e:=LocalExecutor{Allowed:map[string]bool{"true":true}};if _,err:=e.Run(context.Background(),"false");err!=ErrCommandNotAllowed{t.Fatalf("got %v",err)}}
