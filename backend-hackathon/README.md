# Backend Hackathon - Go Clean Architecture

A clean architecture implementation using Gin framework, GORM, and PostgreSQL.

## Project Structure

```
backend-hackathon/
├── cmd/
│   └── api/              # Application entry point
│       └── main.go
├── internal/
│   ├── config/           # Configuration management
│   ├── domain/           # Domain entities
│   ├── handler/          # HTTP handlers (presentation layer)
│   ├── middleware/       # HTTP middlewares
│   ├── repository/       # Data access layer
│   └── usecase/          # Business logic layer
├── migrations/           # Database migrations
├── docker-compose.yml    # Docker compose for PostgreSQL
├── Makefile             # Make commands for project tasks
└── go.mod               # Go module file
```

## Architecture Layers

- **Domain**: Core business entities
- **Repository**: Data access layer (GORM)
- **Usecase**: Business logic layer
- **Handler**: HTTP handlers (Gin)
- **Middleware**: Cross-cutting concerns (auth, logging, etc.)

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Make (optional but recommended)

## Getting Started

### 1. Install Dependencies

```bash
make install
# or
go mod download
```

### 2. Start Database

```bash
make start
# or
docker-compose up -d
```

This will start a PostgreSQL container on port 5432.

### 3. Run Migrations

```bash
make migrate-up
```

Alternatively, you can use the helper script:

```bash
# Show current version
./scripts/migrate.sh version

# Run all migrations up
./scripts/migrate.sh up

# Create a new migration
./scripts/migrate.sh create add_new_table

# Rollback all migrations
./scripts/migrate.sh down
```

### 4. Run the Application

```bash
make run
# or
go run cmd/api/main.go
```

The server will start on port 8080.

## Available Commands

```bash
# Database management
make start              # Start PostgreSQL container
make stop              # Stop PostgreSQL container
make restart           # Restart PostgreSQL container
make clean             # Remove containers and volumes

# Migrations
make migrate-up        # Run all migrations
make migrate-down      # Rollback all migrations
make migrate-version   # Show current migration version
make migrate-create NAME=create_table  # Create new migration

# Application
make install           # Install Go dependencies
make build            # Build the application
make run              # Run the application
make test             # Run tests
```

## API Endpoints

### Health Check
- `GET /ping` - Check server status

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

## Configuration

Environment variables can be set in `.env` file (create from `.env.example`):

- `SERVER_PORT`: Server port (default: 8080)
- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database user (default: postgres)
- `DB_PASSWORD`: Database password (default: postgres)
- `DB_NAME`: Database name (default: hackathon_db)
- `DB_SSL_MODE`: SSL mode (default: disable)

## Database Schema

### Overview

The database consists of three main tables:

![Database Schema](docs/db/db_schema.png)

### Tables

#### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    revoked_at TIMESTAMP NULL
);
```

Tokens table đã được loại bỏ trong thiết kế hiện tại (JWT stateless + revoke theo thời điểm bằng `users.revoked_at`).

#### Images Table
```sql
CREATE TABLE images (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE SET NULL,
    filename VARCHAR(255) NOT NULL,
    content_type VARCHAR(255) NOT NULL,
    size_bytes BIGINT NOT NULL,
    path TEXT NOT NULL,
    uploaded_at TIMESTAMP DEFAULT NOW()
);
```

### Relationships

- **users → images**: One-to-Many (One user can have many images)

## Authentication Middleware

The authentication middleware is initialized but not yet implemented. To implement authentication:

1. Edit `internal/middleware/auth.go`
2. Add your authentication logic (JWT, session, etc.)
3. Apply the middleware to protected routes in `cmd/api/main.go`

## Development

### Adding a New Feature

1. Create domain entity in `internal/domain/`
2. Create repository in `internal/repository/`
3. Create usecase in `internal/usecase/`
4. Create handler in `internal/handler/`
5. Register routes in `cmd/api/main.go`

### Running Tests

```bash
make test
```

## License

MIT
## Migration Helper Script

For convenience, a wrapper script is available at `scripts/migrate.sh`.

Environment overrides:
- `DB_URL` (default: `postgres://postgres:postgres@localhost:5432/hackathon_db?sslmode=disable`)
- `MIGRATIONS_DIR` (default: `migrations`)

Examples:

```bash
./scripts/migrate.sh up
./scripts/migrate.sh create add_http_metadata_to_images
./scripts/migrate.sh down
```

