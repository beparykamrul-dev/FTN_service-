# FTN Global DNS Mesh & FTN.OS Router Control Plane

## Production topology

User → authoritative/global DNS → nearest healthy Edge PoP → edge proxy/filtering → Aether-Core encrypted overlay → FTN.OS/Core Router → application services.

The design removes dependency on a local fiber backbone. It still requires real upstream Internet/WAN/cloud connectivity; Aether-Core is an overlay/control plane, not physical connectivity.

## DNS

`ns1.familytimenet.com` and `ns2.familytimenet.com` remain the existing authoritative anchors. New DNS providers/servers can be enrolled through the universal DNS adapter/template system. Zone imports must validate SOA/NS/DNSSEC and detect conflicts before activation.

## Edge

Each PoP exposes only the services required for that PoP. Envoy/Traefik/NGINX/HAProxy are interchangeable adapter boundaries. eBPF/XDP and FastNetMon may perform defensive filtering where supported.

## Backbone

Aether-Core selects healthy encrypted paths using telemetry. GoBGP/FRR provide routing adapters. SR/SRv6 can be used where the upstream network supports it. Route changes are validated and approval-controlled.

## FTN.OS Router Control Plane

FTN.OS is treated as the router-agent/control-plane contract:

- device enrollment and identity
- TPM/HSM attestation metadata
- configuration inventory and hash
- interface/route/tunnel health
- BGP/FRR state
- firewall policy version
- DNS role
- firmware/version state
- signed configuration and rollback status

The control plane observes and proposes changes; destructive or externally visible changes require explicit policy/approval.

## Global telemetry

Flow telemetry is normalized before storage. ClickHouse is the high-volume analytical store; PostgreSQL remains the authoritative relational/audit store. ELK/SiLK/YAF/Nfdump/pmacct and other collectors can feed the normalized model.

## Traffic handling

Use standards-compliant encrypted transports, authenticated service identity and defensive filtering. No component is designed to make traffic "unblockable" or to evade lawful network/security controls.
