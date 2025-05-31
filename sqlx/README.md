# Task Manager (SQLX Version)

This is a simple Go project using **PostgreSQL**, **sqlx**, and **golang-migrate** to demonstrate a lightweight and flexible approach to working with relational databases using raw SQL and Go structs.

---

## 📁 Project Structure

```bash
├── db/
│   └── migrations/           # SQL migration files (up/down)
├── internal/
│   ├── store/                # Store interface and implementation using sqlx
│   │   ├── models.go         # Structs: User, Task
│   │   └── store.go          # Store interface + SQLX implementation
│   └── main.go               # Example app using the store
├── go.mod / go.sum           # Go module definition
└── README.md                 # You're here!
```

---

## 🛠️ Setup & Commands

### 1. Install migration manager

Install [`golang-migrate`](https://github.com/golang-migrate/migrate) CLI locally:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# or macos: brew install golang-migrate
```

Make sure `$GOPATH/bin` is in your `PATH`.

---

### 2. Run migrations
```bash
migrate -path db/migrations \
  -database "postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable" \
  up
```

---

### 3. Run the app
```bash
# Optional: get dependencies
go mod tidy

# Run the program
go run internal/main.go
```

---

### 4. Run migrations (down)
```bash
migrate -path db/migrations \
  -database "postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable" \
  down
```
> Optional: Reset everything by running `migrate drop -f`

---

### ✅ Pros
1. Less Boilerplate
* No config files (like sqlc.yaml) or code generation required.
* You just write your Go code and SQL queries directly.

2. Really Simple
* Easy to learn and start using.
* Clear mental model: SQL in Go, structs match DB tables.

3. More Flexible
* You can build dynamic queries easily (e.g., building WHERE clauses based on logic).
* You're not tied to a generator or schema snapshot.

4. Direct SQL Control
* You write the exact SQL you want — no abstractions in the way.
* You can optimize complex queries manually.

5. No Extra Codegen Dependencies
* Unlike sqlc, there's no need to install or manage extra CLI tools or config files.
* Just install sqlx via go get and go.

### ❌ Cons
1. No Central Schema File
* There’s no schema.sql snapshot — to understand your DB structure, you must:
* * Read all migration files, or
* * Inspect the live DB, or
* * Use `pg_dump -s`.

2. No Compile-Time Type Safety
* SQL queries are written as strings and only validated at runtime.
* Mistakes like column typos or type mismatches won't be caught until you run the code.

3. Manual Struct Maintenance
* You have to manually ensure your Go structs match your DB schema.
* If the schema changes, it’s easy to forget to update your struct.

4. No Interface for Queries by Default
* Without wrapping sqlx in a Store interface (like I did here), the code can get messy in large projects.
