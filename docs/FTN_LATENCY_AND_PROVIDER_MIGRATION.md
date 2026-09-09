# FTN Latency-First and Provider Migration

## Objective

FTN should minimize end-user latency without making BDIX, one transit provider, one CDN, or one hosting provider a single point of dependency.

Latency is treated as a measured property of the available network paths. The controller prefers healthy FTN full-mesh paths, then authorized peering/transit/edge/origin paths according to measured cost and policy.

## Path hierarchy

```text
Customer
   |
Nearest healthy FTN POP
   |
FTN Full Mesh / ECMP
   |
+-- Authorized IX peer
+-- Authorized Internet transit
+-- Authorized global edge/CDN
+-- Authorized global hosting/origin
```

BDIX can remain one peering path, but it is not the only path.

## Latency controller

The controller measures RTT, jitter, packet loss, availability, DNS latency and HTTP health. It calculates a policy cost and selects the lowest-cost healthy path with hysteresis so that small fluctuations do not cause route flapping.

The controller must not infer access rights from latency or public routing data. Provider authorization remains a prerequisite for private provider resources.

## Hosting migration

FTN maintains an origin/hosting registry containing only non-secret references:

`service -> provider -> region/POP -> endpoint -> health policy -> authorization reference`

A service may move between an FTN origin and an authorized global host using:

`inventory -> provision -> sync -> health test -> canary -> traffic shift -> observe -> commit -> retire old origin`

The old origin remains available until the new path passes its health gate.

## Domain migration

Domain movement uses:

`inventory -> DNS synchronization -> DNSSEC/SOA/NS validation -> TTL reduction -> cutover -> health monitoring -> TTL restoration`

Provider credentials, API tokens, TSIG secrets and private keys remain in the runtime secret store.

## CDN/edge migration

FTN can use authorized CDN/edge providers as temporary or permanent delivery paths. Cache health, origin latency, error rate and cache-hit ratio are measured before and after migration.

A CDN is not treated as Internet transit.

## Rollback

Every privileged migration requires an idempotency key, audit record and approval. If the health gate fails, the controller returns to the last known-good desired state rather than attempting an unauthorized workaround.

## Low-latency engineering requirements

- Deploy POPs close to customer concentrations.
- Keep FTN-owned services cached at the edge where appropriate.
- Use ECMP and multiple physically diverse upstream paths.
- Use BFD for fast failure detection.
- Use IPv4 and IPv6 where end-to-end health is verified.
- Monitor MTU, packet loss, jitter and congestion, not RTT alone.
- Prefer direct authorized peering when it is genuinely lower-latency than transit.
- Use global hosting/edge regions when they provide a better measured path.
- Keep DNS authoritative service distributed and health-aware.

## Important boundary

No software configuration can guarantee BDIX-equivalent latency everywhere. Physical path length, fiber topology, peering, transit capacity and congestion determine actual network latency. FTN can, however, make BDIX one option within a multi-path architecture and continuously select the best authorized healthy path.
