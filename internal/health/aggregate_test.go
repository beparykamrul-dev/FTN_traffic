package health

import "testing"

func TestAll(t *testing.T) {
	if got := All(OK("a"), OK("b")); !got.Healthy { t.Fatalf("expected healthy: %+v", got) }
	got := All(OK("a"), Status{Name: "b", Healthy: false, Detail: "down"})
	if got.Healthy || got.Detail != "b: down" { t.Fatalf("unexpected result: %+v", got) }
}
