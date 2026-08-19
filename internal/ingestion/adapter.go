package ingestion

import "context"

type Record struct {
    SourceID   string            `json:"source_id"`
    ObservedAt string            `json:"observed_at"`
    Metrics    map[string]float64 `json:"metrics"`
    Labels     map[string]string `json:"labels"`
}

type Adapter interface {
    ID() string
    Collect(context.Context) ([]Record, error)
}

type Manager struct { adapters map[string]Adapter }

func NewManager() *Manager { return &Manager{adapters: make(map[string]Adapter)} }

func (m *Manager) Register(a Adapter) { m.adapters[a.ID()] = a }

func (m *Manager) Collect(ctx context.Context) ([]Record, error) {
    var all []Record
    for _, a := range m.adapters {
        records, err := a.Collect(ctx)
        if err != nil { return nil, err }
        all = append(all, records...)
    }
    return all, nil
}
