# Todo Backend

A Go backend implementing DDD (Domain Driven Design) and Clean Architecture principles with automatic OpenAPI code generation using ogen.

## Architecture

### Domain Layer (Core Business Logic)
- `internal/domain/todo/entity.go` - Business entities
- `internal/domain/todo/repository.go` - Repository interfaces
- `internal/domain/todo/service.go` - Domain services

### UseCase Layer (Application Logic)
- `internal/usecase/todo/interfaces.go` - UseCase interfaces and DTOs
- `internal/usecase/todo/interactor.go` - UseCase implementations

### Infrastructure Layer (External Dependencies)
- `internal/infrastructure/database/sqlite/` - SQLite repository implementations
- `internal/interfaces/handler/` - HTTP handlers
- `internal/interfaces/server.go` - HTTP server setup

### Entry Point
- `cmd/server/main.go` - Application entry point with dependency injection

## Features

- **Clean Architecture**: Clear separation of concerns with dependency inversion
- **Domain Driven Design**: Rich domain models with business logic encapsulation
- **OpenAPI Code Generation**: Automatic server code generation from OpenAPI specification
- **SQLite Database**: Lightweight database with type-safe queries
- **Dependency Injection**: Manual DI setup following clean architecture principles

## Getting Started

### Prerequisites
- Go 1.24+
- SQLite3

### Installation

```bash
# Install dependencies
make deps

# Generate code from OpenAPI spec
make generate

# Build the application
make build

# Run the development server
make run
```

### Available Commands

```bash
make deps         # Install dependencies
make generate     # Generate code from OpenAPI spec
make build        # Build the application
make run          # Run development server
make test         # Run tests
make fmt          # Format code
make lint         # Run linter
make clean        # Clean build artifacts
```

## Project Structure

```
backend/
├── cmd/
│   └── server/
│       └── main.go                    # Application entry point
├── internal/
│   ├── domain/
│   │   └── todo/
│   │       ├── entity.go              # Todo entity
│   │       ├── repository.go          # Repository interface
│   │       └── service.go             # Domain service
│   ├── usecase/
│   │   └── todo/
│   │       ├── interfaces.go          # UseCase interfaces
│   │       └── interactor.go          # UseCase implementation
│   ├── infrastructure/
│   │   └── database/
│   │       └── sqlite/
│   │           └── todo_repository.go # SQLite repository
│   ├── interfaces/
│   │   ├── handler/
│   │   │   └── todo_handler.go    # HTTP handlers
│   │   └── server.go              # Server setup
│   └── generated/                     # ogen generated code
├── go.mod
├── go.sum
├── tools.go                           # Tool dependencies
├── generate.go                        # Code generation directive
├── ogen.yaml                          # ogen configuration 
└── Makefile                           # Build tasks
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `DB_PATH` | SQLite database path | `./todo.db` |

## API Endpoints

The server implements the OpenAPI specification defined in `../schema/openapi.yaml`:

- `GET /api/todos` - Get all todos
- `POST /api/todos` - Create a new todo
- `GET /api/todos/{id}` - Get a specific todo
- `PUT /api/todos/{id}` - Update a todo
- `DELETE /api/todos/{id}` - Delete a todo