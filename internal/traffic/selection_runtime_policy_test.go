package traffic

import (
	"testing"
	"github.com/beparykamrul-dev/FTN_traffic/internal/latency"
)

func TestRuntimePolicyRejectsUnauthorized(t *testing.T) {
	p := RuntimeSelectionPolicy{RequireAuthorization:true, RequireHealthy:true, AllowProviderFailover:true}
	if err := p.Validate(Candidate{Path:latency.Path{ID:"x",Authorized:false,Healthy:true}}); err != ErrRuntimeSelectionPolicy { t.Fatalf("got %v", err) }
}

func TestRuntimePolicyAcceptsAuthorizedHealthy(t *testing.T) {
	p := RuntimeSelectionPolicy{RequireAuthorization:true, RequireHealthy:true, AllowProviderFailover:false}
	if err := p.Validate(Candidate{Path:latency.Path{ID:"x",Authorized:true,Healthy:true,Class:latency.PathFTNFullMesh}}); err != nil { t.Fatal(err) }
}
