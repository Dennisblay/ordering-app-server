-- name: GetAllSocialMedia :many
SELECT id, user_id, platform, link
FROM social_media;

-- name: GetSocialMediaById :one
SELECT id, user_id, platform, link
FROM social_media
WHERE id = $1
    LIMIT 1;

-- name: GetSocialMediaByUserId :many
SELECT id, user_id, platform, link
FROM social_media
WHERE user_id = $1;

-- name: CreateSocialMedia :one
INSERT INTO social_media (user_id, platform, link)
VALUES ($1, $2, $3)
    RETURNING id, user_id, platform, link, created_at, updated_at;

-- name: UpdateSocialMedia :one
UPDATE social_media
SET platform = $2,
    link = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, platform, link, created_at, updated_at;

-- name: DeleteSocialMedia :exec
DELETE
FROM social_media
WHERE id = $1;
