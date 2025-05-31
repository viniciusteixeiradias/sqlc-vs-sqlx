-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING *;

-- name: CreateTask :one
INSERT INTO tasks (user_id, title)
VALUES ($1, $2)
RETURNING *;

-- name: ListTasksByUser :many
SELECT * FROM tasks
WHERE user_id = $1;

-- name: MarkTaskDone :exec
UPDATE tasks
SET done = TRUE
WHERE id = $1;

-- new
-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1;
