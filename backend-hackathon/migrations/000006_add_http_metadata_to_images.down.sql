-- Remove HTTP metadata columns from images
ALTER TABLE IF EXISTS images
    DROP COLUMN IF EXISTS user_agent,
    DROP COLUMN IF EXISTS client_ip;