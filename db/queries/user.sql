-- name: createUser :one
INSERT INTO "users"(
    username
) VALUES (
    $1
) RETURNING *;
