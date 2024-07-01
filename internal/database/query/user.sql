-- name: GetAllUsers :many
SELECT id, first_name, last_name, email, phone, address
FROM "user";

-- name: GetUserById :one
SELECT id, first_name, last_name, email, phone, address
FROM "user"
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, phone, address
FROM "user"
WHERE email = $1
LIMIT 1;

-- name: GetUserByPhone :one
SELECT id, first_name, last_name, email, phone, address
FROM "user"
WHERE phone = $1
LIMIT 1;

-- name: GetUserByEmailAndPassword :one
SELECT id, first_name, last_name, email, phone, address
FROM "user"
WHERE email = $1
  AND password_hash = $2
LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user" (first_name, last_name, email, password_hash, phone, address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, first_name, last_name, email, phone, address, created_at, updated_at;

-- name: UpdateUserName :one
UPDATE "user"
SET first_name = $2,
    last_name  = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, first_name, last_name, email, phone, address, created_at, updated_at;

-- name: UpdateUserEmail :one
UPDATE "user"
SET email = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, first_name, last_name, email, phone, address, created_at, updated_at;

-- name: UpdateUserPhone :one
UPDATE "user"
SET phone = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, first_name, last_name, email, phone, address, created_at, updated_at;

-- name: UpdateUserAddress :one
UPDATE "user"
SET address = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, first_name, last_name, email, phone, address, created_at, updated_at;

-- name: UpdateUserPassword :one
UPDATE "user"
SET password_hash = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, first_name, last_name, email, phone, address, created_at, updated_at;

-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE id = $1;
