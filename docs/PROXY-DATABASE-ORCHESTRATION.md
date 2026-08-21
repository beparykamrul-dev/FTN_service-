# FTN Proxy & Database Orchestration

## Proxy layers

Aether-Core may coordinate these distinct boundaries:

- Sidecar proxy: per-service traffic and mTLS boundary
- Mesh/mTLS proxy: authenticated service-to-service communication
- Edge proxy: internet-facing ingress/egress boundary
- Database proxy: connection pooling, routing and health checks
- Envoy: service-mesh/edge proxy adapter
- Traefik: ingress/reverse-proxy adapter
- NGINX: edge/reverse/forward/caching proxy adapter

Proxy selection is policy-driven; multiple proxies are not required on every node.

## Database topology

PostgreSQL remains the authoritative relational store. A master/replica topology may be used for read scaling and resilience. PgBouncer sits at the connection-pooling boundary. Writes must have a single authoritative write path; replica lag is monitored before read routing.

## Service scheduling

HashiCorp Nomad is an optional workload scheduler. Consul supplies service discovery and health information. Neither replaces FTN's central inventory/audit database.

## Cryptography

Go's standard crypto libraries should be preferred for application cryptography. AES-256-GCM is available for authenticated encryption of appropriate application payloads; keys are managed through the PKI/TPM/HSM boundary rather than hard-coded or stored in PostgreSQL.

## ACME / PKI

ACME is used for publicly trusted certificate automation where appropriate. Private services use FTN Private PKI/mTLS. Certificate renewal, status and revocation events are centrally audited.

## Frontend visualization

React + WebGL is the target for the high-density FTN topology/NOC visualization. It must degrade gracefully to ordinary DOM/canvas layouts on devices without suitable WebGL capability and automatically adapt to screen size/pixel density.

## Gugi

"Gugi" is retained as a named integration placeholder until its exact product/library identity is confirmed; no unverified dependency is added under that name.

## Silk Road

"Silk Road" is treated only as the user's label for the Traefik proxy integration; it is not an external dependency or security feature.
