-- name: GetAllRoles :many
SELECT id, name
FROM "role";

-- name: GetRoleById :one
SELECT id, name
FROM "role"
WHERE id = $1
    LIMIT 1;

-- name: GetRoleByName :one
SELECT id, name
FROM "role"
WHERE name = $1
    LIMIT 1;

-- name: CreateRole :one
INSERT INTO "role" (name)
VALUES ($1)
    RETURNING id, name, created_at, updated_at;

-- name: UpdateRoleName :one
UPDATE "role"
SET name = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id, name, created_at, updated_at;

-- name: DeleteRole :exec
DELETE
FROM "role"
WHERE id = $1;
