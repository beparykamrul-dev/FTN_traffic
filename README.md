# FTN Traffic

FTN Traffic is the traffic/edge integration plane for Family Time Network.

## Scope

- GitHub-origin public asset distribution
- CDN/edge integration contracts
- Traffic telemetry and health signals
- Transit/IX/BGP traffic kept as a separate network plane
- Provider-neutral adapters
- Approval-first privileged changes
- Secrets stay outside Git

## Architecture

```text
GitHub / FTN Origin -> CDN/Edge -> Global Users

Transit / IX / BGP -> FTN Core -> FTN POPs -> FTN Clients
```

These planes are integrated for observability and policy, but CDN delivery does not substitute for Internet transit.

## Safety

Do not commit provider credentials, API tokens, private keys, customer data, or routing secrets.
