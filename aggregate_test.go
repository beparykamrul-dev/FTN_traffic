package traffic

import "testing"

func TestAdd(t *testing.T) { a := Add(Aggregate{}, Event{Bytes:10, Packets:2}); a = Add(a, Event{Bytes:5, Packets:1}); if a.Bytes != 15 || a.Packets != 3 || a.Events != 2 { t.Fatalf("unexpected aggregate: %+v", a) } }
