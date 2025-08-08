-- name: CreateUser :one
INSERT INTO users (
    id,
    user_name,
    first_name,
    last_name,
    email,
    phone
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_name;

-- name: UpdateUser :one
UPDATE users
SET
    user_name = $2,
    first_name = $3,
    last_name = $4,
    email = $5,
    phone = $6
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1;
