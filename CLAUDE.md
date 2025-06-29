# Todo App

## Tech Stack

- **Backend:** Go 1.24, sqlc for SQL generation, ogen for OpenAPI
- **Frontend:** TypeScript, React Router v7, @tanstack/react-query
- **Database:** SQLite
- **API:** OpenAPI 3.0 YAML specification
- **Testing:** Vitest (unit), Playwright (e2e)
- **Package Manager:** npm/go mod

## Project Structure

- `schema/` - OpenAPI 3.0 YAML specifications
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
- `go generate ./...` - Generate Go server from OpenAPI spec
- `npm run generate:api` - Generate TypeScript API client
- `npm run validate:api` - Validate OpenAPI specification

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

## API Design

- **スキーマファースト:** OpenAPI 3.0 YAML 仕様で API 定義
- **自動生成:** スキーマからコード生成で開発効率向上
- **型安全性:** フロントエンド・バックエンド間の型整合性を保証

## Code Generation

- **Backend:** ogen で OpenAPI から Go サーバーコード生成
- **Frontend:** @tanstack/react-query + openapi-typescript でクライアントコード生成
- **Validation:** OpenAPI スキーマでリクエスト・レスポンス検証

## Development Workflow

1. `schema/openapi.yaml` で API 仕様を定義
2. `go generate ./...` でバックエンドコード生成
3. `npm run generate:api` でフロントエンドクライアント生成
4. 生成されたコードをベースに実装開始

## Workflow

- Run tests before committing changes
- Use conventional commit messages
- Keep database schema changes in separate commits
- Run `sqlc generate` after SQL query changes
