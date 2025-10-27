-- Add HTTP metadata columns to images
ALTER TABLE IF EXISTS images
    ADD COLUMN IF NOT EXISTS user_agent TEXT,
    ADD COLUMN IF NOT EXISTS client_ip TEXT;