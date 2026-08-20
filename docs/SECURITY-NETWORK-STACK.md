# FTN Security & Network Analysis Stack

## Private PKI / mTLS

FTN services should use a private CA hierarchy for service identity and mutual TLS. Root keys remain offline; intermediates issue short-lived service certificates. Certificate issuance, renewal, revocation and ownership are centrally tracked without storing private keys in PostgreSQL.

## Revocation

Support both OCSP and CRL workflows. OCSP stapling is preferred for TLS services where supported; CRL distribution remains the recovery/fallback mechanism. Certificate status must be checked before privileged service-to-service operations.

## Routing policy

`go-bgp` is an adapter boundary for BGP state/control. Route policy must be explicit, validated and approval-controlled before changes are applied.

## Traffic analytics

`rwfilter` and PySiLK are treated as flow-analysis adapters. They consume authorized NetFlow/IPFIX/SiLK datasets and produce normalized telemetry; they do not directly modify network state.

## Central tracking

PKI events, route decisions, certificate status, flow-analysis jobs and security findings map to the existing centralized tracking/audit model.
