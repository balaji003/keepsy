# Completed Tasks

## Initial Setup & Fixes (2026-02-09)
- [x] Fix `github.com/go-sql-driver/mysql` import error by running `go get`.
- [x] Clean up `go.mod` to remove unused PostgreSQL dependencies (`pgx`).

## Authentication Implementation (2026-02-09)
- [x] Implement Simple Authentication (Email/Phone + Password).
- [x] Create database migration `000002_create_credentials_table.up.sql` for `user_credentials`.
- [x] Implement `internal/auth` package (Handler, Repository).
- [x] Wire up Auth routes in `main.go` (`POST /auth/register`, `POST /auth/login`).
- [x] Verify authentication endpoints using `curl`.

## API Documentation & Refactoring (2026-02-09)
- [x] Create OpenAPI Specification at `backend/api/openapi.yaml`.
- [x] Refactor backend to use Service Layer Architecture:
  - [x] Refactor `internal/auth` (Handler -> Service -> Repository).
  - [x] Refactor `internal/users` (Handler -> Service -> Repository).
  - [x] Refactor `internal/products` (Handler -> Service -> Repository).
  - [x] Refactor `internal/categories` (Handler -> Service -> Repository).

## Testing (2026-02-09)
- [x] Write Unit Tests for all Services using `testify/mock`.
- [x] Verify all tests pass (`go test ./internal/...`).

## Document Upload & Bills Service (2026-02-15)
- [x] Design Implementation Plan for Dynamic Document Upload.
- [x] Implement `internal/storage` package:
  - [x] Define `Service` interface for dynamic storage (Local/S3).
  - [x] Implement `LocalFileSystem` storage.
- [x] Implement `internal/bills` package:
  - [x] Create `bills` database migration.
  - [x] Implement Model, Repository, and Service.
  - [x] Implement HTTP Handler for `POST /bills/upload` and `GET /bills`.
- [x] Wire up Storage and Bills services in `main.go`.
- [x] Add Unit Tests for `bills` Service.

## Download Endpoint (2026-02-15)
- [x] Refactor `GetDownloadURL` into `storage.Service` interface for future proofing (presigned URLs).
- [x] Implement `internal/bills` download logic:
  - [x] Add `GetBillDownloadURL` to Bills Service with ownership check.
  - [x] Implement `GET /bills/download` endpoint with security checks.
- [x] Verify with Unit Tests covering Success, Unauthorized, and Not Found cases.

## Authentication Improvements (2026-02-15)
- [x] Update `Login` service to return `ErrUserNotFound` when user is not found.
- [x] Update `Login` handler to return 404 with specific error code `user_not_found`.
- [x] Add Unit Test for `UserNotFound` scenario.

## API Cleanup (2026-02-15)
- [x] Remove redundant `POST /users` endpoint (since `POST /auth/register` covers it).
- [x] Remove unused `CreateUser` service method and handler.
- [x] Remove proper `users` Service and Handler (Option A). `users` package is now just Repository & Models.

## Service Refactoring (2026-02-15)
- [x] Create `internal/services` directory.
- [x] Move `internal/auth` to `internal/services/auth`.
- [x] Move `internal/storage` to `internal/services/storage`.
- [x] Update imports in `main.go` and `bills/service.go`.

## User UUID Implementation (2026-02-15)
- [x] Create migration `000004_add_uuid_to_users.up.sql`.
- [x] Update `User` model and `Repository` to include `UUID`.
- [x] Update `Auth` service to generate UUID v5 (Name + Phone) on registration.
- [x] Update `Bills` service to store files in `<uuid>/bills/<filename>`.
- [x] Update tests for `Auth` and `Bills` services.
