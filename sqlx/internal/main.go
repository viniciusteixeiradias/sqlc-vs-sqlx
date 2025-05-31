package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/viniciusteixeiradias/sqlx/internal/store"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable"

	dbConn, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	s := store.NewStore(dbConn)

	// --- Create a new user ---
	user, err := s.CreateUser("Vinicius")

	if err != nil {
		log.Fatal("Error creating user:", err)
	}

	fmt.Printf("User created: %+v\n", user)

	// --- Create a task ---
	task, err := s.CreateTask(user.ID, "Learn SQLX")

	if err != nil {
		log.Fatal("Error creating task:", err)
	}

	fmt.Printf("Task created: %+v\n", task)

	// --- List tasks for the user ---
	tasks, err := s.ListTasksByUser(user.ID)

	if err != nil {
		log.Fatal("Error listing tasks:", err)
	}

	fmt.Printf("Tasks for user %d:\n", user.ID)

	for _, t := range tasks {
		fmt.Printf("- Task ID: %d, Title: %s, Done: %t\n", t.ID, t.Title, t.Done)
	}

	// --- Mark task as done ---
	time.Sleep(1 * time.Second) // just for logs

	err = s.MarkTaskDone(task.ID)

	if err != nil {
		log.Fatal("Error marking task as done:", err)
	}

	fmt.Printf("Task %d marked as done.\n", task.ID)
}
