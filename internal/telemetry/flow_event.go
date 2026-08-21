package telemetry

import "time"

type FlowEvent struct {
	Timestamp   time.Time `json:"timestamp"`
	Source      string    `json:"source"`
	Exporter    string    `json:"exporter"`
	SrcAddr     string    `json:"src_addr"`
	DstAddr     string    `json:"dst_addr"`
	SrcPort     uint16    `json:"src_port"`
	DstPort     uint16    `json:"dst_port"`
	Protocol    uint8     `json:"protocol"`
	Bytes       uint64    `json:"bytes"`
	Packets     uint64    `json:"packets"`
	Region      string    `json:"region,omitempty"`
	EdgeNode    string    `json:"edge_node,omitempty"`
}

// Normalized flow events are suitable for ClickHouse/ELK/SiLK adapters.
// This type deliberately contains no packet payload or credential material.
