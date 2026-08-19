package frontier

import (
    "encoding/json"
    "net/http"
    "sync"
)

type Classification string

const (
    Real          Classification = "real"
    Experimental  Classification = "experimental"
    Simulation    Classification = "simulation"
    Theoretical   Classification = "theoretical"
    Hypothetical  Classification = "hypothetical"
    Unsupported   Classification = "unsupported"
)

type Status string

const (
    Active   Status = "active"
    Disabled Status = "disabled"
)

type Module struct {
    ID             string         `json:"id"`
    Name           string         `json:"name"`
    Domain         string         `json:"domain"`
    Classification Classification `json:"classification"`
    Status         Status         `json:"status"`
}

type Registry struct {
    mu      sync.RWMutex
    modules map[string]Module
}

func NewRegistry() *Registry { return &Registry{modules: make(map[string]Module)} }

func (r *Registry) Register(m Module) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.modules[m.ID] = m
}

func (r *Registry) Snapshot() []Module {
    r.mu.RLock()
    defer r.mu.RUnlock()
    out := make([]Module, 0, len(r.modules))
    for _, m := range r.modules { out = append(out, m) }
    return out
}

func (r *Registry) Handler(w http.ResponseWriter, _ *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(r.Snapshot())
}
