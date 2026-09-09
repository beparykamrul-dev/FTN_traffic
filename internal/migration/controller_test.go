package migration
import "testing"
func TestMigrationRetryAfterShiftRollback(t *testing.T){e:=&fakeExec{failShift:true};c:=NewController(e,fakeApproval{true},fakeHealth{true});r:=req("retry");if err:=c.Run(nil,r);err==nil{t.Fatal("expected first failure")};e.failShift=false;if err:=c.Run(nil,r);err!=nil{t.Fatalf("retry failed: %v",err)}}
func TestMigrationRejectsDifferentRequestForActiveKey(t *testing.T){e:=&fakeExec{};c:=NewController(e,fakeApproval{true},fakeHealth{true});r:=req("same");if err:=c.Run(nil,r);err!=nil{t.Fatal(err)};r.To="other";if err:=c.Run(nil,r);err!=ErrMigrationFingerprint{t.Fatalf("got %v",err)}}
func TestIdempotencyLegacyClaim(t *testing.T){i:=NewIdempotency();if !i.Claim("legacy"){t.Fatal("first claim failed")};if i.Claim("legacy"){t.Fatal("duplicate claim accepted")};if _,ok:=i.State("legacy");!ok{t.Fatal("state missing")}}
