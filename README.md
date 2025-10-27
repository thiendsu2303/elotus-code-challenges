# elotus-code-challenges

Quickstart to run the backend API and the frontend demo locally, then try the upload flow.

## Prerequisites

- Docker Desktop (Compose v2)
- Go 1.23.x (the module uses `go 1.23.0`)
- Node.js >= 18.18 or 20.x (recommended Node 20 LTS) and `npm`
- golang-migrate CLI (to run migrations)

Install golang-migrate CLI:
- macOS: `brew install golang-migrate`
- Or: `go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

Node version management (optional):
- Install `nvm`: `brew install nvm`
- Install Node 20 LTS: `nvm install 20 && nvm use 20`

## Start Backend (API)

1. `cd backend-hackathon`
2. Copy env and adjust if needed: `cp env.example .env`
3. Start database: `make start` (runs Postgres on `localhost:5432`)
4. Run migrations: `make migrate-up`
5. Start API server: `make run`

The API listens on `http://localhost:8080`.

Notes:
- Common Make targets: `make stop`, `make restart`, `make clean`.
- You can also run directly: `go run cmd/api/main.go`.

## Start Frontend (Next.js)

1. Open a new terminal
2. `cd front-end-demo`
3. Install deps: `npm install`
4. Start dev server: `npm run dev`

The app is available at `http://localhost:3000`.

By default, the frontend calls the API at `http://localhost:8080` (`NEXT_PUBLIC_API_BASE_URL` is optional). If you need to point elsewhere, set `NEXT_PUBLIC_API_BASE_URL` in `.env.local` or your shell.

## Swagger

- When the backend is running, open Swagger UI: `http://localhost:8080/swagger/index.html`
- Raw docs: `backend-hackathon/docs/swagger.yaml` and `backend-hackathon/docs/swagger.json`
- Regenerate docs (optional): `cd backend-hackathon && make swagger`

## Try It Out

- Register: visit `http://localhost:3000/register`
- Login: visit `http://localhost:3000/login`
- Protected page: `http://localhost:3000/resource`
- Upload:
  - The frontend lets you choose any file; backend only accepts `image/*` and size ≤ 8MB.
  - On error (non-image or too large), the UI shows a message and a “Choose another file” button, and resets the previous selection.
  - On success, the uploaded image appears in the list.

## Docs

- API Swagger docs are maintained at `backend-hackathon/docs/swagger.yaml` and `backend-hackathon/docs/swagger.json`.
- Upload flow details: `backend-hackathon/docs/flow/upload_image.md`.

## Troubleshooting

- Ensure Docker is running and Postgres is up (`make start`).
- If migrations fail, verify DB connection settings in `.env` match the Makefile `DB_URL`.
- If a port is in use, stop the conflicting process or change the port(s) in env/config.