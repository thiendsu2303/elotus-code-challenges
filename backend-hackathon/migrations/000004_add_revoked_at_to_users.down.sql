-- Remove revoked_at column from users
ALTER TABLE IF EXISTS users
    DROP COLUMN IF EXISTS revoked_at;