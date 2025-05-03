# Configuration
MIGRATIONS_DIR=./migrations

## Run the application
run:
	@go run cmd/api/main.go


# Usage: make create-migration NAME=create_users
create-migration:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) $(NAME)
