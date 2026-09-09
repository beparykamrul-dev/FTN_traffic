# FTN DDNS Provider Integrations

## DuckDNS

The integration is represented as a credential-backed dynamic DNS provider. The FTN control plane stores only a credential reference/environment binding; the DuckDNS token itself must remain in the deployment secret store.

## Porkbun

The integration uses a credential-backed DNS API profile. API key and secret are deployment secrets and are never committed to Git or included in audit payloads.

## Caddy DNS

Caddy DNS is represented as a provider-plugin integration. Caddy/provider credentials are supplied through the deployment environment or secret manager, not repository configuration.

## Common execution policy

1. Validate the DDNS update and IP address.
2. Require an explicit approval identifier for mutation.
3. Require active provider authorization before remote execution.
4. Apply A or AAAA according to the address family.
5. Record success/failure without logging credentials.
6. Keep provider fan-out idempotent and fail closed on missing configuration.

These integrations are provider adapters for zones FTN owns or is authorized to manage; they do not provide access to unrelated third-party zones.
