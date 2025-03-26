-- name: AddPdf :one
INSERT INTO pdfs (uuid, user_id, name)
VALUES (?, ?, ?)
RETURNING *;

-- name: DeletePdf :exec
DELETE FROM pdfs
WHERE id = ?;

-- name: GetPdfByUUID :one
SELECT * FROM pdfs
WHERE uuid = ?;

-- name: GetPdfsByUserID :many
SELECT * FROM pdfs
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: UpdateLastOpenTime :exec
UPDATE pdfs
SET last_open_at = CURRENT_TIMESTAMP
WHERE id = ?;