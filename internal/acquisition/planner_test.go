package acquisition

import "testing"

func TestPublicCDNAllowsPublicDistribution(t *testing.T) {
	s, err := Plan(Request{Mode: PublicCDN, ThirdParty: false, PublicOrAuthorized: true})
	if err != nil || s.Name == "" {
		t.Fatalf("expected public CDN plan, got source=%+v err=%v", s, err)
	}
}

func TestTransitRequiresAuthorization(t *testing.T) {
	if _, err := Plan(Request{Mode: InternetTransit, ThirdParty: true}); err == nil {
		t.Fatal("expected authorization failure")
	}
}

func TestAuthorizedTransitAllowed(t *testing.T) {
	s, err := Plan(Request{Mode: InternetTransit, ThirdParty: true, PublicOrAuthorized: true})
	if err != nil || s.Mode != InternetTransit {
		t.Fatalf("expected authorized transit, got source=%+v err=%v", s, err)
	}
}
