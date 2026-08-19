package backbone

import "context"

type Adapter interface {
    Transport() Transport
    Probe(context.Context, Link) (State, error)
}

// Registry keeps transport implementations behind a common Aether-Core API.
type AdapterRegistry struct { adapters map[Transport]Adapter }

func NewAdapterRegistry() *AdapterRegistry { return &AdapterRegistry{adapters: make(map[Transport]Adapter)} }

func (r *AdapterRegistry) Register(a Adapter) { r.adapters[a.Transport()] = a }

func (r *AdapterRegistry) Get(t Transport) (Adapter, bool) { a, ok := r.adapters[t]; return a, ok }
