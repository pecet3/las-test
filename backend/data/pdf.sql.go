// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: pdf.sql

package data

import (
	"context"
)

const addPdf = `-- name: AddPdf :one
INSERT INTO pdfs (uuid, user_id, name)
VALUES (?, ?, ?)
RETURNING id, uuid, user_id, name, created_at, last_open_at
`

type AddPdfParams struct {
	Uuid   string `json:"uuid"`
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

func (q *Queries) AddPdf(ctx context.Context, arg AddPdfParams) (Pdf, error) {
	row := q.db.QueryRowContext(ctx, addPdf, arg.Uuid, arg.UserID, arg.Name)
	var i Pdf
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.LastOpenAt,
	)
	return i, err
}

const deletePdf = `-- name: DeletePdf :exec
DELETE FROM pdfs
WHERE id = ?
`

func (q *Queries) DeletePdf(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePdf, id)
	return err
}

const getPdfByUUID = `-- name: GetPdfByUUID :one
SELECT id, uuid, user_id, name, created_at, last_open_at FROM pdfs
WHERE uuid = ?
`

func (q *Queries) GetPdfByUUID(ctx context.Context, uuid string) (Pdf, error) {
	row := q.db.QueryRowContext(ctx, getPdfByUUID, uuid)
	var i Pdf
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.LastOpenAt,
	)
	return i, err
}

const getPdfsByUserID = `-- name: GetPdfsByUserID :many
SELECT id, uuid, user_id, name, created_at, last_open_at FROM pdfs
WHERE user_id = ?
ORDER BY created_at DESC
`

func (q *Queries) GetPdfsByUserID(ctx context.Context, userID int64) ([]Pdf, error) {
	rows, err := q.db.QueryContext(ctx, getPdfsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pdf
	for rows.Next() {
		var i Pdf
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.UserID,
			&i.Name,
			&i.CreatedAt,
			&i.LastOpenAt,
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

const updateLastOpenTime = `-- name: UpdateLastOpenTime :exec
UPDATE pdfs
SET last_open_at = CURRENT_TIMESTAMP
WHERE id = ?
`

func (q *Queries) UpdateLastOpenTime(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, updateLastOpenTime, id)
	return err
}
