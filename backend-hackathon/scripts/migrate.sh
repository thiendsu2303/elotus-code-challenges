#!/usr/bin/env bash
set -euo pipefail

# Simple migration helper wrapper
# Usage:
#   ./scripts/migrate.sh up
#   ./scripts/migrate.sh down
#   ./scripts/migrate.sh drop
#   ./scripts/migrate.sh version
#   ./scripts/migrate.sh force <version>
#   ./scripts/migrate.sh create <name>

MIGRATIONS_DIR=${MIGRATIONS_DIR:-migrations}
DB_URL=${DB_URL:-postgres://postgres:postgres@localhost:5432/hackathon_db?sslmode=disable}

ensure_migrate() {
  if ! command -v migrate >/dev/null 2>&1; then
    echo "Installing golang-migrate CLI..."
    go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    # Ensure GOPATH/bin is on PATH for current session
    export PATH="${GOPATH:-$HOME/go}/bin:$PATH"
  fi
}

cmd=${1:-}
case "$cmd" in
  up)
    ensure_migrate
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" up
    ;;
  down)
    ensure_migrate
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" down
    ;;
  drop)
    ensure_migrate
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" drop
    ;;
  version)
    ensure_migrate
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" version || true
    ;;
  force)
    ensure_migrate
    ver=${2:?"force requires version number"}
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" force "$ver"
    ;;
  create)
    ensure_migrate
    name=${2:?"create requires name, e.g. create_images_table"}
    migrate create -ext sql -dir "$MIGRATIONS_DIR" -seq "$name"
    ;;
  *)
    echo "Usage: $0 {up|down|drop|version|force <v>|create <name>}"
    exit 1
    ;;
esac