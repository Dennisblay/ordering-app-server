-- name: GetAllTestimonials :many
SELECT id, user_id, message
FROM testimonial;

-- name: GetTestimonialById :one
SELECT id, user_id, message
FROM testimonial
WHERE id = $1
    LIMIT 1;

-- name: GetTestimonialsByUserId :many
SELECT id, user_id, message
FROM testimonial
WHERE user_id = $1;

-- name: CreateTestimonial :one
INSERT INTO testimonial (user_id, message)
VALUES ($1, $2)
    RETURNING id, user_id, message, created_at, updated_at;

-- name: UpdateTestimonial :one
UPDATE testimonial
SET message = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, message, created_at, updated_at;

-- name: DeleteTestimonial :exec
DELETE
FROM testimonial
WHERE id = $1;
