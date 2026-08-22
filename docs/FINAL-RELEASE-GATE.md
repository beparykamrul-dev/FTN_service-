# FTN Service — Final Release Gate

## Release workflow

```text
Source
  -> CI
  -> Unit Tests
  -> Integration Tests
  -> Security/Secret Scan
  -> Staging
  -> DNS/PKI/BGP/Tunnel Validation
  -> Edge + Database Failover Tests
  -> FTNVPN Validation
  -> Observability Verification
  -> Production Approval
```

## Runtime control loop

```text
Observe -> Normalize -> Correlate -> Health Score -> Policy Check
        -> Authorized Change -> Verify -> Audit -> Rollback if unhealthy
```

## Release invariants

- No production credentials or private keys in Git.
- BGP advertisements require validated policy and explicit production authorization.
- DNS changes are versioned and reversible.
- PKI changes are audited and revocable.
- eBPF/XDP changes have a rollback path.
- Database writes use the authoritative primary; replicas are health/lag checked.
- Telemetry is observable before automated policy is enabled.
- FTNVPN profiles are issued only to enrolled devices.
- Manual emergency actions are authenticated and audited.

## Definition of done

The feature scope is frozen. A release is complete when the implementation passes this gate in the target environment. Architecture documentation alone is not treated as proof that an external provider, server, BGP peer, DNS service, HSM/TPM or Android tunnel is live.
