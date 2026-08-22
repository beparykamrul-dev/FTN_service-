# FTN Workflow Component Matrix

| Workflow stage | Component | Completion contract |
|---|---|---|
| Source/CI | Go, Android | Build/test gate |
| Security | PKI, mTLS, ACME, TPM/HSM | Identity + revocation gate |
| DNS | PowerDNS, Technitium, CoreDNS, Unbound, dnsdist, GoDNS, Hickory | Health + DNSSEC/config gate |
| Edge | Envoy, Traefik, NGINX | Health/failover gate |
| Backbone | Aether-Core, WireGuard, AmneziaWG | Path/health gate |
| Routing | GoBGP, FRR, SR/SRv6 | Policy validation gate |
| Dataplane | eBPF/XDP | Staged load/rollback gate |
| Services | Go + ASP.NET Core | API/integration gate |
| Database | PostgreSQL, PgBouncer, replicas | Backup/restore/lag gate |
| Analytics | ClickHouse, SiLK/YAF/Nfdump/pmacct/ELK | Ingestion/retention gate |
| NOC | React/WebGL | Live telemetry/UI gate |
| Client | FTNVPN Android | Enrollment/profile/reconnect gate |
| Orchestration | Consul/Nomad | Discovery/scheduling gate |
| Release | GitHub Actions | Promotion gate |

## Operational rule

A workflow stage is not marked production-ready merely because its configuration exists. It becomes ready only after its implementation, health check, test evidence and rollback procedure pass in the target environment.

## Final state

Feature scope remains frozen. This matrix is the final mapping between the previously defined FTN components and the release workflow.
