package acquisition

type Mode string

const (
	DirectOrigin      Mode = "direct_origin"
	PublicCDN         Mode = "public_cdn"
	OpenDistribution  Mode = "open_distribution"
	P2PDistribution   Mode = "p2p_distribution"
	CDNContract       Mode = "cdn_contract"
	PartnerEdge       Mode = "partner_edge"
	IXPeering         Mode = "ix_peering"
	InternetTransit   Mode = "internet_transit"
	PublicTelemetry   Mode = "public_telemetry"
)

func (m Mode) String() string { return string(m) }
