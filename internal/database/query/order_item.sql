-- name: GetAllOrderItems :many
SELECT id, order_id, product_id, quantity, price
FROM order_item;

-- name: GetOrderItemById :one
SELECT id, order_id, product_id, quantity, price
FROM order_item
WHERE id = $1
    LIMIT 1;

-- name: GetOrderItemsByOrderId :many
SELECT id, order_id, product_id, quantity, price
FROM order_item
WHERE order_id = $1;

-- name: CreateOrderItem :one
INSERT INTO order_item (order_id, product_id, quantity, price)
VALUES ($1, $2, $3, $4)
    RETURNING id, order_id, product_id, quantity, price, created_at, updated_at;

-- name: UpdateOrderItem :one
UPDATE order_item
SET quantity = $2,
    price = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, order_id, product_id, quantity, price, created_at, updated_at;

-- name: DeleteOrderItem :exec
DELETE
FROM order_item
WHERE id = $1;
