# Expense Tracker Backend (Go + Gin)

A backend template built using **Golang**, **Gin**, **PostgreSQL**, **Redis**, and **JWT authentication**.  
This project focuses on clean code structure, modular design, and practical features that are commonly used in real-world backend services.

---

## 🚀 Features

- User Register & Login
- JWT Access Tokens (1 hour)
- Refresh Tokens (7 days)
- Refresh Token Rotation (old token becomes invalid)
- Redis-backed refresh sessions
- IP + User-Agent checks for refresh tokens
- Password hashing (bcrypt)
- Config-driven setup using YAML
- Auto database migrations (GORM)
- Clean routing with versioning (`/api/v1`)
- Graceful shutdown (server closes connections safely)

---

## 📁 Project Structure
ackend/
│
├── cmd/server/ # Application entrypoint
│
├── internal/
│ ├── app/ # Dependency wiring
│ ├── cache/ # Redis client
│ ├── config/ # Config loader (YAML + env overrides)
│ ├── controller/ # HTTP handlers
│ ├── database/ # Postgres connection
│ ├── domain/ # Models + DTOs
│ ├── repository/ # Database operations
│ ├── router/ # Route groups
│ └── service/ # Business logic
│
├── pkg/
│ └── response/ # Standard API response format
│
├── config.example.yaml
└── README.md


The code is organized in a way that makes it easy to extend the project with new modules (expenses, accounts, categories, etc.).

---

## 🛠 Requirements

- Go 1.22+
- Docker (PostgreSQL + Redis)
- Postman / Thunder Client to test APIs

---

## 🐳 Running Postgres & Redis with Docker

### PostgreSQL

```sh
docker run --name exp-postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=expense_tracker \
  -p 5432:5432 -d postgres
```

Setup Instructions
Install dependencies:
go mod tidy

Copy the example config file:
cp config.example.yaml config.yaml

Start the server (with Air):
air


Server will start on the port defined in config.yaml.
