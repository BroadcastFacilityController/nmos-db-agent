-- name: ListSenderGroupsByDeviceID :many
SELECT
    sqlc.embed(d),
    sg.group_name,
    sg.group_member,
    sqlc.embed(s)
FROM sender_groups sg
JOIN senders s
    ON s.id = sg.sender_id
JOIN devices d
    ON d.id = sg.device_id
WHERE
    sg.device_id = @device_id;

-- name: ListSenderGroupsByDeviceIDAndGroupName :many
SELECT
    sqlc.embed(d),
    sg.group_name,
    sg.group_member,
    sqlc.embed(s)
FROM sender_groups sg
JOIN senders s
    ON s.id = sg.sender_id
JOIN devices d
    ON d.id = sg.device_id
WHERE
    sg.device_id = @device_id AND sg.group_name = @group_name;

-- name: ListSenderGroupsBySenderID :many
SELECT sg.*
FROM sender_groups sg
WHERE
    sg.sender_id = @sender_id;