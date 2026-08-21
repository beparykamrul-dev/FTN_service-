# FTN Hickory DNS Integration

Hickory DNS (formerly Trust-DNS) is added as a Rust-based DNS integration option in the FTN DNS Mesh.

## Role

Hickory is an additional resolver/authoritative DNS adapter alongside PowerDNS, Technitium, CoreDNS, Unbound, dnsdist and GoDNS. It is not a replacement for the existing DNS stack unless explicitly selected by policy.

## Aether-Core integration

Aether-Core can consume Hickory health telemetry through the DNS adapter boundary:

- query latency
- response/error rate
- DNSSEC validation state where enabled
- endpoint reachability
- process/service health

DNS routing decisions remain policy-controlled and centrally tracked.

## Security

Use DNSSEC where applicable, TLS/mTLS for administrative/control-plane communication, and centralized certificate lifecycle management. Do not store private keys in PostgreSQL.

## Deployment

The repository records Hickory as an integration target; installation and production activation require an explicit environment-specific deployment configuration.
