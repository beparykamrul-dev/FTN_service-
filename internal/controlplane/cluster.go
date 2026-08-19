package controlplane

import "sync"

type Cluster struct {
    ID string `json:"id"`
    Name string `json:"name"`
    DesiredCapacity int `json:"desired_capacity"`
    ServerIDs []string `json:"server_ids"`
    LoadPolicy string `json:"load_policy"`
    AutoBalance bool `json:"auto_balance"`
}

type ClusterRegistry struct { mu sync.RWMutex; clusters map[string]Cluster }
func NewClusterRegistry() *ClusterRegistry { return &ClusterRegistry{clusters:make(map[string]Cluster)} }
func (r *ClusterRegistry) Upsert(c Cluster) { r.mu.Lock(); defer r.mu.Unlock(); r.clusters[c.ID]=c }
func (r *ClusterRegistry) Snapshot() []Cluster { r.mu.RLock(); defer r.mu.RUnlock(); out:=make([]Cluster,0,len(r.clusters)); for _,c:=range r.clusters{out=append(out,c)}; return out }
