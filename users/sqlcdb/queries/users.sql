-- name: InsertUser :one
INSERT INTO users(id, first_name, last_name, created_at)
VALUES (gen_random_uuid(), $1, $2, now())
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserV2 :many
SELECT id, id
FROM users;
