-- name: ListReceiverGroupsByDeviceID :many
SELECT
    sqlc.embed(d),
    rg.group_name,
    rg.group_member,
    sqlc.embed(r)
FROM receiver_groups rg
JOIN receivers r
    ON r.id = rg.receiver_id
JOIN devices d
    ON d.id = rg.device_id
WHERE
    rg.device_id = @device_id;

-- name: ListReceiverGroupsBySenderID :many
SELECT rg.*
FROM receiver_groups rg
WHERE
    rg.receiver_id = @receiver_id;