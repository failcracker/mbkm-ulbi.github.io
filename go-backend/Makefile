include .env

.PHONY: run build test clean seed docker-up docker-down

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/mbkm-backend main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Seed the database
seed:
	go run scripts/seed.go

# Start services with Docker Compose
docker-up:
	docker-compose up -d

# Stop services
docker-down:
	docker-compose down

# Start services and follow logs
docker-logs:
	docker-compose up

# Rebuild and start services
docker-rebuild:
	docker-compose up --build

# Run database migrations
migrate:
	go run main.go

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate API documentation
docs:
	swag init

# Development setup
dev-setup: deps seed

# Production build
prod-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/mbkm-backend main.go