-- name: createMessageGroup :one
INSERT INTO "messageGroup"(
    identifier,
    from_user,
    to_user
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetMessageGroup :one
SELECT * FROM "messageGroup"
WHERE id = $1 LIMIT 1;

-- name: GetMessageGroupFromIdentifier :one
SELECT * FROM "messageGroup"
where identifier = $1 LIMIT 1;

-- name: ListGroups :many
SELECT * FROM "messageGroup"
where from_user = $1 OR
where to_user = $1
ORDER BY id;

-- name: DeleteMessageGroup :exec
DELETE FROM "messageGroup" WHERE id = $1;