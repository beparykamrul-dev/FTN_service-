# FTN Service

Production-oriented service foundation for Family Time Network (FTN).

## Completed foundation

- 001 — Frontier Registry & Unified Telemetry
- 002 — Physics Data Ingestion & Source Adapter Engine
- 003 — Central Telemetry Bus, Metrics/Event retention and API tracking
- Central database contract with PostgreSQL production target
- Central audit/tracking records
- Data-source health schema

## HTTP

- `GET /healthz`
- `GET /readyz`
- `GET /api/v1/frontier/modules`
- `GET /api/v1/telemetry/events`
- `GET /api/v1/tracking/events`

## Data plane

`migrations/001_core.sql` defines the central PostgreSQL schema for modules, telemetry, tracking and data sources. `docker-compose.yml` provides the database service for deployment.

The domain layer is kept behind interfaces so another database implementation can be attached without rewriting the service contracts.

## Governance

State-changing automation must remain approval-first and auditable. Hypothetical physics concepts are classified as research/simulation rather than presented as real communication transports.
