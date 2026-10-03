-- name: GetFlow :one
SELECT * FROM flows
WHERE id = $1 LIMIT 1;

-- name: ListFlows :many
SELECT * FROM flows
ORDER BY id;

-- name: ListFlowsPaginated :many
SELECT * FROM flows
WHERE id >= @starting_id
ORDER BY id
LIMIT @page_size;

-- name: ListFlowsByFormat :many
SELECT * FROM flows
WHERE format = $1
ORDER BY id;

-- name: ListFlowsByCreatedSince :many
SELECT * FROM flows
WHERE meta_created_at >= @created_since
ORDER BY meta_created_at;

-- name: ListFlowsByDeviceID :many
SELECT * FROM flows
WHERE device_id = $1
ORDER BY id;

-- name: ListFlowsByDeviceIDAndFormat :many
SELECT * FROM flows
WHERE device_id = $1 AND format = $2
ORDER BY id;

-- name: UpsertFlow :one
INSERT INTO flows (
    id, resource_version, label, description, tags, source_id, device_id, parents, grain_rate, format, media_type, frame_width, frame_height, interlace_mode, colorspace, transfer_characteristic, components, sample_rate, bit_depth, event_type, did_sdid, meta_api_version
) VALUES (
    @id, @resource_version, @label, @description, @tags, @source_id, @device_id, @parents, @grain_rate, @format, @media_type, @frame_width, @frame_height, @interlace_mode, @colorspace, @transfer_characteristic, @components, @sample_rate, @bit_depth, @event_type, @did_sdid, @meta_api_version
)
ON CONFLICT (id) 
DO UPDATE
SET
    resource_version = EXCLUDED.resource_version,
    label = EXCLUDED.label,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    source_id = EXCLUDED.source_id,
    device_id = EXCLUDED.device_id,
    parents = EXCLUDED.parents,
    grain_rate = EXCLUDED.grain_rate,
    format = EXCLUDED.format,
    media_type = EXCLUDED.media_type,
    frame_width = EXCLUDED.frame_width,
    frame_height = EXCLUDED.frame_height,
    interlace_mode = EXCLUDED.interlace_mode,
    colorspace = EXCLUDED.colorspace,
    transfer_characteristic = EXCLUDED.transfer_characteristic,
    components = EXCLUDED.components,
    sample_rate = EXCLUDED.sample_rate,
    bit_depth = EXCLUDED.bit_depth,
    event_type = EXCLUDED.event_type,
    did_sdid = EXCLUDED.did_sdid
WHERE
    flows.meta_api_version = EXCLUDED.meta_api_version
    AND flows.resource_version < EXCLUDED.resource_version
RETURNING *;