.PHONY: help start stop restart migrate-up migrate-down migrate-create clean build run

# Variables
DB_URL=postgres://postgres:postgres@localhost:5432/hackathon_db?sslmode=disable
MIGRATE_DRIVER=postgres
MIGRATIONS_DIR=migrations

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

start: ## Start the database container
	docker-compose up -d
	@echo "Waiting for database to be ready..."
	@sleep 3
	@echo "Database is ready!"

stop: ## Stop the database container
	docker-compose down
	@echo "Database stopped!"

restart: stop start ## Restart the database container

migrate-up: ## Run all migrations up
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) up

migrate-down: ## Run all migrations down
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down

migrate-drop: ## Drop all database tables
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) drop

migrate-force: ## Force database version (usage: make migrate-force VERSION=1)
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) force $(VERSION)

migrate-version: ## Show current migration version
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) version

migrate-create: ## Create a new migration (usage: make migrate-create NAME=create_users_table)
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=your_migration_name"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

clean: ## Remove database volume and containers
	docker-compose down -v
	@echo "Database cleaned!"

install: ## Install dependencies
	go mod download
	go mod tidy

build: ## Build the application
	go build -o bin/api cmd/api/main.go

run: ## Run the application
	go run cmd/api/main.go

test: ## Run tests
	go test -v ./...

.DEFAULT_GOAL := help

