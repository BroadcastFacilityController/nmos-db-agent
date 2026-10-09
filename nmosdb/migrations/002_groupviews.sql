-- +goose Up
CREATE VIEW receiver_groups AS
SELECT DISTINCT
    r.device_id,
    split_part(t.grouphint, ':', 1) AS group_name,
    split_part(t.grouphint, ':', 2) AS group_member,
    r.id AS receiver_id
FROM receivers r
CROSS JOIN LATERAL jsonb_array_elements_text(
    r.tags -> 'urn:x-nmos:tag:grouphint/v1.0'
) AS t(grouphint);

CREATE VIEW sender_groups AS
SELECT DISTINCT
    s.device_id,
    split_part(t.grouphint, ':', 1) AS group_name,
    split_part(t.grouphint, ':', 2) AS group_member,
    s.id AS sender_id
FROM senders s
CROSS JOIN LATERAL jsonb_array_elements_text(
    s.tags -> 'urn:x-nmos:tag:grouphint/v1.0'
) AS t(grouphint);

-- +goose Down
DROP VIEW receiver_groups;
DROP VIEW sender_groups;
