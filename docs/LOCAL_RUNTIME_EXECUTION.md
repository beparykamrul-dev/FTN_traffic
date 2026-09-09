# FTN local runtime execution

FTN_traffic owns orchestration and policy locally. FRR, BIRD, GoBGP, iproute2, nftables, edge caches, DNS, flow collectors and observability systems are replaceable local backends.

## Execution boundary

1. Desired state is validated.
2. Authorization and approval are required for mutations.
3. Backend health is checked before execution.
4. Route policy enforces RPKI and max-prefix constraints.
5. Reconciliation applies only validated desired state.
6. Failed health gates fail closed; rollback is handled by the migration controller.

GitHub is source/update distribution only. Runtime state, credentials, customer records and production secrets remain local.

This layer deliberately does not pretend to be a live FRR/BIRD/GoBGP installation: concrete adapters execute only when those services are installed and explicitly authorized in the target FTN environment.
