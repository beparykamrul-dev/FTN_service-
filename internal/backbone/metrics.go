package backbone

import "time"

type LinkMetrics struct {
    LinkID string `json:"link_id"`
    Transport Transport `json:"transport"`
    State State `json:"state"`
    LatencyMs float64 `json:"latency_ms"`
    PacketLossPct float64 `json:"packet_loss_pct"`
    RxBytes uint64 `json:"rx_bytes"`
    TxBytes uint64 `json:"tx_bytes"`
    HandshakeAgeSec float64 `json:"handshake_age_sec"`
    ObservedAt time.Time `json:"observed_at"`
}
