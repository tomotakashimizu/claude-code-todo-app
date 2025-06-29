# Todo App

## Tech Stack

- **Backend:** Go 1.24, sqlc for SQL generation
- **Frontend:** TypeScript, React Router v7
- **Database:** SQLite
- **Testing:** Vitest (unit), Playwright (e2e)
- **Package Manager:** npm/go mod

## Project Structure

- `backend/` - Go backend with sqlc-generated database layer
- `frontend/` - React TypeScript application
- `database/` - SQLite schema and migrations
- `tests/` - Unit and e2e test suites

## Commands

- `go run ./backend` - Start Go backend server
- `npm run dev` - Start frontend development server
- `npm run build` - Build frontend for production
- `npm run test` - Run Vitest unit tests
- `npm run test:e2e` - Run Playwright e2e tests
- `go test ./...` - Run Go backend tests
- `sqlc generate` - Generate Go code from SQL queries

## Code Style

- **Go:** Use standard Go formatting with `gofmt`
- **TypeScript:** Use strict TypeScript configuration
- **React:** Prefer function components with hooks
- **Database:** Use prepared statements via sqlc
- **Imports:** Use absolute imports with path aliases

## Database

- SQLite for simplicity and portability
- Use sqlc for type-safe database operations
- Store migrations in `database/migrations/`
- Schema definitions in `database/schema.sql`

## Testing

- Unit tests with Vitest for React components
- Go standard testing for backend logic
- Playwright for full e2e user workflows
- Test database operations with in-memory SQLite

## Workflow

- Run tests before committing changes
- Use conventional commit messages
- Keep database schema changes in separate commits
- Run `sqlc generate` after SQL query changes
