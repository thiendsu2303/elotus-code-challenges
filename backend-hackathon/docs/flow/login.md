# Login Flow

- Endpoint: `POST /api/v1/auth/login`
- Payload: `{ username, password }`

Steps (idea):
- Handler parses JSON and validates `username`, `password`.
- Service loads user by `username` (repo `GetByUsername`), compare password using `bcrypt.CompareHashAndPassword`.
- Issue JWT (HS256) with claims: `sub` (user id), `iat` (issued at), `exp` (expires at), `iss` (`backend-hackathon`).
 - Issue JWT (HS256) with claims: `sub` (user id), `iat` (issued at), `exp` (expires at), `iss` (from env `JWT_ISSUER`).
- Sign with `JWT_SECRET`, TTL from `ACCESS_TOKEN_TTL`.
- Return `200 OK` with `{ access_token, token_type: "Bearer", expires_at }`.

 Auth & revoke (idea):
 - Middleware: verify JWT signature, check `iss` equals env `JWT_ISSUER`, check `exp`, load user by `sub`, reject if `iat <= users.revoked_at`.
- Time-based revoke: set `users.revoked_at = NOW()` to invalidate all older tokens (no `tokens` table).

Common errors:
- `400 Bad Request`: invalid payload.
- `401 Unauthorized`: wrong credentials.
- `500 Internal Server Error`: system/DB error.