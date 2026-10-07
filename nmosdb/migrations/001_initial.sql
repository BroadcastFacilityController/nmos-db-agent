-- +goose Up
CREATE TABLE nodes (
    -- Resource Core
    id                      UUID PRIMARY KEY,
    resource_version        TEXT NOT NULL,
    label                   TEXT NOT NULL,
    description             TEXT,
    tags                    JSONB,
    -- Node
    api_versions            TEXT[],
    api_endpoints           JSONB,
    caps                    JSONB,
    services                JSONB NOT NULL,
    clocks                  JSONB,
    interfaces              JSONB,
    -- Metadata
    meta_user_label         TEXT,
    meta_api_version        TEXT NOT NULL,
    meta_created_at         TIMESTAMPTZ DEFAULT NOW()
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
    meta_user_label         TEXT,
    meta_api_version        TEXT NOT NULL,
    meta_created_at         TIMESTAMPTZ DEFAULT NOW()
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
    device_id                UUID REFERENCES devices(id),
    parents                 TEXT[] NOT NULL,
    clock_name              TEXT,
    format                  TEXT NOT NULL,
    -- Audio Source
    audio_channels          JSONB,
    -- Metadata
    meta_user_label         TEXT,
    meta_api_version        TEXT NOT NULL,
    meta_created_at         TIMESTAMPTZ DEFAULT NOW()
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
    meta_user_label         TEXT,
    meta_api_version        TEXT NOT NULL,
    meta_created_at         TIMESTAMPTZ DEFAULT NOW()
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
    meta_user_label         TEXT,
    meta_api_version        TEXT NOT NULL,
    meta_created_at         TIMESTAMPTZ DEFAULT NOW()
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
    meta_user_label         TEXT,
    meta_api_version        TEXT NOT NULL,
    meta_created_at         TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE receivers;
DROP TABLE senders;
DROP TABLE flows;
DROP TABLE sources;
DROP TABLE devices;
DROP TABLE nodes;