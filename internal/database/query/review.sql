-- name: GetAllReviews :many
SELECT id, user_id, product_id, rating, comment
FROM review;

-- name: GetReviewById :one
SELECT id, user_id, product_id, rating, comment
FROM review
WHERE id = $1
    LIMIT 1;

-- name: GetReviewsByProductId :many
SELECT id, user_id, product_id, rating, comment
FROM review
WHERE product_id = $1;

-- name: CreateReview :one
INSERT INTO review (user_id, product_id, rating, comment)
VALUES ($1, $2, $3, $4)
    RETURNING id, user_id, product_id, rating, comment, created_at, updated_at;

-- name: UpdateReview :one
UPDATE review
SET rating = $2,
    comment = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, product_id, rating, comment, created_at, updated_at;

-- name: DeleteReview :exec
DELETE
FROM review
WHERE id = $1;
