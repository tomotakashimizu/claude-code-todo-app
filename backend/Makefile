.PHONY: generate build run test clean deps

# Generate code from OpenAPI spec
generate:
	go generate ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Build the application
build: generate
	go build -o bin/server ./cmd/server

# Run the development server
run: generate
	go run ./cmd/server

# Run tests
test: generate
	go test ./...

# Run tests with coverage
test-coverage: generate
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf internal/generated/
	rm -f coverage.out coverage.html

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Full check (generate, format, lint, test)
check: generate fmt lint test

# Development setup
dev-setup: deps generate

# Build for production
build-prod: generate
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server