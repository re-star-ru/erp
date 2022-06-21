// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package repo

import (
	"context"
)

const createRestaritem = `-- name: CreateRestaritem :one
INSERT INTO restaritems (
	name, onceGUID
) VALUES (
			 $1, $2
		 )
RETURNING id, onceguid, name, sku, itemguid
`

type CreateRestaritemParams struct {
	Name     string
	Onceguid string
}

func (q *Queries) CreateRestaritem(ctx context.Context, arg CreateRestaritemParams) (Restaritem, error) {
	row := q.db.QueryRowContext(ctx, createRestaritem, arg.Name, arg.Onceguid)
	var i Restaritem
	err := row.Scan(
		&i.ID,
		&i.Onceguid,
		&i.Name,
		&i.Sku,
		&i.Itemguid,
	)
	return i, err
}

const deleteRestaritem = `-- name: DeleteRestaritem :exec
DELETE FROM restaritems
WHERE id = $1
`

func (q *Queries) DeleteRestaritem(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteRestaritem, id)
	return err
}

const getRestaritem = `-- name: GetRestaritem :one
SELECT id, onceguid, name, sku, itemguid FROM restaritems
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRestaritem(ctx context.Context, id int64) (Restaritem, error) {
	row := q.db.QueryRowContext(ctx, getRestaritem, id)
	var i Restaritem
	err := row.Scan(
		&i.ID,
		&i.Onceguid,
		&i.Name,
		&i.Sku,
		&i.Itemguid,
	)
	return i, err
}

const listRestaritems = `-- name: ListRestaritems :many
SELECT id, onceguid, name, sku, itemguid FROM restaritems
ORDER BY id
`

func (q *Queries) ListRestaritems(ctx context.Context) ([]Restaritem, error) {
	rows, err := q.db.QueryContext(ctx, listRestaritems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Restaritem
	for rows.Next() {
		var i Restaritem
		if err := rows.Scan(
			&i.ID,
			&i.Onceguid,
			&i.Name,
			&i.Sku,
			&i.Itemguid,
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

const updateRestaritem = `-- name: UpdateRestaritem :one
UPDATE restaritems
set name = $2,
	onceGUID = $3
WHERE id = $1
RETURNING id, onceguid, name, sku, itemguid
`

type UpdateRestaritemParams struct {
	ID       int64
	Name     string
	Onceguid string
}

func (q *Queries) UpdateRestaritem(ctx context.Context, arg UpdateRestaritemParams) (Restaritem, error) {
	row := q.db.QueryRowContext(ctx, updateRestaritem, arg.ID, arg.Name, arg.Onceguid)
	var i Restaritem
	err := row.Scan(
		&i.ID,
		&i.Onceguid,
		&i.Name,
		&i.Sku,
		&i.Itemguid,
	)
	return i, err
}