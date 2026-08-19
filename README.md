# FTN Service

Production-oriented service foundation for Family Time Network (FTN).

## Unified platform

FTN operates many physical servers as one logical global service cluster. The service fabric does **not** depend on a local fiber backbone; local fiber can be an optional transport path, not a control-plane requirement.

## Foundation

- 001 — Frontier Registry & Unified Telemetry
- 002 — Physics Data Ingestion & Source Adapter Engine
- 003 — Central Telemetry Bus, Metrics/Event retention and API tracking
- Central PostgreSQL source of truth + PgBouncer
- Central audit/tracking
- Multi-server cluster registry and geographic topology
- Aether-Core transport abstraction: WireGuard / experimental WGQ / AmneziaWG
- Global DNS mesh and DNS portal architecture
- ASP.NET Core portal/control-plane integration contract
- Go service backend contract

## DNS platform

- PowerDNS Enterprise — authoritative
- Technitium DNS Enterprise
- CoreDNS
- Unbound
- dnsdist
- GoDNS
- Anycast DNS
- DNSPod / Tencent Cloud
- Cloudflare DNS
- Akamai DNS
- Global DNS Mesh

Current infrastructure inventory includes `ns1.familytimenet.com` at `100.52.87.3` / `138.252.113.19` and `ns2.familytimenet.com` at `100.52.86.20` / `138.252.113.2`, as supplied for FTN deployment inventory.

## Service integrations

- Cloudflare
- Netflix
- Fastly

Provider APIs and credentials remain adapter-driven and must be configured through runtime secrets.

## Central monitoring

- Server health and capacity
- DNS authoritative/recursive health
- Anycast and propagation health
- Transport health
- Latency and packet loss
- Geographic NOC map
- Incidents and alerts
- Approval queue
- Full change/audit tracking

## HTTP foundation

- `GET /healthz`
- `GET /readyz`
- `GET /api/v1/frontier/modules`
- `GET /api/v1/telemetry/events`
- `GET /api/v1/tracking/events`

## Database

`migrations/002_platform_mesh.sql` adds centralized platform servers, provider inventory, DNS endpoints/health, topology links and approval-controlled intents. PostgreSQL remains the central source of truth; schema changes are forward-only migrations.

## Governance

State-changing automation is approval-first and auditable. Hypothetical physics concepts remain research/simulation classifications rather than being represented as real communication transports.
