-- name: GetAllSessions :many
SELECT id, user_id, token, expires_at
FROM "session";

-- name: GetSessionById :one
SELECT id, user_id, token, expires_at
FROM "session"
WHERE id = $1
    LIMIT 1;

-- name: GetSessionByToken :one
SELECT id, user_id, token, expires_at
FROM "session"
WHERE token = $1
    LIMIT 1;

-- name: CreateSession :one
INSERT INTO "session" (user_id, token, expires_at)
VALUES ($1, $2, $3)
    RETURNING id, user_id, token, expires_at, created_at, updated_at;

-- name: DeleteSession :exec
DELETE
FROM "session"
WHERE id = $1;
