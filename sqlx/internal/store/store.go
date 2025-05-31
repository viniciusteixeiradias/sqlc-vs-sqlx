package store

import "github.com/jmoiron/sqlx"

type Store interface {
	CreateUser(name string) (User, error)
	CreateTask(userID int, title string) (Task, error)
	ListTasksByUser(userID int) ([]Task, error)
	MarkTaskDone(taskID int) error
}

type SQLStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &SQLStore{db: db}
}

func (s *SQLStore) CreateUser(name string) (User, error) {
	var user User
	err := s.db.Get(&user, "INSERT INTO users (name) VALUES ($1) RETURNING *", name)
	return user, err
}

func (s *SQLStore) CreateTask(userID int, title string) (Task, error) {
	var task Task
	err := s.db.Get(&task, "INSERT INTO tasks (user_id, title) VALUES ($1, $2) RETURNING *", userID, title)
	return task, err
}

func (s *SQLStore) ListTasksByUser(userID int) ([]Task, error) {
	var tasks []Task
	err := s.db.Select(&tasks, "SELECT * FROM tasks WHERE user_id = $1", userID)
	return tasks, err
}

func (s *SQLStore) MarkTaskDone(taskID int) error {
	_, err := s.db.Exec("UPDATE tasks SET done = TRUE WHERE id = $1", taskID)
	return err
}
