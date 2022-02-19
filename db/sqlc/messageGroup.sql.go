// Code generated by sqlc. DO NOT EDIT.
// source: messageGroup.sql

package db

import (
	"context"
)

const deleteMessageGroup = `-- name: DeleteMessageGroup :exec
DELETE FROM "messageGroup" WHERE "id" = $1
`

func (q *Queries) DeleteMessageGroup(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMessageGroup, id)
	return err
}

const getMessageGroup = `-- name: GetMessageGroup :one
SELECT id, identifier, from_user, to_user FROM "messageGroup"
WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetMessageGroup(ctx context.Context, id int64) (MessageGroup, error) {
	row := q.db.QueryRowContext(ctx, getMessageGroup, id)
	var i MessageGroup
	err := row.Scan(
		&i.ID,
		&i.Identifier,
		&i.FromUser,
		&i.ToUser,
	)
	return i, err
}

const getMessageGroupFromIdentifier = `-- name: GetMessageGroupFromIdentifier :one
SELECT id, identifier, from_user, to_user FROM "messageGroup"
WHERE "identifier" = $1 LIMIT 1
`

func (q *Queries) GetMessageGroupFromIdentifier(ctx context.Context, identifier string) (MessageGroup, error) {
	row := q.db.QueryRowContext(ctx, getMessageGroupFromIdentifier, identifier)
	var i MessageGroup
	err := row.Scan(
		&i.ID,
		&i.Identifier,
		&i.FromUser,
		&i.ToUser,
	)
	return i, err
}

const listGroups = `-- name: ListGroups :many
SELECT id, identifier, from_user, to_user FROM "messageGroup"
WHERE "from_user" = $1 OR "to_user" = $1
ORDER BY "id"
`

func (q *Queries) ListGroups(ctx context.Context, fromUser int64) ([]MessageGroup, error) {
	rows, err := q.db.QueryContext(ctx, listGroups, fromUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MessageGroup
	for rows.Next() {
		var i MessageGroup
		if err := rows.Scan(
			&i.ID,
			&i.Identifier,
			&i.FromUser,
			&i.ToUser,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createMessageGroup = `-- name: createMessageGroup :one
INSERT INTO "messageGroup"(
    "identifier",
    "from_user",
    "to_user"
) VALUES (
    $1,
    $2,
    $3
) RETURNING id, identifier, from_user, to_user
`

type createMessageGroupParams struct {
	Identifier string `json:"identifier"`
	FromUser   int64  `json:"from_user"`
	ToUser     int64  `json:"to_user"`
}

func (q *Queries) createMessageGroup(ctx context.Context, arg createMessageGroupParams) (MessageGroup, error) {
	row := q.db.QueryRowContext(ctx, createMessageGroup, arg.Identifier, arg.FromUser, arg.ToUser)
	var i MessageGroup
	err := row.Scan(
		&i.ID,
		&i.Identifier,
		&i.FromUser,
		&i.ToUser,
	)
	return i, err
}
