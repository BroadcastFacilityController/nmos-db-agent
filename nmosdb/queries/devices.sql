-- name: GetDevice :one
SELECT * FROM devices
WHERE id = $1 LIMIT 1;

-- name: GetDeviceByReceiver :one
SELECT * FROM devices D
WHERE D.id = (
    SELECT device_id FROM receivers R
    WHERE R.id = $1 LIMIT 1
) LIMIT 1;

-- name: ListDevices :many
SELECT * FROM devices
ORDER BY id;

-- name: ListDevicesPaginated :many
SELECT * FROM devices
WHERE id >= @starting_id
ORDER BY id
LIMIT @page_size;

-- name: ListDevicesByNodeID :many
SELECT * FROM devices
WHERE node_id = $1
ORDER BY id;

-- name: ListDevicesByCreatedSince :many
SELECT * FROM devices
WHERE meta_created_at >= @created_since
ORDER BY meta_created_at;

-- name: UpdateDeviceUserLabel :exec
UPDATE devices
SET meta_user_label = @meta_user_label
WHERE id = @id;

-- name: UpsertDevice :one
INSERT INTO devices (
    id, resource_version, label, description, tags, type, receivers, senders, node_id, controls, meta_api_version
) VALUES (
    @id, @resource_version, @label, @description, @tags, @type, @receivers, @senders, @node_id, @controls, @meta_api_version
)
ON CONFLICT (id) 
DO UPDATE
SET
    resource_version = EXCLUDED.resource_version,
    label = EXCLUDED.label,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    type = EXCLUDED.type,
    receivers = EXCLUDED.receivers,
    senders = EXCLUDED.senders,
    node_id = EXCLUDED.node_id,
    controls = EXCLUDED.controls
WHERE
    devices.meta_api_version = EXCLUDED.meta_api_version
    AND devices.resource_version < EXCLUDED.resource_version
RETURNING *;