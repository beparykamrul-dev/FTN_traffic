# Provider Authorization Dashboard

Each provider is managed independently. The dashboard exposes authorization state and operational capabilities without storing secrets in Git.

## Provider card

- Provider name and category
- Authorization status: pending, connected, expired, disabled
- Method: API, OAuth, contract, or peering
- Account/reference identifier
- Granted scopes/capabilities
- Expiration and last verification time
- Credential reference (secret-store reference only)
- Health and traffic summary

## Actions

`Connect` -> establish the provider's supported authorization flow.

`Re-authenticate` -> refresh an expired or invalid authorization.

`Disable` -> stop using the provider without deleting its configuration.

`Test` -> verify authorization and provider health.

`Configure` -> manage provider-specific routing, edge, cache, telemetry, and delivery settings allowed by the provider.

Privileged route or edge mutations remain approval-gated. Provider secrets are references to an external secret store, never plaintext repository configuration.
