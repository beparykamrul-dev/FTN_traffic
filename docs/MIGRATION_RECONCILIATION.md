# Traffic migration reconciliation

FTN_traffic treats traffic migration as desired-state reconciliation rather than an ad-hoc route command.

## Lifecycle

`preflight -> approval -> health_gate -> shift -> observe -> commit`

A failed shift or post-shift health gate enters rollback. Every migration requires a unique idempotency key, explicit approval, authorization and an auditable lifecycle.

Supported strategies are ECMP, weighted, canary, failover and blue/green. The controller passes the requested strategy to the execution backend instead of silently selecting one.

GitHub is not a runtime dependency. Production state and credentials remain local. Concrete route/DNS/edge mutations occur only through installed and authorized backends.
