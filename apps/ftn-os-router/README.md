# FTN.OS Router Control Plane Agent

FTN.OS is the node-side contract for the FTN control plane. The agent is responsible for authenticated enrollment, health reporting, configuration inventory, route/tunnel telemetry and signed configuration lifecycle.

## Required controls

- mTLS using FTN Private PKI
- TPM 2.0/HSM-backed identity where available
- signed configuration manifests
- monotonic configuration versions
- health-gated activation
- deterministic rollback
- no central storage of device private keys

## Adapters

- GoBGP / FRRouting
- WireGuard / AmneziaWG
- eBPF/XDP policy
- DNS role adapters
- Prometheus telemetry

The agent must not accept arbitrary unauthenticated commands from the public Internet.
