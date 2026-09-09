package providers

// CatalogEntry describes an integration target without implying ownership
// of, or access to, third-party traffic. Access must come from an authorized
// CDN contract, peering session, transit service, or public distribution source.
type CatalogEntry struct {
	Name         string   `json:"name"`
	Role         string   `json:"role"`
	Capabilities []string `json:"capabilities"`
	Mode         string   `json:"mode"`
}

var GlobalCatalog = []CatalogEntry{
	{Name: "facebook", Role: "content_platform", Capabilities: []string{"telemetry", "edge"}, Mode: "authorized_only"},
	{Name: "google", Role: "content_platform", Capabilities: []string{"telemetry", "edge"}, Mode: "authorized_only"},
	{Name: "netflix", Role: "content_platform", Capabilities: []string{"telemetry", "edge"}, Mode: "authorized_only"},
	{Name: "tiktok", Role: "content_platform", Capabilities: []string{"telemetry", "edge"}, Mode: "authorized_only"},
	{Name: "imo", Role: "content_platform", Capabilities: []string{"telemetry"}, Mode: "authorized_only"},
	{Name: "pubg", Role: "game_platform", Capabilities: []string{"telemetry"}, Mode: "authorized_only"},
	{Name: "freefire", Role: "game_platform", Capabilities: []string{"telemetry"}, Mode: "authorized_only"},
	{Name: "edgenext", Role: "cdn", Capabilities: []string{"cache", "delivery", "health"}, Mode: "contract_or_public"},
	{Name: "akamai", Role: "cdn", Capabilities: []string{"cache", "delivery", "health"}, Mode: "contract_or_public"},
	{Name: "aws", Role: "cloud_cdn", Capabilities: []string{"cache", "delivery", "health"}, Mode: "contract_or_public"},
	{Name: "cloudflare", Role: "cdn", Capabilities: []string{"cache", "delivery", "health", "flow"}, Mode: "contract_or_public"},
	{Name: "fastly", Role: "cdn", Capabilities: []string{"cache", "delivery", "health"}, Mode: "contract_or_public"},
	{Name: "bunny", Role: "cdn", Capabilities: []string{"cache", "delivery", "health"}, Mode: "contract_or_public"},
	{Name: "tencent_cloud", Role: "cloud_cdn", Capabilities: []string{"cache", "delivery", "health"}, Mode: "contract_or_public"},

	// Public distribution intermediaries. These distribute public/authorized
	// artifacts; they are not Internet-transit providers and cannot be used to
	// obtain arbitrary private third-party traffic.
	{Name: "jsdelivr", Role: "public_distribution", Capabilities: []string{"cache", "delivery"}, Mode: "public_only"},
	{Name: "unpkg", Role: "public_distribution", Capabilities: []string{"cache", "delivery"}, Mode: "public_only"},
	{Name: "cdnjs", Role: "public_distribution", Capabilities: []string{"cache", "delivery"}, Mode: "public_only"},
	{Name: "github_pages", Role: "public_distribution", Capabilities: []string{"origin", "delivery"}, Mode: "public_only"},
	{Name: "npm_registry", Role: "public_distribution", Capabilities: []string{"artifact", "delivery"}, Mode: "public_only"},
	{Name: "webtorrent", Role: "open_distribution", Capabilities: []string{"p2p", "delivery"}, Mode: "owned_or_authorized_only"},
}
