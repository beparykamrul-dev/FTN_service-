package backbone

import (
	"sort"
	"time"
)

type EdgeNode struct {
	ID           string    `json:"id"`
	Region       string    `json:"region"`
	Endpoint     string    `json:"endpoint"`
	Healthy      bool      `json:"healthy"`
	LatencyMS    float64   `json:"latency_ms"`
	PacketLoss   float64   `json:"packet_loss_pct"`
	LastChecked  time.Time `json:"last_checked"`
}

// RankEdges returns healthy edge candidates by a deterministic health score.
// It does not change routing state; a separate approved adapter performs changes.
func RankEdges(nodes []EdgeNode) []EdgeNode {
	out := make([]EdgeNode, 0, len(nodes))
	for _, n := range nodes {
		if n.Healthy {
			out = append(out, n)
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		si := out[i].LatencyMS + (out[i].PacketLoss * 100)
		sj := out[j].LatencyMS + (out[j].PacketLoss * 100)
		return si < sj
	})
	return out
}
