# FTN Service — Production Closure

## Scope

This repository now defines the production implementation boundaries for the FTN unified network/control platform:

- Aether-Core backbone orchestration
- FTN.OS router control-plane contract
- WireGuard and AmneziaWG production transport profiles; WGQ remains experimental
- GoBGP/FRR routing integration and SR/SRv6 policy boundary
- eBPF/XDP defensive dataplane boundary
- mTLS, Private PKI, ACME, OCSP/CRL, TPM 2.0 and HSM boundaries
- Global DNS mesh including PowerDNS, Technitium, CoreDNS, Unbound, dnsdist, GoDNS and Hickory adapters
- Global Edge PoP / proxy architecture using policy-selected Envoy, Traefik or NGINX
- PostgreSQL authoritative state with PgBouncer and master/replica read scaling
- ClickHouse high-volume telemetry
- SiLK/YAF/Nfdump/NfSen/pmacct/ELK/ElastFlow and compatible flow analytics boundaries
- React/WebGL NOC visualization
- FTNVPN Android client foundation and control-plane lifecycle
- Consul/Nomad integration boundaries
- Lua sandbox and AI-assisted observability/policy recommendations

## Production truth

Architecture files and integration boundaries are committed. A component is considered production-ready only after its environment-specific adapter, credentials, tests, health checks and deployment manifest have passed validation. This document does not claim that every third-party system is installed or that every edge/PoP is live.

## Change-control policy

State-changing operations must be policy validated and auditable. Automatic routing, firewall, DNS, certificate and deployment actions require explicit enablement per environment. Manual emergency controls must be authenticated and logged.

## Security boundary

FTN provides encryption, redundancy, failover, authenticated routing and defensive filtering for infrastructure it operates or is authorized to administer. No component is designed to defeat third-party security controls or provide firewall/DPI evasion.

## Final acceptance checklist

- [ ] Build and unit tests pass for each implementation package
- [ ] Integration tests pass against staging DNS, routing, PKI and database services
- [ ] WireGuard/AmneziaWG profiles pass connectivity and failover tests
- [ ] BGP policies are validated before route activation
- [ ] eBPF/XDP programs are staged with rollback
- [ ] PKI/ACME renewal and revocation are tested
- [ ] PostgreSQL backup/restore and replica failover are tested
- [ ] ClickHouse retention and ingestion limits are validated
- [ ] DNS zone transfer/DNSSEC/health checks are validated
- [ ] FTNVPN enrollment, profile provisioning and reconnect are tested
- [ ] Edge PoP health and failover are validated
- [ ] Central audit and telemetry are verified
- [ ] Production secrets are provisioned outside source control

## Completion state

Feature scope is frozen. Remaining work is environment-specific build, deployment, validation and operational acceptance—not additional technology accumulation.
