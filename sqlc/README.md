# Task Manager (SQLC Version)

This is a simple Go project using **PostgreSQL**, **sqlc**, and **golang-migrate** to demonstrate a type-safe and organized approach to working with databases.

---

## 📁 Project Structure

```bash
├── db/
│ ├── migrations/ # SQL migration files (up/down)
│ ├── queries.sql # Raw SQL queries used by sqlc
│ └── schema.sql # Full schema snapshot (used by sqlc)
│
├── internal/
│ ├── db/ # Auto-generated Go code (from sqlc)
│ └── main.go # Example app using the queries
│
├── sqlc.yaml # sqlc config file
├── go.mod / go.sum # Go module definition
└── README.md # You're here!
```


---

## 🛠️ Setup & Commands

### 1. Install migration manager

Install [`golang-migrate`](https://github.com/golang-migrate/migrate) CLI locally:

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# or macos: brew install golang-migrate
```

Make sure `$GOPATH/bin` is in your `PATH`.

---

### 2. Generate Go code from SQL
Install `sqlc` if you haven't already:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Then run the following command to generate Go code from your SQL queries:

```bash
sqlc generate
```
> This will read the `sqlc.yaml` config file and generate Go code in the `internal/db` directory based on the SQL queries defined in `db/queries.sql`.

---

### 3. Run migrations
```bash
migrate -path db/migrations \
  -database "postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable" \
  up
```

---

### 4. Run the app
```bash
# Optional: get dependencies
go mod tidy

# Run the program
go run internal/main.go
```

---

### 5. Run migrations (down)
```bash
migrate -path db/migrations \
  -database "postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable" \
  down
```
> Optional: Reset everything by running `migrate drop -f`

---

### ✅ Pros
1. Type Safety
* All SQL query parameters and return types are type-checked at compile time.
* No risk of mismatched types or missing fields in your Go structs.

2. Organization
* Clean separation between raw SQL, Go code, and database migrations.
* Easy to reason about changes and track DB evolution.

### ❌ Cons
1. Boilerplate
* You need to write extra .sql files and manage a sqlc.yaml file.
* Requires a bit more setup than raw SQL with database/sql.

2. Duplication
* You write your schema once in migrations/*.up.sql...
* ...and again (as a snapshot) in schema.sql for sqlc code generation.

3. Manual Dependency
* sqlc doesn't handle DB migrations for you.
* You must manage tools like golang-migrate separately.
