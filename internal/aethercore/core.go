package aethercore

import "time"

// Aether-Core is FTN's transport-agnostic backbone control abstraction.
// It orchestrates approved transport adapters; it is not a new cryptographic protocol.
type State string

const (
	Unknown State = "unknown"
	Healthy State = "healthy"
	Degraded State = "degraded"
	Offline State = "offline"
)

type TransportProfile struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mode      string `json:"mode"`
	Enabled   bool   `json:"enabled"`
	Research  bool   `json:"research"`
}

type LinkHealth struct {
	ProfileID   string        `json:"profile_id"`
	State       State         `json:"state"`
	LatencyMs   float64       `json:"latency_ms"`
	PacketLoss  float64       `json:"packet_loss_percent"`
	RxBytes     uint64        `json:"rx_bytes"`
	TxBytes     uint64        `json:"tx_bytes"`
	LastSeen    time.Time     `json:"last_seen"`
}

type Core struct {
	Profiles []TransportProfile `json:"profiles"`
}

func DefaultCore() Core {
	return Core{Profiles: []TransportProfile{
		{ID: "wireguard", Name: "WireGuard", Mode: "production", Enabled: true},
		{ID: "amneziawg", Name: "AmneziaWG", Mode: "production-profile", Enabled: true},
		{ID: "wgq", Name: "WGQ", Mode: "experimental", Enabled: false, Research: true},
	}}
}
