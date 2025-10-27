# Flow: Token Revocation (Concept & Operation)

This document explains the idea and mechanism of token revocation in the system.

## Objectives
- Invalidate all JWTs issued before a given point in time without maintaining a `jti` blacklist.
- Support bulk revocation or selective revocation per user efficiently.
- Keep checks lightweight in middleware, avoiding heavy DB work.

## How It Works
- Add the `users.revoked_at` field (UTC). When set, any token with `iat` (Issued At) less than or equal to `revoked_at` is rejected.
- The middleware compares `token_iat <= users.revoked_at`. If true, it returns 401.
- Tokens issued after `revoked_at` remain valid.

## Why Time-Based Revocation
- Simple, O(1) check, no blacklist table.
- Fits bulk revocation (all users) or per-user revocation.
- Preserves performance: each request loads one user row and compares timestamps.

## Typical Flow
- Login: Issue token with `iat = now`, `revoked_at = NULL` → valid.
- Admin revoke all: set `revoked_at = now` for all users → any token issued before `now` becomes invalid.
- Admin revoke a list: set `revoked_at = now` for selected users → those users need to log in again.
- Logout: Not required to delete tokens server-side; rely on `exp` (TTL) and revocation when needed.

## Middleware Checks
- Validate `Authorization: Bearer <token>` and `HS256` signature with `JWT_SECRET`.
- Verify `iss == JWT_ISSUER` (from environment), ensure `exp` not expired.
- Load user from DB and compare `iat` against `revoked_at`.
- If `iat <= revoked_at`, return 401; otherwise attach user to context and allow the request.

## Edge Cases
- Time synchronization: Use `UTC` for both `revoked_at` and `iat`. Ensure NTP sync to avoid clock drift.
- Race conditions: If revocation happens while a request is being processed, that request may pass; subsequent requests will be blocked.
- Un-revoke: You can clear/reset `revoked_at` manually, but the CLI does not support it by default. Consider security implications.

## Scalability
- Add `jti` blacklist if you need to revoke individual tokens instead of time-based cutoff.
- Use event-based or message queue propagation for revocation in distributed systems.
- Rotate `JWT_SECRET` periodically to harden security.

## Quick Test
- Obtain a token (login) → call a protected endpoint like `GET /api/v1/ping-auth` → 200.
- Run revoke-all: `make revoke-all` → call again with old token → 401.
- Log in again to get a new token → call again → 200.

## Related Configuration
- `JWT_SECRET`: secret used to sign tokens.
- `JWT_ISSUER`: `iss` claim value verified by middleware.
- `ACCESS_TOKEN_TTL`: token lifetime (used for the `exp` claim).