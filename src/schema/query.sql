-- name: GetList :one
SELECT *
FROM list
WHERE id = ?
LIMIT 1;
-- name: ShowList :many
SELECT *
FROM list
ORDER BY name;
-- name: CreateList :one
INSERT INTO list (name, description)
VALUES (?, ?)
RETURNING *;
-- name: UpdateList :one
UPDATE list
set name = ?,
    description = ?
WHERE id = ?
RETURNING *;
-- name: Delete :exec
DELETE FROM list
WHERE id = ?;