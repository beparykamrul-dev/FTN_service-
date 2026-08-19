# FTN Service

Production-oriented service foundation for Family Time Network (FTN).

## Frontier Control Plane

A unified registry and telemetry foundation for real, experimental, simulation, theoretical, and hypothetical frontier-research modules.

### Current modules

- 001 — Frontier Registry & Unified Telemetry
- 002 — Physics Data Ingestion & Source Adapter Engine

### Design rules

- Explicit classification of scientific claims
- Source adapters isolated behind interfaces
- Deterministic telemetry records
- Approval-first state-changing controls
- Auditability for control operations
- Hypothetical physics is never presented as production transport

## HTTP

- `GET /healthz`
- `GET /readyz`
- `GET /api/v1/frontier/modules`
