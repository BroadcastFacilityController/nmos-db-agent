-- name: ListSenderGroupsByDeviceID :many
SELECT
    sg.device_id,
    sg.group_name,
    sg.group_member,
    sqlc.embed(s)
FROM sender_groups sg
JOIN senders s
    ON s.id = sg.sender_id
WHERE
    sg.device_id = @device_id;

-- name: ListSenderGroupsBySenderID :many
SELECT sg.*
FROM sender_groups sg
WHERE
    sg.sender_id = @sender_id;