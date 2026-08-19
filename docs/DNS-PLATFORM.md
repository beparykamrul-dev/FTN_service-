# FTN Global DNS Platform

FTN DNS is designed as a provider-neutral global DNS mesh. No local fiber backbone is a required dependency for service continuity.

## Control layers

```text
ASP.NET Core Portal / Control Plane
              |
        Go Service Backend
              |
       Central PostgreSQL
              |
       DNS Control Engine
              |
  +-----------+-------------------+
  |           |                   |
Authoritative Recursive      Global Providers
  |           |                   |
PowerDNS    Unbound        DNSPod / Cloudflare / Akamai
Technitium  CoreDNS
            dnsdist
            GoDNS
```

## Registered infrastructure

### NS1 / authoritative backend

- PowerDNS Enterprise: `100.52.87.3`, `138.252.113.19`
- DNS name: `ns1.familytimenet.com`

### NS2 / frontend/backend service endpoint

- `100.52.86.20`, `138.252.113.2`
- DNS name: `ns2.familytimenet.com`

These values are treated as infrastructure inventory, not credentials.

## DNS roles

- PowerDNS Enterprise — authoritative DNS
- Technitium DNS Enterprise — authoritative/recursive platform adapter
- CoreDNS — service/discovery DNS adapter
- Unbound — recursive resolver
- dnsdist — DNS traffic distribution and health routing
- GoDNS — FTN DNS orchestration/control adapter
- Anycast DNS — global endpoint routing abstraction
- DNSPod / Tencent Cloud — external DNS provider adapter
- Cloudflare DNS — external DNS provider adapter
- Akamai DNS — external DNS provider adapter
- Global DNS Mesh — unified health, routing and propagation layer

## Portal controls

The portal should expose provider health, zones, records, nameserver status, resolver status, propagation state, Anycast state, latency, packet loss, capacity, incidents, change history and approval queues.

State-changing DNS operations must pass policy and approval checks and produce a central audit event.

## Central data model

PostgreSQL remains the source of truth for configuration metadata, provider inventory, DNS health observations, telemetry, tracking and audit records. PgBouncer is the connection-pooling layer.

Provider credentials, API tokens, signing keys and private keys are runtime secrets and must not be stored in Git.
