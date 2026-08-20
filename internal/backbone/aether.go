package backbone

import (
    "sync"
    "time"
)

// Aether-Core is the FTN backbone orchestration layer. It is transport-agnostic:
// production transports, research profiles, and future adapters can coexist
// without making Aether-Core itself a new cryptographic protocol or kernel.
type Transport string

const (
    WireGuard Transport = "wireguard"
    AmneziaWG Transport = "amneziawg"
    QUIC Transport = "quic"
    TLS Transport = "tls"
    IPsec Transport = "ipsec"
    WGQ Transport = "wgq-experimental"
    PostQuantum Transport = "post-quantum-experimental"
    QuantumResearch Transport = "quantum-research"
    PhotonicResearch Transport = "photonic-research"
    GRE Transport = "gre"
    VXLAN Transport = "vxlan"
    DNSMesh Transport = "dns-mesh"
    Anycast Transport = "anycast"
)

type State string

const (
    Healthy State = "healthy"
    Degraded State = "degraded"
    Offline State = "offline"
    Unknown State = "unknown"
)

type Capability string

const (
    Routing Capability = "routing"
    Encryption Capability = "encryption"
    Overlay Capability = "overlay"
    DNS Capability = "dns"
    Telemetry Capability = "telemetry"
    Experimental Capability = "experimental"
)

type Link struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Transport Transport `json:"transport"`
    State State `json:"state"`
    Capabilities []Capability `json:"capabilities,omitempty"`
    Endpoint string `json:"endpoint,omitempty"`
    Region string `json:"region,omitempty"`
    LatencyMS float64 `json:"latency_ms,omitempty"`
    PacketLoss float64 `json:"packet_loss,omitempty"`
    RXBytes uint64 `json:"rx_bytes,omitempty"`
    TXBytes uint64 `json:"tx_bytes,omitempty"`
    LastChecked time.Time `json:"last_checked"`
}

type Backbone struct {
    mu sync.RWMutex
    links map[string]Link
}

func New() *Backbone { return &Backbone{links: make(map[string]Link)} }

func (b *Backbone) Upsert(link Link) {
    b.mu.Lock()
    defer b.mu.Unlock()
    if link.LastChecked.IsZero() { link.LastChecked = time.Now().UTC() }
    b.links[link.ID] = link
}

func (b *Backbone) Snapshot() []Link {
    b.mu.RLock()
    defer b.mu.RUnlock()
    out := make([]Link, 0, len(b.links))
    for _, link := range b.links { out = append(out, link) }
    return out
}
