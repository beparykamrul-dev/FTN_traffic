package migration

import "testing"

func TestPlanValidation(t *testing.T) {
	p := Plan{Request:Request{ID:"m1",From:"edge-a",To:"edge-b",ApprovalID:"ap-1",Strategy:"weighted"},CanaryPercent:10,ObserveSeconds:30,RollbackOnHealthFailure:true}
	if err:=p.Validate(); err!=nil { t.Fatal(err) }
	p.CanaryPercent=0
	if err:=p.Validate(); err!=ErrInvalidMigrationPlan { t.Fatalf("got %v",err) }
}
