CREATE TABLE IF NOT EXISTS platform_servers (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    cluster_id TEXT NOT NULL,
    address TEXT,
    public_address TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    region TEXT,
    state TEXT NOT NULL DEFAULT 'unknown',
    last_seen TIMESTAMPTZ,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS service_providers (
    id TEXT PRIMARY KEY,
    kind TEXT NOT NULL,
    name TEXT NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT true,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS dns_endpoints (
    id TEXT PRIMARY KEY,
    provider_id TEXT NOT NULL REFERENCES service_providers(id),
    hostname TEXT,
    address TEXT NOT NULL,
    role TEXT NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT true,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS dns_health_observations (
    id BIGSERIAL PRIMARY KEY,
    endpoint_id TEXT NOT NULL REFERENCES dns_endpoints(id),
    observed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    state TEXT NOT NULL,
    latency_ms DOUBLE PRECISION,
    packet_loss_percent DOUBLE PRECISION,
    rcode TEXT,
    recursion_available BOOLEAN,
    authoritative BOOLEAN,
    propagation JSONB NOT NULL DEFAULT '{}'::jsonb,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb
);

CREATE TABLE IF NOT EXISTS topology_links (
    id TEXT PRIMARY KEY,
    from_server_id TEXT NOT NULL REFERENCES platform_servers(id),
    to_server_id TEXT NOT NULL REFERENCES platform_servers(id),
    transport TEXT,
    state TEXT NOT NULL DEFAULT 'unknown',
    latency_ms DOUBLE PRECISION,
    packet_loss_percent DOUBLE PRECISION,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS control_intents (
    id TEXT PRIMARY KEY,
    target_type TEXT NOT NULL,
    target_id TEXT NOT NULL,
    operation TEXT NOT NULL,
    requested_by TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    requires_approval BOOLEAN NOT NULL DEFAULT true,
    approved_by TEXT,
    approved_at TIMESTAMPTZ,
    executed_at TIMESTAMPTZ,
    result JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_dns_health_endpoint_time
    ON dns_health_observations(endpoint_id, observed_at DESC);
CREATE INDEX IF NOT EXISTS idx_servers_cluster_state
    ON platform_servers(cluster_id, state);
CREATE INDEX IF NOT EXISTS idx_control_intents_status
    ON control_intents(status, created_at DESC);
