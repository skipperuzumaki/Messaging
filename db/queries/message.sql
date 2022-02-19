-- name: createMessage :one
INSERT INTO "message"(
    "group",
    "message",
    "sent_from",
    "sent_to"
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetLatestUnreadMessage :one
SELECT * FROM "message"
WHERE "sent_to" = $1 AND "read" = FALSE
ORDER BY "sent_at" desc
LIMIT 1;

-- name: ReadMessageGroup :exec
UPDATE "message"
SET "read" = TRUE
WHERE "group" = $1;

-- name: RetrieveAll :many
SELECT * FROM "message"
where "group" = $1
ORDER BY "sent_at"
LIMIT $2
OFFSET $3;