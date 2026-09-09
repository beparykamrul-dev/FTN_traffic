# FTN Dataplane Mutation Pipeline

Route mutations follow a fail-closed pipeline:

1. Normalize and canonicalize prefixes/next-hops.
2. Reject unauthorized route intents and duplicate routes.
3. Require explicit approval and request metadata.
4. Enforce RPKI and max-prefix policy.
5. Require healthy router, BGP, and configured BFD gates.
6. Apply or withdraw through the provider-neutral `RouterAdapter`.
7. Emit an audit event containing actor, request ID, approval ID, and result.

Telemetry is validated and sensitive labels are removed before publication. Runtime command paths remain allow-listed and argument-validated.

The pipeline does not capture customer payloads or perform customer traffic interception. Third-party private access remains authorization-gated.
