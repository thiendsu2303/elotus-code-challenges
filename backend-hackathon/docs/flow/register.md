# Register Flow

- Endpoint: `POST /api/v1/auth/register`
- Payload: `{ username, password }`

Steps (idea):
- Handler parses JSON and validates `username`, `password`.
- Service checks username uniqueness (repo `GetByUsername`).
- Hash password with `bcrypt` and persist user (repo `Create`).
- Return `201 Created` with user info (`id`, `username`, `created_at`).

Common errors:
- `400 Bad Request`: invalid payload.
- `409 Conflict`: username already exists.
- `500 Internal Server Error`: system/DB error.

Security notes:
- Do not return `password_hash`.
- Use `bcrypt.DefaultCost` for hashing.