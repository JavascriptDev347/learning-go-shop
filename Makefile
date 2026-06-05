# Makefile for building the application

help:
	@echo "Makefile for building the application"
	@echo "Usage:"
	@echo "  make build       - Compiles the application and outputs the binary to the bin directory"
	@echo "  make run         - Builds the application and runs it"
	@echo "  make dev         - Runs the application in development mode"
	@echo "  make lint        - Runs golangci-lint on the codebase"
	@echo "  make format      - Formats the code using gofmt and goimports"
	@echo "  make migrate-up   - Applies database migrations"
	@echo "  make migrate-down - Rolls back database migrations"

# Build: compiles the application and outputs the binary to the bin directory
# -o: output file name
build:
	go build -o bin/app ./cmd/api

# Run: builds the application and runs it
run:
	go run ./cmd/api

dev:
	go run ./cmd/api

#lint
lint: format
	golangci-lint run ./...

#format
format:
	@gofmt -s -w .
	@goimports -w .

migrate-up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/go_ecommerce_shop?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/go_ecommerce_shop?sslmode=disable" down

docker-up:
	docker compose -f docker/docker-compose.yml up -d
docker-down:
	docker compose -f docker/docker-compose.yml down -d

