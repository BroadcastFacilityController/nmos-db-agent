-- name: GetReceiver :one
SELECT * FROM receivers
WHERE id = $1 LIMIT 1;

-- name: ListReceivers :many
SELECT * FROM receivers
ORDER BY id;

-- name: ListReceiversPaginated :many
SELECT * FROM receivers
WHERE id >= @starting_id
ORDER BY id
LIMIT @page_size;

-- name: ListReceiversByTransport :many
SELECT * FROM receivers
WHERE transport = $1
ORDER BY id;

-- name: ListReceiversByCreatedSince :many
SELECT * FROM receivers
WHERE meta_created_at >= @created_since
ORDER BY meta_created_at;

-- name: ListReceiversByDeviceID :many
SELECT * FROM receivers
WHERE device_id = $1
ORDER BY id;

-- name: ListReceiversByDeviceIDAndTransport :many
SELECT * FROM receivers
WHERE device_id = $1 AND transport = $2
ORDER BY id;

-- name: UpsertReceiver :one
INSERT INTO receivers (
    id, resource_version, label, description, tags, device_id, transport, interface_bindings, subscription_sender, subscription_active, format, caps, meta_user_label, meta_api_version
) VALUES (
    @id, @resource_version, @label, @description, @tags, @device_id, @transport, @interface_bindings, @subscription_sender, @subscription_active, @format, @caps, @meta_user_label, @meta_api_version
)
ON CONFLICT (id) 
DO UPDATE
SET
    resource_version = EXCLUDED.resource_version,
    label = EXCLUDED.label,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    device_id = EXCLUDED.device_id,
    transport = EXCLUDED.transport,
    interface_bindings = EXCLUDED.interface_bindings,
    subscription_sender = EXCLUDED.subscription_sender,
    format = EXCLUDED.format,
    caps = EXCLUDED.caps,
    meta_user_label = EXCLUDED.meta_user_label
WHERE
    receivers.meta_api_version = EXCLUDED.meta_api_version
    AND receivers.resource_version < EXCLUDED.resource_version
RETURNING *;