package controlplane

type MapNode struct {
    ServerID string `json:"server_id"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Label string `json:"label,omitempty"`
    State ServerState `json:"state"`
}

type MapLink struct {
    ID string `json:"id"`
    From string `json:"from"`
    To string `json:"to"`
    LatencyMs float64 `json:"latency_ms,omitempty"`
    PacketLossPercent float64 `json:"packet_loss_percent,omitempty"`
    State ServerState `json:"state"`
}

type TopologyMap struct { Nodes []MapNode `json:"nodes"`; Links []MapLink `json:"links"` }
