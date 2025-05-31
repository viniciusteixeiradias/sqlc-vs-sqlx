# Task Manager Project

## 📂 Project
This project demonstrates two approaches to using Go with PostgreSQL:
* [SQLC Version](./sqlc/README.md) — Type-safe, code-generated queries
* [SQLX Version](./sqlx/README.md) — Flexible, raw SQL with minimal setup

---

### 🔗 Quick Start
* Project URL: `http://localhost:5050` (pgAdmin)
* Database URL: `postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable`
```bash
docker-compose up -d  # Start PostgreSQL and PgAdmin
docker-compose down   # Stop when done
```
