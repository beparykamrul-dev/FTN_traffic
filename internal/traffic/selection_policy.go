package traffic

import "strings"

type AuthorizationClass string

const (
	OwnedOrAuthorized AuthorizationClass = "owned_or_authorized"
	PublicDistribution AuthorizationClass = "public"
	ContractRequired AuthorizationClass = "contract_required"
	PartnerRequired AuthorizationClass = "partner_required"
	PeerRequired AuthorizationClass = "peer_required"
	TransitContractRequired AuthorizationClass = "transit_contract_required"
)

type Source struct {
	ID string
	Authorization AuthorizationClass
	Enabled bool
}

func AuthorizedSource(s Source) bool {
	if !s.Enabled || strings.TrimSpace(s.ID) == "" { return false }
	switch s.Authorization {
	case OwnedOrAuthorized, PublicDistribution, ContractRequired, PartnerRequired, PeerRequired, TransitContractRequired:
		return true
	default:
		return false
	}
}
