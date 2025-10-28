# Backend Hackathon - Go Clean Architecture

## Quick Overview

- Stack: Go (Gin), GORM, PostgreSQL.
- Prerequisites: Go 1.23.x, Docker Desktop (Compose v2), golang-migrate CLI.
- API base: `http://localhost:8080`.
- Swagger UI: `http://localhost:8080/swagger/index.html` (sources: `docs/swagger.yaml`, `docs/swagger.json`; regenerate with `make swagger`).

## Start

- Change to project: `cd backend-hackathon`
- Setup env: `cp env.example .env`
  - Edit `.env` if needed (DB host, port, user, password, JWT settings).
  - Without `.env`, sensible defaults are used; you can override via shell, e.g. `export DB_HOST=localhost`.
- Start database: `make start`
  - Requires Docker Desktop running.
  - Postgres listens on `localhost:5432` (`postgres/postgres`, DB `hackathon_db`).
  - Stop with `make stop`, clean volumes with `make clean`.
- Run migrations: `make migrate-up`
  - Requires `golang-migrate` CLI installed.
  - If it fails, confirm Postgres is up (`docker ps`) and `DB_URL` matches your setup.
- Start API: `make run`
  - Server runs at `http://localhost:8080`.
  - Health check: `curl http://localhost:8080/ping` (should return OK JSON).
  - Swagger UI: open `http://localhost:8080/swagger/index.html`.

## API Endpoints

- `GET /ping` — Health check
- `POST /api/v1/auth/register` — Register user
- `POST /api/v1/auth/login` — Obtain JWT access token
- `POST /api/v1/auth/logout` — Logout (requires Bearer token)
- `GET /api/v1/ping-auth` — Authenticated ping (requires Bearer token)
- `GET /api/v1/resource/images` — List images for the authenticated user
- `POST /api/v1/upload` — Upload image (multipart/form-data; field `file`, legacy fallback `data`; accepts `image/*`, ≤ 8MB)

## Swagger

- View UI: `http://localhost:8080/swagger/index.html`
- Regenerate: `make swagger` (requires `swag` on PATH)
- Sources: `docs/swagger.yaml`, `docs/swagger.json`, generated `docs.go`

## Configuration

Set via `.env` (copy from `env.example`):
- `SERVER_PORT` (default: 8080)
- `DB_HOST` (default: localhost)
- `DB_PORT` (default: 5432)
- `DB_USER` (default: postgres)
- `DB_PASSWORD` (default: postgres)
- `DB_NAME` (default: hackathon_db)
- `DB_SSL_MODE` (default: disable)
- `JWT_SECRET`, `JWT_ISSUER`, `ACCESS_TOKEN_TTL`

## Useful Commands

- DB lifecycle: `make start`, `make stop`, `make restart`, `make clean`
- Migrations: `make migrate-up`, `make migrate-down`, `make migrate-version`, `make migrate-create NAME=...`
- App lifecycle: `make install`, `make build`, `make run`, `make test`