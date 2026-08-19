package controlplane

import (
    "sync"
    "time"
)

type ServerState string
const (
    Online ServerState = "online"
    Degraded ServerState = "degraded"
    Offline ServerState = "offline"
    Unknown ServerState = "unknown"
)

type Server struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Address string `json:"address"`
    Latitude float64 `json:"latitude,omitempty"`
    Longitude float64 `json:"longitude,omitempty"`
    Region string `json:"region,omitempty"`
    State ServerState `json:"state"`
    LastSeen time.Time `json:"last_seen"`
    CPUPercent float64 `json:"cpu_percent,omitempty"`
    MemoryPercent float64 `json:"memory_percent,omitempty"`
    BandwidthMbps float64 `json:"bandwidth_mbps,omitempty"`
}

type ServerRegistry struct { mu sync.RWMutex; servers map[string]Server }
func NewServerRegistry() *ServerRegistry { return &ServerRegistry{servers: make(map[string]Server)} }
func (r *ServerRegistry) Upsert(s Server) { r.mu.Lock(); defer r.mu.Unlock(); if s.LastSeen.IsZero(){s.LastSeen=time.Now().UTC()}; r.servers[s.ID]=s }
func (r *ServerRegistry) Snapshot() []Server { r.mu.RLock(); defer r.mu.RUnlock(); out:=make([]Server,0,len(r.servers)); for _,s:=range r.servers{out=append(out,s)}; return out }
