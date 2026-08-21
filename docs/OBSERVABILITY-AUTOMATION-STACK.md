# FTN Observability, Automation & Edge Stack

This document records integration boundaries for the FTN unified control plane. Components are adapters, not assumed installed services.

## Control plane

- ASP.NET Core — portal/control-plane integration boundary
- Go backend — core service APIs
- Aether-Core — backbone orchestration
- WebSockets/AJAX — live dashboard transport
- net.Conn — low-level Go networking boundary
- Mutex + worker pools — bounded concurrent processing

## Flow telemetry & analysis

- tools.netsa.cert.org / SiLK ecosystem
- YAF + SiLK
- PySiLK
- rwfilter
- Nfdump + NfSen
- ElastFlow
- pmacct
- Plixer Scrutinizer
- Kentik
- FastNetMon
- ELK Stack (Elasticsearch, Logstash, Kibana)
- ClickHouse

All flow sources normalize into the central telemetry model. Collection and analysis are read-only unless an explicit control adapter is approved.

## Routing & network automation

- GoBGP
- FRRouting
- Segment Routing / SRv6
- SaltStack
- Ansible
- Lua policy extensions
- NGINX edge/reverse/forward proxy boundary
- Transparent/caching proxy boundary
- SOCKS5 and approved tunnel adapters

## Geo / enrichment

- GeoIP / GeoIP2
- P2Location
- DB-IP Free

Dataset versions and licensing metadata are tracked centrally.

## Security / trust

- SASE / Zero Trust architecture
- eBPF/XDP enforcement boundary
- Private PKI
- Let's Encrypt for publicly trusted certificates where appropriate
- HashiCorp integrations (including Consul)
- Go crypto primitives
- checksum/integrity verification

## Platform / research references

- FreeBSD
- openSUSE Software
- IEEE Computer Society
- Jupyter

These are platform/research integration references, not blanket runtime dependencies.

## Go supply-chain integrity

`proxy.golang.org` and the Go Checksum Database are treated as dependency-distribution/integrity services. Dependency versions and checksums should be pinned and verified in CI.

## AI-Native Autonomous Networks

AI can observe telemetry, correlate events, score anomalies and propose actions. State-changing operations remain deterministic, policy-validated and approval-controlled unless an explicitly approved automation policy permits them.
