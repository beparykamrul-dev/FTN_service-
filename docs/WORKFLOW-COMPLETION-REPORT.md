# FTN Service — Workflow Completion Report

## Completed in repository

- Final workflow and release gates
- Component-to-workflow matrix
- CI validation workflow
- Global DNS / Edge / Aether-Core architecture
- FTN.OS router control-plane contract
- Proxy and database orchestration boundaries
- Hickory DNS integration boundary
- FTNVPN Android foundation
- Security/PKI and telemetry boundaries

## Remaining execution-only gates

These require the actual target environment and cannot truthfully be completed by repository documentation alone:

1. Build every package in CI.
2. Execute unit and integration tests.
3. Provision real PKI/ACME credentials and validate renewal/revocation.
4. Validate real DNS authoritative/recursive nodes and DNSSEC.
5. Validate real BGP peers and route policies before advertisement.
6. Deploy and test WireGuard/AmneziaWG paths.
7. Load and rollback-test eBPF/XDP programs.
8. Test PostgreSQL backup/restore, replication and PgBouncer failover.
9. Verify ClickHouse ingestion, retention and query limits.
10. Test every Edge PoP failover path.
11. Build/sign/test FTNVPN Android against real control-plane endpoints.
12. Run security, recovery and load tests.
13. Promote only validated artifacts to production.

## Rule

No placeholder is presented as a live production service. Secrets, private keys and provider credentials stay outside Git. Feature scope is frozen; only environment-specific execution and acceptance remain.
