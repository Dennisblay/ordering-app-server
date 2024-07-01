-- name: GetAllUserRoles :many
SELECT user_id, role_id
FROM "user_role";

-- name: GetUserRolesByUserId :many
SELECT user_id, role_id
FROM "user_role"
WHERE user_id = $1;

-- name: GetUserRolesByRoleId :many
SELECT user_id, role_id
FROM "user_role"
WHERE role_id = $1;

-- name: CreateUserRole :one
INSERT INTO "user_role" (user_id, role_id)
VALUES ($1, $2)
    RETURNING user_id, role_id;

-- name: DeleteUserRole :exec
DELETE
FROM "user_role"
WHERE user_id = $1 AND role_id = $2;
