# FTN DNS Portal API contract

The control plane exposes read-oriented DNS status and provider-management contracts.

## Read endpoints

- `GET /api/v1/dns/portal` — engine, provider and Anycast-node snapshot.
- `GET /api/v1/dns/providers` — configured provider status.
- `GET /api/v1/dns/engines` — DNS engine status.
- `GET /api/v1/dns/nodes` — mesh/Anycast node health.

## Mutations

Zone/record changes and provider connect/re-authenticate/disable operations must use the existing authorization layer and approval workflow. The API must reject privileged mutations unless the provider is authorized and the requested operation has the required approval.

## Secrets

Requests use credential references. API keys, OAuth secrets, DNS TSIG keys and certificate private keys are never returned by the portal and are never committed to Git.
