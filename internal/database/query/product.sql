-- name: GetAllProducts :many
SELECT id, name, description, category, price, stock
FROM product;

-- name: GetProductById :one
SELECT id, name, description, category, price, stock
FROM product
WHERE id = $1
    LIMIT 1;

-- name: CreateProduct :one
INSERT INTO product (name, description, category, price, stock)
VALUES ($1, $2, $3, $4, $5)
    RETURNING id, name, description, category, price, stock, created_at, updated_at;

-- name: UpdateProduct :one
UPDATE product
SET name = $2,
    description = $3,
    category = $4,
    price = $5,
    stock = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, name, description, category, price, stock, created_at, updated_at;

-- name: DeleteProduct :exec
DELETE
FROM product
WHERE id = $1;
