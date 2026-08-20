# FTN Edge Trust & Automation

## Hardened Edge-to-Core Headend

The edge/headend is a trust boundary between untrusted networks and FTN core services. Use deny-by-default ingress, mTLS service identity, certificate revocation checks, rate limits, management-plane isolation, signed configuration, health-gated promotion, and deterministic rollback.

## TPM 2.0 / HSM

TPM 2.0 is the preferred endpoint hardware-root-of-trust for measured boot, device identity and protected key operations. HSMs are the preferred boundary for high-value CA/signing keys. Private keys must not be copied into the central PostgreSQL database.

## Consul

HashiCorp Consul is an optional service-discovery/configuration adapter. It must not become the sole source of truth: FTN's centralized database remains authoritative for inventory and audit state. Consul health checks can feed Aether-Core telemetry.

## GeoIP

P2Location and DB-IP Free are optional GeoIP data adapters. Store provider/version/license metadata and normalized region/country information; do not treat GeoIP as precise physical location.

## Lua scripting

Lua is an embedded, sandboxed policy/automation extension point. Scripts are versioned, resource-limited, denied network/filesystem access by default, and cannot bypass approval or security policy.

## AI traffic classification

AI may classify and score traffic anomalies or protocol/application fingerprints for defensive routing and QoS decisions. It must not be used to conceal malicious activity or bypass provider/security controls. Model decisions are advisory unless an approved deterministic policy permits automation.

## Central tracking

Record TPM/HSM attestation state, certificate status, Consul health, GeoIP dataset version, Lua policy version, AI classification, edge health, configuration hash and approval/audit events centrally.
