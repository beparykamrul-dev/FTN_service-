package backbone

import (
    "sync"
    "time"
)

// Aether-Core is the FTN control-plane abstraction for secure transport,
// telemetry and policy orchestration. It is not a new kernel or cryptographic
// protocol; transport implementations remain replaceable adapters.
type Transport string

const (
    WireGuard Transport = "wireguard"
    WGQ Transport = "wgq-experimental"
    AmneziaWG Transport = "amneziawg"
)

type State string

const (
    Healthy State = "healthy"
    Degraded State = "degraded"
    Offline State = "offline"
    Unknown State = "unknown"
)

type Link struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Transport Transport `json:"transport"`
    State State `json:"state"`
    Endpoint string `json:"endpoint,omitempty"`
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
