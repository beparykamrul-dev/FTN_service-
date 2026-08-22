# FTN Service — Implementation Readiness

This file closes the feature-definition phase and defines the executable acceptance gate.

## Done in repository scope

- Aether-Core secure mesh architecture
- WireGuard / AmneziaWG production profiles; WGQ experimental profile
- GoBGP / FRR / SR-SRv6 routing boundaries
- eBPF/XDP defensive dataplane boundary
- mTLS / Private PKI / ACME / OCSP / CRL
- TPM 2.0 / HSM trust boundaries
- Global DNS Mesh including Hickory adapter
- Edge proxy and service-mesh proxy boundaries
- PostgreSQL / PgBouncer / replica architecture
- ClickHouse telemetry architecture
- Flow analytics normalization
- Consul / Nomad integration boundaries
- FTNVPN Android control-plane foundation
- React/WebGL NOC target

## Execution gate

No component is marked live merely because an architecture file exists. Each environment must pass build, integration, health, security and rollback checks before activation.

## Required production gates

1. Build and static checks
2. Unit/integration tests
3. Secrets and certificate provisioning outside source control
4. DNS and DNSSEC validation
5. BGP policy validation
6. Tunnel connectivity and failover validation
7. eBPF/XDP staging and rollback validation
8. PKI issuance/renewal/revocation validation
9. Database backup/restore and replica validation
10. ClickHouse ingestion/retention validation
11. Edge PoP health/failover validation
12. FTNVPN enrollment/profile/reconnect validation
13. Central telemetry/audit validation
14. Load and recovery testing
15. Human production acceptance

## Operational rule

Automatic actions are policy-gated. Emergency/manual actions require authentication and audit logging. FTN resilience features are for infrastructure operated or authorized by FTN and are not an external firewall/DPI bypass mechanism.

## Completion semantics

Feature scope is complete. The remaining status is deployment validation, because live infrastructure credentials, provider accounts, hardware and peer configuration cannot be fabricated inside source control.
