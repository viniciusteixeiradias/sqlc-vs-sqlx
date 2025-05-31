package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/viniciusteixeiradias/sqlc/internal/db"
	"log"
	"time"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable"

	dbConn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	queries := db.New(dbConn)
	ctx := context.Background()

	// --- Create a new user ---
	user, err := queries.CreateUser(ctx, "Vinicius")

	if err != nil {
		log.Fatal("Error creating user:", err)
	}

	fmt.Printf("User created: %+v\n", user)

	// --- Create a task ---
	task, err := queries.CreateTask(ctx, db.CreateTaskParams{
		UserID: user.ID,
		Title:  "Learn SQLC",
	})

	if err != nil {
		log.Fatal("Error creating task:", err)
	}

	fmt.Printf("Task created: %+v\n", task)

	// --- List tasks for the user ---
	tasks, err := queries.ListTasksByUser(ctx, user.ID)

	if err != nil {
		log.Fatal("Error listing tasks:", err)
	}

	fmt.Printf("Tasks for user %d:\n", user.ID)

	for _, t := range tasks {
		fmt.Printf("- Task ID: %d, Title: %s, Done: %t\n", t.ID, t.Title, t.Done)
	}

	// --- Mark task as done ---
	time.Sleep(1 * time.Second) // just for logs
	err = queries.MarkTaskDone(ctx, task.ID)

	if err != nil {
		log.Fatal("Error marking task as done:", err)
	}

	fmt.Printf("Task %d marked as done.\n", task.ID)
}
