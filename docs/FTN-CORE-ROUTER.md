# FTN Core Router & Health-Driven Firmware Architecture

## Objective

FTN Core Router is the unified routing/control profile for FTN infrastructure. Firmware is selected and built from detected hardware capabilities, supported kernel/network features, and an explicit device profile.

The system must never flash firmware automatically. Build, validation, signing, staging, and deployment are separate approval-controlled steps.

## Health-driven build pipeline

```text
Router Discovery
  -> Hardware/OS Capability Probe
  -> Firmware Profile Selection
  -> Config Compatibility Check
  -> Build
  -> Reproducibility Check
  -> Static/Security Validation
  -> Artifact Hash + Manifest
  -> Staging
  -> Approval
  -> Controlled Deployment
  -> Post-boot Health Verification
  -> Rollback on failed health gate
```

## Health signals

- CPU / memory / storage
- interface state and errors
- throughput
- packet loss / latency
- routing/BGP state
- DNS reachability
- WireGuard/AmneziaWG tunnel state
- kernel/service health
- configuration drift
- firmware version and artifact hash

## Supported transport profiles

- Standard WireGuard
- AmneziaWG where the target platform supports it
- WGQ remains experimental/research-only until a defined, reviewed implementation exists

## Firmware safety

No unverified firmware is promoted to production. Each artifact receives a version, target hardware profile, source revision, build metadata, cryptographic hash, validation result, and deployment approval record.
