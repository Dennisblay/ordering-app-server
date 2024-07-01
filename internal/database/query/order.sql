-- name: GetAllOrders :many
SELECT id, user_id, total_price, status, delivery_date
FROM "order";

-- name: GetOrderById :one
SELECT id, user_id, total_price, status, delivery_date
FROM "order"
WHERE id = $1
    LIMIT 1;

-- name: GetOrdersByUserId :many
SELECT id, user_id, total_price, status, delivery_date
FROM "order"
WHERE user_id = $1;

-- name: CreateOrder :one
INSERT INTO "order" (user_id, total_price, status, delivery_date)
VALUES ($1, $2, $3, $4)
    RETURNING id, user_id, total_price, status, delivery_date, created_at, updated_at;

-- name: UpdateOrderStatus :one
UPDATE "order"
SET status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, user_id, total_price, status, delivery_date, created_at, updated_at;

-- name: DeleteOrder :exec
DELETE
FROM "order"
WHERE id = $1;
