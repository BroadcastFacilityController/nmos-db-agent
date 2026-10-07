-- name: GetSource :one
SELECT * FROM sources
WHERE id = $1 LIMIT 1;

-- name: ListSources :many
SELECT * FROM sources
ORDER BY id;

-- name: ListSourcesPaginated :many
SELECT * FROM sources
WHERE id >= @starting_id
ORDER BY id
LIMIT @page_size;

-- name: ListSourcesByFormat :many
SELECT * FROM sources
WHERE format = $1
ORDER BY id;

-- name: ListSourcesByCreatedSince :many
SELECT * FROM sources
WHERE meta_created_at >= @created_since
ORDER BY meta_created_at;

-- name: ListSourcesByDeviceID :many
SELECT * FROM sources
WHERE device_id = $1
ORDER BY id;

-- name: ListSourcesByDeviceIDAndFormat :many
SELECT * FROM sources
WHERE device_id = $1 AND format = $2
ORDER BY id;

-- name: UpsertSource :one
INSERT INTO sources (
    id, resource_version, label, description, tags, grain_rate, caps, device_id, parents, clock_name, format, audio_channels, meta_user_label, meta_api_version
) VALUES (
    @id, @resource_version, @label, @description, @tags, @grain_rate, @caps, @device_id, @parents, @clock_name, @format, @audio_channels, @meta_user_label, @meta_api_version
)
ON CONFLICT (id) 
DO UPDATE
SET
    resource_version = EXCLUDED.resource_version,
    label = EXCLUDED.label,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    grain_rate = EXCLUDED.grain_rate,
    caps = EXCLUDED.caps,
    device_id = EXCLUDED.device_id,
    parents = EXCLUDED.parents,
    clock_name = EXCLUDED.clock_name,
    format = EXCLUDED.format,
    audio_channels = EXCLUDED.audio_channels,
    meta_user_label = EXCLUDED.meta_user_label
WHERE
    sources.meta_api_version = EXCLUDED.meta_api_version
    AND sources.resource_version < EXCLUDED.resource_version
RETURNING *;