# FTN DNS Global Registry

## Primary principle

DNS is the primary global entry layer. Every FTN-operated DNS server is an independently addressable node that can be registered, health-checked, synchronized and promoted without replacing the existing FTN DNS service.

### FamilyTimeNet authoritative identity

```text
ns1.familytimenet.com
  100.52.87.3
  138.252.113.19

ns2.familytimenet.com
  100.52.86.20
  138.252.113.20
```

Additional FTN DNS nodes are added through the same registry and synchronization contract. The registry supports IPv4/IPv6, authoritative/recursive roles, DNSSEC capability, health endpoints, upstreams, Anycast eligibility and provider metadata.

## Universal DNS onboarding

A new DNS provider/server follows:

```text
Discover / Import
      -> Validate IP + hostname
      -> Detect DNS engine
      -> Generate provider-specific template
      -> Import existing zones/records
      -> Conflict analysis
      -> FamilyTimeNet alignment
      -> DNSSEC/TTL/SOA validation
      -> Health check
      -> Staged activation
      -> Synchronize
      -> Monitor
```

Existing records are preserved unless an explicit change policy authorizes replacement. Every change has a version, audit record and rollback representation.

## Engines

Adapters are supported for PowerDNS, Technitium, CoreDNS, Unbound, dnsdist, GoDNS, Hickory and other compatible authoritative/recursive services. Provider integrations may include Cloudflare, DNSPod/Tencent and Akamai where credentials and APIs are supplied by the operator.

## GeoIP2

GeoIP2 is used only for legitimate geolocation-aware policy and telemetry. It can provide country/region/ASN metadata for DNS health analysis, Edge PoP selection and observability; it is not used to infer private identity.

## GoBGP

GoBGP is the application-facing BGP control component. It is distinct from FRRouting and is integrated through an adapter so FTN can choose the appropriate routing implementation per node. Route advertisements require explicit policy validation and operator authorization.

## Traffic path

```text
Client
  -> FamilyTimeNet DNS / Anycast DNS
  -> selected healthy FTN DNS/Edge node
  -> Edge proxy
  -> Aether-Core
  -> FTN.OS / GoBGP
  -> WireGuard/AmneziaWG where required
  -> FTN service
```

The DNS registry is provider-neutral: adding another provider or FTN DNS server does not require redesigning the existing FamilyTimeNet DNS service.

## Safety and operational controls

- DNS changes are staged and reversible.
- Credentials remain outside Git.
- BGP announcements require authorization.
- DNSSEC and certificate state are monitored.
- Node health is measured independently of routing preference.
- A failed node is removed from eligible service only after policy-defined health thresholds.
