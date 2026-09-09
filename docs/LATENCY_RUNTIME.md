# FTN Latency Runtime

The runtime path layer exposes measured path state and Prometheus metrics. It is intentionally provider-neutral: FTN full mesh, authorized IX, authorized transit, authorized global edge, and authorized origin paths can all be represented by the same model.

## Runtime flow

`probe -> measurement -> eligibility -> cost -> hysteresis -> selected path -> migration controller`

Only authorized and healthy paths can be selected. A route or provider mutation still requires the existing approval boundary.

## Metrics

- `ftn_latency_selected`
- `ftn_path_rtt_p50_ms`
- `ftn_path_rtt_p95_ms`
- `ftn_path_jitter_ms`
- `ftn_path_loss_percent`
- `ftn_path_availability`

This layer does not claim to create physical peering or guarantee a fixed latency. Actual BDIX-like performance requires appropriately located FTN POPs, capacity, fiber paths, and authorized peering/transit.
