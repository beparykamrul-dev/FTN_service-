# FTN Unified Control Plane

FTN is designed to operate many physical servers as one logical service cluster.

## Target scale

The control plane is designed for 20–30 servers initially and remains extensible beyond that range.

## Model

```text
Servers -> Agent/Telemetry -> Central Event Bus -> Central Database
                         -> Health/Capacity Engine
                         -> Geographic Topology
                         -> Policy/Approval
                         -> Control Adapters
                         -> Audit/Tracking
```

## Unified dashboard

- Global service health
- Server and cluster health
- CPU/memory/storage/network capacity
- Latency and packet loss
- WireGuard/WGQ/AmneziaWG transport status
- DNS/BGP/service health
- Geographic map with server nodes and inter-server links
- Alerts and incidents
- Pending control actions
- Audit trail

## Control model

Actions are represented as intents and routed through approved adapters. A monitoring failure must not automatically become an arbitrary infrastructure change. Destructive or state-changing operations require policy/approval checks.

## Map model

Every server may have latitude/longitude and region metadata. Links can expose latency, packet loss and state. This supports a geographic NOC view without coupling the domain model to a particular map vendor.
