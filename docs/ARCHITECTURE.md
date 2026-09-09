# FTN Traffic Architecture

## Planes

### CDN/Edge plane
1. FTN public repository/release is the versioned source.
2. CI validates and publishes immutable assets.
3. CDN/edge providers pull/cache public assets.
4. FTN tracks cache health, latency, egress and availability.

### Network plane
1. Authorized transit providers and IX peers terminate on FTN routers.
2. BGP distributes only authorized prefixes.
3. Routing policy selects transit/peering paths based on policy and health.
4. NetFlow/IPFIX and router telemetry feed the traffic analytics layer.

## Integration

The CDN and network planes share identity, telemetry, health and policy, but remain technically independent. A CDN provider cannot be treated as an Internet transit provider.

## Future FTN POP

```text
                 FTN Control Plane
                        |
              +---------+---------+
              |                   |
          BGP/Transit          CDN Control
              |                   |
          FTN Core            FTN Origin
              |                   |
           FTN POP ----------- Edge Cache
              |
           Customers
```

## Provider policy

- Capability and health based selection
- Credentials from runtime secret store only
- Privileged mutations require explicit approval
- Every provider mutation is audited
- Fail closed when authorization or health validation fails
