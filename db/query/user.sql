-- name: CreateUser :one
INSERT INTO users (
    name
) VALUES (
    $1
) RETURNING *;

-- name: DisplayUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: DisplayAllUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;