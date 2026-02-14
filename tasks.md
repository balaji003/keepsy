# Completed Tasks

## Initial Setup & Fixes
- [x] Fix `github.com/go-sql-driver/mysql` import error by running `go get`.
- [x] Clean up `go.mod` to remove unused PostgreSQL dependencies (`pgx`).

## Authentication Implementation
- [x] Implement Simple Authentication (Email/Phone + Password).
- [x] Create database migration `000002_create_credentials_table.up.sql` for `user_credentials`.
- [x] Implement `internal/auth` package (Handler, Repository).
- [x] Wire up Auth routes in `main.go` (`POST /auth/register`, `POST /auth/login`).
- [x] Verify authentication endpoints using `curl`.

## API Documentation & Refactoring
- [x] Create OpenAPI Specification at `backend/api/openapi.yaml`.
- [x] Refactor backend to use Service Layer Architecture:
  - [x] Refactor `internal/auth` (Handler -> Service -> Repository).
  - [x] Refactor `internal/users` (Handler -> Service -> Repository).
  - [x] Refactor `internal/products` (Handler -> Service -> Repository).
  - [x] Refactor `internal/categories` (Handler -> Service -> Repository).

## Testing
- [x] Write Unit Tests for all Services using `testify/mock`.
- [x] Verify all tests pass (`go test ./internal/...`).

## Repository Organization & Deployment
- [x] Reorganize project structure into Monorepo (moved backend to `backend/` folder).
- [x] Initialize Git repository at root.
- [x] Create root `.gitignore`.
- [x] Push code to GitHub: `https://github.com/balaji003/keepsy.git`.
