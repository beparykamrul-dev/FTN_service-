# Aether-Core Backbone

Aether-Core is the FTN overlay/control-plane orchestration layer for a provider-neutral backbone. Local fiber is not a required dependency.

## Path selection

The path policy evaluates health telemetry such as latency and packet loss before a path is eligible. Additional production adapters can supply bandwidth, jitter, IPv4/IPv6 reachability, NAT state, routing state, and tunnel health.

## Failover model

1. Collect path telemetry.
2. Apply policy and security requirements.
3. Mark paths healthy/degraded/unavailable.
4. Select an eligible path.
5. Record the decision centrally.
6. Re-evaluate continuously.

State-changing network operations remain approval-controlled. Aether-Core does not claim to replace BGP/OSPF/IS-IS or to provide physical connectivity by itself.

## Target transports

Production profiles: WireGuard, AmneziaWG, QUIC, TLS, IPsec.

Research profiles: WGQ and other explicitly experimental adapters.

Overlay integrations: GRE, VXLAN, DNS Mesh, Anycast.
