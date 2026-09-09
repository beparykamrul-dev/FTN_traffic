package metrics

import "testing"

func TestRecord(t *testing.T) { var s Snapshot; s.Record(100, false); s.Record(20, true); if s.Requests != 2 || s.Bytes != 120 || s.Errors != 1 { t.Fatalf("unexpected snapshot: %+v", s) } }
