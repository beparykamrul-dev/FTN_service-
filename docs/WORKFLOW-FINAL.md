# FTN Service — Final Workflow

```text
Device / User
    -> Global DNS Mesh
    -> Healthy Edge PoP
    -> Edge Proxy / Security
    -> Aether-Core path selection
    -> FTN.OS Router Control Plane
    -> GoBGP / FRR / SR policy
    -> WireGuard / AmneziaWG transport
    -> eBPF/XDP defensive dataplane
    -> Go / ASP.NET services
    -> PostgreSQL/PgBouncer + ClickHouse telemetry
    -> React/WebGL NOC + FTNVPN telemetry
```

## Control loop

```text
Observe -> Normalize -> Correlate -> Health Score -> Policy Validate
        -> Change Proposal -> Apply (if authorized) -> Verify -> Rollback on failure
```

## Completion gate

1. Build all implementation packages.
2. Run unit and integration tests.
3. Validate DNS, PKI, routing, tunnel and database adapters in staging.
4. Validate edge failover and telemetry ingestion.
5. Validate FTNVPN enrollment/profile lifecycle.
6. Run security, recovery and load tests.
7. Promote only validated artifacts to production.

No production secret, BGP credential, private key or provider credential belongs in Git.

Feature scope is frozen; remaining work is executable validation/deployment for the actual environment.
