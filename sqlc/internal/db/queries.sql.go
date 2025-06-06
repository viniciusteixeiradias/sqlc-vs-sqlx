// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: queries.sql

package db

import (
	"context"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (user_id, title)
VALUES ($1, $2)
RETURNING id, user_id, title, done
`

type CreateTaskParams struct {
	UserID int32
	Title  string
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.UserID, arg.Title)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Done,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, user_id, title, done FROM tasks
WHERE id = $1
`

// new
func (q *Queries) GetTaskByID(ctx context.Context, id int32) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Done,
	)
	return i, err
}

const listTasksByUser = `-- name: ListTasksByUser :many
SELECT id, user_id, title, done FROM tasks
WHERE user_id = $1
`

func (q *Queries) ListTasksByUser(ctx context.Context, userID int32) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Done,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markTaskDone = `-- name: MarkTaskDone :exec
UPDATE tasks
SET done = TRUE
WHERE id = $1
`

func (q *Queries) MarkTaskDone(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, markTaskDone, id)
	return err
}
