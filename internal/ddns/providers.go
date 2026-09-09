package ddns

// ProviderProfile describes an integration without embedding credentials.
type ProviderProfile struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	CredentialEnv []string `json:"credential_env"`
	SupportsIPv4 bool `json:"supports_ipv4"`
	SupportsIPv6 bool `json:"supports_ipv6"`
}

var ProviderProfiles = []ProviderProfile{
	{Name: "duckdns", Kind: "dynamic_dns", CredentialEnv: []string{"FTN_DDNS_DUCKDNS_TOKEN"}, SupportsIPv4: true, SupportsIPv6: true},
	{Name: "porkbun", Kind: "dns_api", CredentialEnv: []string{"FTN_DDNS_PORKBUN_API_KEY", "FTN_DDNS_PORKBUN_SECRET_API_KEY"}, SupportsIPv4: true, SupportsIPv6: true},
	{Name: "caddy_dns", Kind: "dns_provider_plugin", CredentialEnv: []string{"FTN_DDNS_CADDY_CREDENTIAL_REF"}, SupportsIPv4: true, SupportsIPv6: true},
	{Name: "route53_aws", Kind: "authoritative_dns", CredentialEnv: []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_REGION"}, SupportsIPv4: true, SupportsIPv6: true},
}
