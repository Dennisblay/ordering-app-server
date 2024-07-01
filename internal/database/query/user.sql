-- name: GetAllUsers :many
SELECT id, name, email, phone, address
FROM "user";

-- name: GetUserById :one
SELECT id, name, email, phone, address
FROM "user"
where id = $1
limit 1;

-- name: GetUserByEmail :one
SELECT id, name, email, phone, address
FROM "user"
where email = $1
limit 1;

-- name: GetUserByPhone :one
SELECT id, name, email, phone, address
FROM "user"
where phone = $1
limit 1;

-- name: GetUserByEmailAndPassword :one
SELECT id, name, email, phone, address
FROM "user"
where email = $1
  and password_hash = $2
LIMIT 1;
