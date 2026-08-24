# FTN DNS Automatic Adjustment Runbook

## Objective

Onboard any new FTN DNS server/provider while preserving `familytimenet.com` as the authoritative FTN identity and avoiding unplanned record replacement.

## Pipeline

```text
New Server / Provider
        ↓
DNS engine detection
        ↓
IP + hostname validation
        ↓
Zone/record import
        ↓
Normalize A/AAAA/CNAME/MX/NS/SRV/TXT/CAA/PTR/HTTPS/SVCB
        ↓
SOA + TTL + DNSSEC analysis
        ↓
Diff against FTN DNS registry
        ↓
Conflict report
        ↓
Operator-approved change set
        ↓
Staged sync
        ↓
Health + resolution tests
        ↓
Eligible for FTN mesh
```

## Supported adjustment targets

- PowerDNS
- Technitium
- CoreDNS
- Unbound
- dnsdist
- GoDNS
- Hickory DNS
- BIND-compatible zone sources
- Cloudflare API
- DNSPod/Tencent DNS API
- Akamai DNS API

Provider credentials are supplied at deployment time and are never committed to Git.

## familytimenet.com alignment

The registry treats the existing FTN authoritative identity as the source of policy, not as a reason to blindly overwrite another server. New nodes receive a generated change set containing only the required differences. Existing records outside the managed policy remain untouched unless explicitly adopted.

## Safety

- Dry-run diff before activation
- No destructive default operation
- SOA serial/version tracking
- DNSSEC state verification
- staged health checks
- rollback snapshot
- audit event for every accepted change
- automatic exclusion of unhealthy nodes

## Anycast and edge readiness

A DNS node becomes Anycast/edge eligible only after resolver/authoritative health, DNSSEC status, network reachability and telemetry checks pass. DNS health does not by itself authorize BGP route advertisement; routing remains a separate GoBGP/FRR policy gate.
