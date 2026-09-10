# FTN Global Traffic Runtime

The runtime path is authorization-first and health-first. A candidate is eligible only when its path is authorized, healthy, within packet-loss/latency/availability gates, and represented by a supported path class.

Selection uses weighted latency, loss, jitter and availability cost, with capacity/locality adjustments and deterministic path-ID tie breaking. Hysteresis reduces unnecessary path churn.

Migration is separate from selection. A selected-path change enters the migration controller, which requires an approval reference, preflight, health gate, gradual shift, observation and commit. Shift or post-shift failures trigger rollback; idempotency prevents accidental duplicate execution.

The design supports FTN full-mesh paths, authorized IX peers, authorized transit, authorized global edge providers and authorized origin hosts. Public routing information is treated as telemetry only and never as permission to access a third party.

Observability can remain local in Prometheus or use the configured remote-write boundary. Telemetry is redacted and policy-validated before external publication. No customer payloads, credentials or secrets are part of the runtime metric model.
