CREATE TABLE nodes (
    -- Resource Core
    id                      UUID PRIMARY KEY,
    resource_version        TEXT NOT NULL,
    label                   TEXT NOT NULL,
    description             TEXT,
    tags                    JSONB,
    -- Node
    href                    TEXT,           -- Officially depricated v1.1+
    hostname                TEXT,           -- Officially depricated v1.1+
    api_versions            TEXT[],         
    api_endpoints           JSONB,          
    caps                    JSONB,          
    services                JSONB NOT NULL,
    clocks                  JSONB,
    interfaces              JSONB,
    supported_versions      TEXT[] NOT NULL,
    -- Metadata
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE devices (
    -- Resource Core
    id                      UUID PRIMARY KEY,
    resource_version        TEXT NOT NULL,
    label                   TEXT NOT NULL,
    description             TEXT,
    tags                    JSONB,
    -- Device
    type                    TEXT NOT NULL,
    receivers               TEXT[],
    senders                 TEXT[],
    node_id                 UUID REFERENCES nodes(id),
    controls                JSONB,
    -- Metadata
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE sources (
    -- Resource Core
    id                      UUID PRIMARY KEY,
    resource_version        TEXT NOT NULL,
    label                   TEXT NOT NULL,
    description             TEXT,
    tags                    JSONB,
    -- Source Core
    grain_rate              JSONB,
    caps                    JSONB,
    deviceid                UUID REFERENCES devices(id),
    parents                 TEXT[] NOT NULL,
    clock_name              TEXT,
    format                  TEXT NOT NULL,
    -- Audio Source
    audio_channels          JSONB,
    -- Metadata
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE flows (
    -- Resource Core
    id                      UUID PRIMARY KEY,
    resource_version        TEXT NOT NULL,
    label                   TEXT NOT NULL,
    description             TEXT,
    tags                    JSONB,
    -- Flow Core
    source_id               UUID REFERENCES sources(id),
    device_id               UUID REFERENCES devices(id),
    parents                 UUID[],
    grain_rate              JSONB,
    format                  TEXT NOT NULL,
    media_type              TEXT,
    -- Video
    frame_width             INT,
    frame_height            INT,
    interlace_mode          TEXT,
    colorspace              TEXT,
    transfer_characteristic TEXT,
    -- Video Raw
    components              JSONB,
    -- Audio
    sample_rate             JSONB,
    -- Audio Raw
    bit_depth               INT,
    -- JSON Data
    event_type              TEXT,
    -- SDI_ANC Data
    did_sdid                JSONB,
    -- Metadata
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE senders (
    -- Resource Core
    id                      UUID PRIMARY KEY,
    resource_version        TEXT NOT NULL,
    label                   TEXT NOT NULL,
    description             TEXT,
    tags                    JSONB,
    -- Sender
    caps                    JSONB,
    flow_id                 UUID REFERENCES flows(id),
    transport               TEXT NOT NULL,
    device_id               UUID REFERENCES devices(id),
    manifest_href           TEXT,
    interface_bindings      TEXT[],
    subscription_receiver   UUID,
    subscription_active     BOOLEAN,
    transport_file          BYTEA,
    -- Metadata
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE receivers (
    -- Resource Core
    id                      UUID PRIMARY KEY,
    resource_version        TEXT NOT NULL,
    label                   TEXT NOT NULL,
    description             TEXT,
    tags                    JSONB,
    -- Receiver
    device_id               UUID REFERENCES devices(id),
    transport               TEXT,
    interface_bindings      TEXT[],
    subscription_sender     UUID,
    subscription_active     BOOLEAN,
    format                  TEXT,
    caps                    JSONB,
    -- Metadata
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);