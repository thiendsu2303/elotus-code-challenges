-- Drop tokens table as we now use stateless JWT with per-user revoke timestamp
DROP TABLE IF EXISTS tokens;