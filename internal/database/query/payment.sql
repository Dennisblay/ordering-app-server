-- name: GetAllPayments :many
SELECT id, order_id, amount, payment_method, payment_date
FROM payment;

-- name: GetPaymentById :one
SELECT id, order_id, amount, payment_method, payment_date
FROM payment
WHERE id = $1
    LIMIT 1;

-- name: GetPaymentsByOrderId :many
SELECT id, order_id, amount, payment_method, payment_date
FROM payment
WHERE order_id = $1;

-- name: CreatePayment :one
INSERT INTO payment (order_id, amount, payment_method, payment_date)
VALUES ($1, $2, $3, $4)
    RETURNING id, order_id, amount, payment_method, payment_date, created_at, updated_at;

-- name: UpdatePayment :one
UPDATE payment
SET amount = $2,
    payment_method = $3,
    payment_date = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, order_id, amount, payment_method, payment_date, created_at, updated_at;

-- name: DeletePayment :exec
DELETE
FROM payment
WHERE id = $1;
