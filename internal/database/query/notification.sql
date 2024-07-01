-- name: GetAllNotifications :many
SELECT id, user_id, message, sent_at
FROM notification;

-- name: GetNotificationById :one
SELECT id, user_id, message, sent_at
FROM notification
WHERE id = $1
    LIMIT 1;

-- name: GetNotificationsByUserId :many
SELECT id, user_id, message, sent_at
FROM notification
WHERE user_id = $1;

-- name: CreateNotification :one
INSERT INTO notification (user_id, message, sent_at)
VALUES ($1, $2, $3)
    RETURNING id, user_id, message, sent_at, created_at, updated_at;

-- name: UpdateNotification :one
UPDATE notification
SET message = $2,
    sent_at = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, message, sent_at, created_at, updated_at;

-- name: DeleteNotification :exec
DELETE
FROM notification
WHERE id = $1;
