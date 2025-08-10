# Variables
DB_USER=postgres
DB_PASS=postgres
DB_NAME=flowers4you
DB_HOST=localhost
DB_PORT=5432
MIGRATIONS_DIR=./migrations
MIGRATE_CMD=migrate -path $(MIGRATIONS_DIR) -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

# Build the Go app
build:
	go build -o main .

# Run the Go app
run:
	go run main.go

# Start Docker services
up:
	docker-compose up -d

# Stop Docker services
down:
	docker-compose down

# Create a new migration
# Usage: make migration name=create_users_table
migration:
	@migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

# Apply migrations
migrate-up:
	@$(MIGRATE_CMD) up

# Rollback last migration
migrate-down:
	@$(MIGRATE_CMD) down 1

# Check migration version
migrate-version:
	@$(MIGRATE_CMD) version