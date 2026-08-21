# FTN Service — Project Closure & Production Readiness

## Scope frozen

The FTN platform scope is now frozen around these production domains:

- FTN Global DNS Mesh and DNS adapter framework
- FTN.OS Router Control Plane
- Aether-Core provider-neutral backbone orchestration
- WireGuard and AmneziaWG production profiles; WGQ remains experimental
- GoBGP/FRR routing integration boundary
- eBPF/XDP defensive policy boundary
- Global Edge/PoP architecture
- Sidecar, mTLS, Edge and Database proxy boundaries
- PostgreSQL authoritative state + PgBouncer + replica/read-scaling architecture
- ClickHouse high-volume telemetry
- Private PKI, ACME, OCSP/CRL, TPM 2.0 and HSM boundaries
- Flow telemetry adapters: YAF, SiLK, Nfdump/NfSen, pmacct, ElastFlow and compatible sources
- React/WebGL NOC visualization boundary
- FTNVPN Android client foundation
- Central audit, inventory, health and route-decision tracking

## Operational safety

AI may observe, correlate, score and recommend. State-changing operations are policy-gated and approval-controlled unless an explicitly approved deterministic automation policy permits them.

Traffic-resilience features are for FTN-owned or authorized infrastructure. The platform does not implement mechanisms intended to bypass third-party firewalls, DPI controls or provider restrictions.

## Production completion criteria

The repository is architecture/consolidation complete. Before declaring a deployment environment production-ready, run the environment-specific implementation, integration, security and load test suites and verify certificates, routing, DNS, database replication, backups, monitoring and rollback procedures.

## Next phase

No additional technology catalog expansion is required. Remaining work is deployment-specific implementation, validation, performance testing, security review and operational rollout.
