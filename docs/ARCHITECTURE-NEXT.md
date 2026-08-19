# FTN Unified Global Service Architecture

## Non-negotiable topology rule

FTN must not require a local fiber backbone as a service dependency. Physical fiber may exist as an optional access/transport path, but the logical service fabric is WAN/global and provider-neutral.

## Planes

```text
                ASP.NET Core Portal
                 Unified Control UI
                         |
                    Go Backend
                         |
                 Central Event Bus
                         |
                  PostgreSQL + PgBouncer
                         |
       +-----------------+------------------+
       |                 |                  |
  Server Fleet       DNS Mesh          Aether-Core
  20–30+ nodes       Anycast           WG/WGQ/AWG
       |                 |                  |
       +-----------------+------------------+
                         |
                Central Monitoring
                         |
                 Geographic Map/NOC
```

## DNS mesh

PowerDNS Enterprise, Technitium, CoreDNS, Unbound, dnsdist and GoDNS are controlled through provider adapters. DNSPod/Tencent Cloud, Cloudflare and Akamai are external-provider adapters. Anycast and propagation monitoring are modeled as platform capabilities rather than tied to one vendor.

## CDN/cache integrations

Cloudflare, Netflix and Fastly are service/provider inventory integrations. Their APIs, credentials and exact capabilities must be verified before an adapter performs changes.

## Map/NOC

The map is a topology view over the centralized server registry and DNS/transport health. It must not assume a particular map vendor. Nodes expose state, capacity, region and last-seen data; links expose transport, latency and loss.

## Central tracking

Every provider operation, server state transition, DNS change, topology event and approval is represented as a trackable event with actor, target, operation, status and timestamps.

## Security boundary

No credentials, private keys, DNS signing keys or provider API tokens are committed to the repository. Runtime secret management is mandatory.
