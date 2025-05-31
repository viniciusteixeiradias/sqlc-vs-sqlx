package store

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Task struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Title  string `db:"title"`
	Done   bool   `db:"done"`
}
