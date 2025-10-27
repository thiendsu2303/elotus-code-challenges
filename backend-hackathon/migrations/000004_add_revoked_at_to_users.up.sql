-- Add revoked_at column to users for time-based token revocation
ALTER TABLE IF EXISTS users
    ADD COLUMN IF NOT EXISTS revoked_at TIMESTAMP NULL;