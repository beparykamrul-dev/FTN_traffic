# FTN as an ISP Upstream Provider

## Target operating model

FTN is designed to become a network provider that downstream ISPs can use as an upstream for authorized Internet connectivity, IPv4/IPv6 transit, peering, DNS, caching and edge services.

```text
                    Global Internet
                         |
             +-----------+-----------+
             |                       |
      Authorized transit       Authorized IX peers
             |                       |
             +-----------+-----------+
                         |
                    FTN Core / AS
                  Full Mesh + ECMP
                         |
        +----------------+----------------+
        |                |                |
      ISP A            ISP B            ISP C
        |                |                |
    Customers        Customers        Customers
```

## What makes FTN an upstream

FTN must operate an independently routed network with its own authorized ASN/prefix resources, border routers, redundant POPs, BGP policy, RPKI/IRR hygiene, route-server/peering fabric, capacity planning and NOC operations.

A downstream ISP receives a contractual BGP service from FTN. The service can expose:

- IPv4 transit
- IPv6 transit
- default route
- partial routes
- full Internet table where capacity and policy permit
- FTN-owned routes and services
- authorized private peering
- DNS resolver/authoritative service
- edge/cache/CDN services

## Independence from a single national upstream

The design goal is not to secretly bypass another carrier or exchange. Instead FTN builds multiple legitimate external paths and makes them internal inputs to the FTN network. This allows downstream ISPs to see FTN as their provider while FTN maintains resilient connectivity through authorized transit, IX peering and FTN-operated peering infrastructure.

During the early phase FTN may still require external transit for global reachability. A software repository cannot remove that physical/commercial dependency. The long-term architecture reduces single-provider dependency through diverse authorized paths and capacity expansion.

## Routing policy

Every downstream session is validated for:

- authorized local/remote ASN
- IPv4/IPv6 address family
- RPKI validation
- IRR/route-policy validation where applicable
- max-prefix protection
- BFD where supported
- graceful restart
- communities and local preference policy
- explicit import/export policy

No public routing observation is treated as permission to originate or transit a prefix.

## FTN peering fabric

FTN can operate an authorized route-server/peering fabric for participating networks. Participants remain isolated by explicit import/export policy. Route-server functionality does not replace bilateral authorization, RPKI/IRR validation or contractual terms.

## Capacity and service tiers

The control plane should model downstream service tiers independently from physical links:

`customer ISP -> service profile -> BGP session -> FTN POP -> FTN core -> authorized external path`

Capacity admission must account for peak traffic, burst headroom, failure of the largest path, IPv4/IPv6 mix, packet-per-second limits and DDoS mitigation capacity.

## Fail-closed requirements

Credentials, BGP passwords, private keys and customer records never enter Git. Route mutations require approval and audit. Unauthorized route origination, unauthorized transit and customer traffic interception are prohibited.

## Reality constraint

FTN can become an upstream provider, but it cannot manufacture Internet reachability from software alone. To provide complete global Internet access, FTN needs its own address/ASN resources, physical POP connectivity and sufficient authorized interconnection capacity. The architecture therefore treats external transit and IX connectivity as replaceable upstream inputs, not as the FTN service itself.
