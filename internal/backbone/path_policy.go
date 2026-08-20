package backbone

// PathPolicy defines deterministic inputs for selecting a healthy backbone path.
// The policy is transport-agnostic; it does not itself configure routers.
type PathPolicy struct {
	MaxLatencyMS       float64 `json:"max_latency_ms"`
	MaxPacketLossPct   float64 `json:"max_packet_loss_pct"`
	MinAvailableMbps   float64 `json:"min_available_mbps"`
	PreferIPv6         bool    `json:"prefer_ipv6"`
	AllowDegraded      bool    `json:"allow_degraded"`
	RequireEncryption  bool    `json:"require_encryption"`
}

func DefaultPathPolicy() PathPolicy {
	return PathPolicy{
		MaxLatencyMS:     150,
		MaxPacketLossPct: 2,
		MinAvailableMbps: 10,
		PreferIPv6:       true,
		AllowDegraded:    false,
		RequireEncryption: true,
	}
}

func Eligible(link Link, p PathPolicy) bool {
	if link.State != Healthy && !(p.AllowDegraded && link.State == Degraded) {
		return false
	}
	if p.MaxLatencyMS > 0 && link.LatencyMS > p.MaxLatencyMS {
		return false
	}
	if p.MaxPacketLossPct > 0 && link.PacketLoss > p.MaxPacketLossPct {
		return false
	}
	return true
}
