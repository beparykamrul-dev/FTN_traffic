# Local routing runtime

FTN_traffic now has a local execution boundary for FRR, BIRD and GoBGP, with kernel nftables/tc command boundaries.

- No shell interpolation: commands are invoked directly.
- Backend binaries are allowlisted.
- Health/status can be checked locally.
- Route, firewall and QoS mutations require both authorization and approval.
- RPKI/max-prefix policy remains mandatory at the route-policy layer.
- GitHub is not a runtime dependency.

The runtime layer does not embed credentials or silently mutate production infrastructure. A target host must explicitly install and authorize the selected backend.
