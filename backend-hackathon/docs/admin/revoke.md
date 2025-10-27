# Admin CLI: Token Revocation

This document explains how to revoke JWTs using the server-side admin CLI. Revocation works by setting `users.revoked_at` to a timestamp. Any token with `iat` <= `revoked_at` is rejected by the auth middleware.

## Prerequisites
- Ensure the app can connect to the database. The CLI loads `.env` via `internal/config`.
- Required env vars: `DATABASE_URL` (and any others your config needs).

## Commands
- `make revoke-all`: Revoke tokens for all users.
- `make revoke-users USERS="alice,bob"`: Revoke tokens for a specific list of usernames.

## Usage Details

### Revoke all users
- Command: `make revoke-all`
- Effect: Sets `users.revoked_at` to current UTC time for all rows.
- Output: Prints rows affected and timestamp.

### Revoke specific users
- Command: `make revoke-users USERS="alice,bob"`
- Effect: Sets `revoked_at` for each provided username.
- Output: Prints success and failed counts.
- Notes: Usernames are comma-separated; whitespace is ignored.

## Direct Go run (optional)
If you prefer not to use Makefile:
- Revoke all: `go run cmd/admin/main.go --all`
- Revoke list: `go run cmd/admin/main.go --users alice,bob`

## How revocation is enforced
- Middleware compares token `iat` with `users.revoked_at`.
- If `revoked_at` is set and `iat` <= `revoked_at`, request is rejected with 401.

## Troubleshooting
- "User not found": Confirm usernames exist in DB.
- DB connection errors: Verify `DATABASE_URL` and network access.
- No rows affected for revoke-all: Ensure the users table has data.