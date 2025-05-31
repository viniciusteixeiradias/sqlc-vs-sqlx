# Task Manager Project

This project demonstrates two approaches to using Go with PostgreSQL:
- Using [`sqlc`](./sqlc/README.md)
- Using [`sqlx`](./sqlx/README.md)

---

### 🔗 Quick Start
* Project URL: `http://localhost:5050` (pgAdmin)
* Database URL: `postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable`
```bash
docker-compose up -d  # Start PostgreSQL and PgAdmin
docker-compose down   # Stop when done
```
