# FTN Traffic Mimicry — Two-Mode Design

FTN supports two distinct, policy-controlled modes.

## Mode 1 — Defensive classification

AI analyzes authorized telemetry to classify application/protocol characteristics, detect anomalies, and improve QoS, routing, capacity planning, and incident response. This mode does not alter packet appearance.

## Mode 2 — Controlled traffic shaping

A sandboxed traffic-shaping adapter may normalize benign application traffic characteristics for FTN-owned services and authorized testing. It is limited to protocol-compliant, non-evasive shaping and must not be used to bypass access controls, provider restrictions, security monitoring, or detection systems.

### Common controls

- explicit policy and scope
- rate/CPU/memory limits
- deterministic configuration versioning
- telemetry before/after changes
- approval for state-changing production policies
- immediate rollback on health regression
- centralized audit/tracking

AI recommendations remain advisory unless an explicit deterministic policy permits automation.
