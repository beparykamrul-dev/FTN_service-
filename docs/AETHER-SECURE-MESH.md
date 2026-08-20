# Aether-Core Secure Mesh

## Architecture

Aether-Core coordinates a provider-neutral FTN overlay. It does not replace the underlying routing or tunnel implementations.

### Secure mesh

- WireGuard mesh: production encrypted transport
- AmneziaWG: production transport profile where operationally appropriate
- GoBGP: dynamic routing adapter and telemetry boundary
- eBPF/XDP firewall: host/edge packet-policy enforcement boundary
- mTLS + private PKI: service identity and authenticated control-plane communication
- PQC: cryptographic-agility layer for post-quantum algorithms and hybrid handshakes where supported

## Routing safety

BGP changes are policy-validated and approval-controlled. Aether-Core may recommend a path or route change from telemetry, but it must not bypass the approval policy for state-changing operations.

## Firewall safety

eBPF/XDP programs must be versioned, validated, resource-bounded and staged before production activation. A failed program or health regression must have a deterministic rollback path.

## PQC

PQC is an algorithm-agility capability, not a claim that every tunnel is already quantum-resistant. The implementation should support standards-based algorithms and hybrid classical+PQC modes only when the selected protocol/library provides a reviewed implementation.

Private keys remain endpoint-controlled and are never stored in the central database. Certificate status, algorithm/profile metadata, peer identity, routing decisions, firewall policy versions and health events are centrally tracked.
