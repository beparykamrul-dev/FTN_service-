# FTN Global Provider / Open-Source Integration Plan

## Goal

Bring FTN traffic into FTN-owned edge/backbone infrastructure through legitimate provider connectivity, using open-source routing, proxy and observability components where their licenses and provider terms permit.

## Provider-neutral edge model

```text
Provider VPS / Edge PoP
  -> Edge health checks
  -> Envoy / Traefik / NGINX
  -> FRRouting / GoBGP
  -> WireGuard / AmneziaWG
  -> Aether-Core
  -> FTN.OS router
  -> FTN services
```

FRRouting is used as the standards-based routing engine. It supports BGP and other routing protocols and can install routing decisions into the host kernel. See the upstream project documentation before enabling a protocol or feature. 

## Open-source components

- FRRouting: BGP/OSPF/IS-IS/EVPN/VRF and related routing capabilities
- Envoy: edge/service proxy
- Traefik: ingress/reverse proxy
- NGINX: edge/reverse/caching proxy
- WireGuard: encrypted transport
- OpenTelemetry Collector: vendor-neutral telemetry collection
- ClickHouse: high-volume telemetry analytics
- PostgreSQL/PgBouncer: authoritative state and connection pooling

OpenTelemetry is designed for generating, collecting and exporting telemetry; ClickHouse documents the OpenTelemetry Collector as an ingestion component for ClickHouse-based observability. 

## Provider onboarding contract

For every hosting/transit provider, the control plane records:

- provider and PoP identifier
- region/ASN
- public prefixes actually assigned to FTN
- IPv4/IPv6 availability
- bandwidth and traffic limits
- BGP capability and authorization status
- tunnel endpoint and health status
- provider API capability
- abuse/contact information
- monthly/usage cost metadata
- deployment status

No provider is treated as connected merely because its name is present in configuration.

## Routing safety

BGP advertisements are disabled by default. A prefix can be advertised only after authorization, prefix validation, RPKI/IRR policy checks where applicable, peer validation, and an explicit production activation step. Provider terms and routing policies must be respected.

## Traffic steering

Aether-Core chooses among healthy, authorized paths using measured latency, loss, capacity and policy. Failed paths are withdrawn from the active set and restored only after health recovery. This is resilience/load distribution, not a mechanism for bypassing third-party security controls.

## License and supply-chain policy

Only upstream repositories with compatible licenses may be incorporated or packaged. Dependencies remain pinned and are subject to checksum/signature verification and vulnerability scanning. Forking a project does not transfer ownership or grant permission to redistribute code outside its license.

## Reality check

Open-source software cannot by itself create Internet transit capacity. Actual global traffic ingress requires real provider-hosted PoPs, addresses/prefixes, transit/peering arrangements and provider authorization. Those credentials and contracts must be supplied in the target deployment environment and must never be committed to Git.
