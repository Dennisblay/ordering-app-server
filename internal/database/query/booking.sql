-- name: GetAllBookings :many
SELECT id, user_id, event_date, event_type, guest_count, location
FROM booking;

-- name: GetBookingById :one
SELECT id, user_id, event_date, event_type, guest_count, location
FROM booking
WHERE id = $1
    LIMIT 1;

-- name: GetBookingsByUserId :many
SELECT id, user_id, event_date, event_type, guest_count, location
FROM booking
WHERE user_id = $1;

-- name: CreateBooking :one
INSERT INTO booking (user_id, event_date, event_type, guest_count, location)
VALUES ($1, $2, $3, $4, $5)
    RETURNING id, user_id, event_date, event_type, guest_count, location, created_at, updated_at;

-- name: UpdateBooking :one
UPDATE booking
SET event_date = $2,
    event_type = $3,
    guest_count = $4,
    location = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, event_date, event_type, guest_count, location, created_at, updated_at;

-- name: DeleteBooking :exec
DELETE
FROM booking
WHERE id = $1;
