# Aether-Core

Aether-Core is the FTN transport-agnostic backbone orchestration layer.

## Transport profiles

- WireGuard — production transport.
- AmneziaWG — production transport profile where supported and approved.
- WGQ — experimental/research profile only; it is not represented as a standardized quantum-safe WireGuard protocol.

## Responsibilities

- Transport profile registry
- Link health aggregation
- Latency and packet-loss telemetry
- RX/TX counters
- Last-seen state
- Policy/approval integration
- Central tracking and audit integration

Aether-Core does not claim to provide FTL, wormholes, quantum teleportation, or other hypothetical transport mechanisms.
