-- name: GetAllUsers :many
SELECT *
FROM users;
-- name: AddUser :one
INSERT INTO users (
        uuid,
        name,
        email,
        folder_uuid,
        is_draft,
        created_at
    )
VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
RETURNING *;
-- name: GetUserByID :one
SELECT *
from users
WHERE id = ?;
-- name: GetUserByUUID :one
SELECT *
from users
WHERE uuid = ?;
-- name: GetUserByEmail :one
SELECT *
from users
WHERE email = ?;

-- name: UpdateUserEmail :one
UPDATE users
SET email = ?
WHERE id = ?
RETURNING *;
-- name: UpdateUserIsDraft :one
UPDATE users
SET is_draft = ?
WHERE id = ?
RETURNING *;
-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = ?;
-- name: AddSession :one
INSERT INTO sessions (
        user_id,
        email,
        expiry,
        token,
        refresh_token,
        activate_code,
        user_ip,
        type,
        post_suspend_expiry,
        is_expired
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;
-- name: GetSessionByToken :one
SELECT *
FROM sessions
WHERE token = ?;
-- name: GetSessionByActivateCode :one
SELECT *
FROM sessions
WHERE activate_code = ?;
-- name: UpdateSessionPostSuspendExpiry :exec
UPDATE sessions
SET post_suspend_expiry = ?
WHERE token = ?;
-- name: UpdateSessionIsExpired :exec
UPDATE sessions
SET is_expired = ?
WHERE token = ?;