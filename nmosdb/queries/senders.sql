-- name: GetSender :one
SELECT * FROM senders
WHERE id = $1 LIMIT 1;

-- name: ListSenders :many
SELECT * FROM senders
ORDER BY id;

-- name: ListSendersByTransport :many
SELECT * FROM senders
WHERE transport = $1
ORDER BY id;

-- name: ListSendersByCreatedSince :many
SELECT * FROM senders
WHERE meta_created_at >= @created_since
ORDER BY meta_created_at;

-- name: ListSendersByTransportAndFlowFormat :many
SELECT s.* FROM senders s
JOIN flows f ON s.flow_id = f.id
JOIN sources src ON f.source_id = src.id
WHERE s.transport = @transport AND f.format = @flow_format AND src.format = @source_format
ORDER BY s.id;

-- name: ListSendersByDeviceIDTransportAndFormat :many
SELECT s.* FROM senders s
JOIN flows f ON s.flow_id = f.id
JOIN sources src ON f.source_id = src.id
WHERE s.device_id = @device_id AND s.transport = @transport AND f.format = @flow_format AND src.format = @source_format
ORDER BY s.id;

-- name: UpsertSender :one
INSERT INTO senders (
    id, resource_version, label, description, tags, caps, flow_id, transport, device_id, manifest_href, interface_bindings, subscription_receiver, subscription_active, transport_file, meta_api_version
) VALUES (
    @id, @resource_version, @label, @description, @tags, @caps, @flow_id, @transport, @device_id, @manifest_href, @interface_bindings, @subscription_receiver, @subscription_active, @transport_file, @meta_api_version
)
ON CONFLICT (id) 
DO UPDATE
SET
    resource_version = EXCLUDED.resource_version,
    label = EXCLUDED.label,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    caps = EXCLUDED.caps,
    flow_id = EXCLUDED.flow_id,
    transport = EXCLUDED.transport,
    device_id = EXCLUDED.device_id,
    manifest_href = EXCLUDED.manifest_href,
    interface_bindings = EXCLUDED.interface_bindings,
    subscription_receiver = EXCLUDED.subscription_receiver,
    subscription_active = EXCLUDED.subscription_active,
    transport_file = EXCLUDED.transport_file
WHERE
    senders.meta_api_version = EXCLUDED.meta_api_version
    AND senders.resource_version < EXCLUDED.resource_version
RETURNING *;