-- name: GetNode :one
SELECT * FROM nodes
WHERE id = $1 LIMIT 1;

-- name: GetNodeByDevice :one
SELECT * FROM nodes N
WHERE N.id = (
    SELECT node_id FROM devices D
    WHERE D.id = $1 LIMIT 1
) LIMIT 1;

-- name: GetNodeByReceiver :one
SELECT * FROM nodes N
WHERE N.id = (
    SELECT node_id FROM devices D
    WHERE D.id = (
        SELECT device_id FROM receivers R
        WHERE R.id = $1 LIMIT 1
    ) LIMIT 1
) LIMIT 1;

-- name: ListNodes :many
SELECT * FROM nodes
ORDER BY id;

-- name: ListNodesPaginated :many
SELECT * FROM nodes
WHERE id >= @starting_id
ORDER BY id
LIMIT @page_size;

-- name: ListNodesByCreatedSince :many
SELECT * FROM nodes
WHERE meta_created_at >= @created_since
ORDER BY meta_created_at;

-- name: UpsertNode :one
INSERT INTO nodes (
    id, resource_version, label, description, tags, api_versions, api_endpoints, caps, services, clocks, interfaces, meta_user_label, meta_api_version
) VALUES (
    @id, @resource_version, @label, @description, @tags, @api_versions, @api_endpoints, @caps, @services, @clocks, @interfaces, @meta_user_label, @meta_api_version
)
ON CONFLICT (id) 
DO UPDATE
SET
    resource_version = EXCLUDED.resource_version,
    label = EXCLUDED.label,
    description = EXCLUDED.description,
    tags = EXCLUDED.tags,
    api_versions = EXCLUDED.api_versions,
    api_endpoints = EXCLUDED.api_endpoints,
    caps = EXCLUDED.caps,
    services = EXCLUDED.services,
    clocks = EXCLUDED.clocks,
    interfaces = EXCLUDED.interfaces,
    meta_user_label = EXCLUDED.meta_user_label
WHERE
    nodes.meta_api_version = EXCLUDED.meta_api_version
    AND nodes.resource_version < EXCLUDED.resource_version
RETURNING *;