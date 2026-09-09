package acquisition

type Source struct {
	Name                 string   `json:"name"`
	Mode                 Mode     `json:"mode"`
	PublicDistribution   bool     `json:"public_distribution"`
	RequiresAuthorization bool    `json:"requires_authorization"`
	CanDeliverContent    bool     `json:"can_deliver_content"`
	CanProvideTransit    bool     `json:"can_provide_transit"`
	CanProvideTelemetry  bool     `json:"can_provide_telemetry"`
	Capabilities         []string `json:"capabilities"`
}

var Sources = []Source{
	{Name: "direct_origin", Mode: DirectOrigin, CanDeliverContent: true, Capabilities: []string{"origin", "cache"}},
	{Name: "jsdelivr", Mode: PublicCDN, PublicDistribution: true, CanDeliverContent: true, Capabilities: []string{"cache", "delivery"}},
	{Name: "unpkg", Mode: PublicCDN, PublicDistribution: true, CanDeliverContent: true, Capabilities: []string{"cache", "delivery"}},
	{Name: "cdnjs", Mode: PublicCDN, PublicDistribution: true, CanDeliverContent: true, Capabilities: []string{"cache", "delivery"}},
	{Name: "github_pages", Mode: PublicCDN, PublicDistribution: true, CanDeliverContent: true, Capabilities: []string{"origin", "delivery"}},
	{Name: "webtorrent", Mode: P2PDistribution, RequiresAuthorization: true, CanDeliverContent: true, Capabilities: []string{"p2p", "delivery"}},
	{Name: "cdn_contract", Mode: CDNContract, RequiresAuthorization: true, CanDeliverContent: true, Capabilities: []string{"cache", "delivery"}},
	{Name: "partner_edge", Mode: PartnerEdge, RequiresAuthorization: true, CanDeliverContent: true, Capabilities: []string{"edge", "delivery"}},
	{Name: "ix_peering", Mode: IXPeering, RequiresAuthorization: true, CanProvideTransit: true, CanProvideTelemetry: true, Capabilities: []string{"bgp", "peering"}},
	{Name: "internet_transit", Mode: InternetTransit, RequiresAuthorization: true, CanProvideTransit: true, Capabilities: []string{"bgp", "transit"}},
	{Name: "public_telemetry", Mode: PublicTelemetry, PublicDistribution: true, CanProvideTelemetry: true, Capabilities: []string{"telemetry"}},
}
