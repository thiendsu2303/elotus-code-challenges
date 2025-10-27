# Upload Image Flow

- Endpoint: `POST /api/v1/upload`
- Auth: Bearer token required
- Content-Type: `multipart/form-data`
- Form field: `file` (image)
- Size limit: 8MB (client and server checks)

## Steps

1. Client selects an image (`image/*`) and submits multipart with field `file`.
2. JWT middleware authenticates the request and injects `userID` into context.
3. Handler `UploadImage`:
   - Read `fileHeader := c.FormFile("file")` (fallback to `data` if missing).
   - Open the file and sniff the first 512 bytes to detect `content-type`.
   - Validate that `content-type` starts with `image/*` (reject otherwise).
   - Only after confirming it's an image, check `size <= 8MB`.
   - Prepare `userAgent` and `clientIP` from the request.
   - Call `imageService.SaveUpload(userID, fileHeader, contentType, userAgent, clientIP)`.
4. Service `SaveUpload`:
   - Ensure directory `backend-hackathon/tmp` exists.
   - Save file into `tmp/img_*` (random name) using `os.CreateTemp`.
   - Persist metadata in DB: `filename`, `content_type`, `size_bytes`, `path` (e.g. `tmp/img_1234`), `user_agent`, `client_ip`.
5. Response:
   - Return `201 Created` with the image object (`id`, `filename`, `content_type`, `size_bytes`, `path`, `uploaded_at`).

## Storage

- Images are stored under `backend-hackathon/tmp/` and this directory is listed in `.gitignore`.
- The `path` field is a repo-relative path like `tmp/img_4018117933`.

## Common errors

- `400 Bad Request`: missing file or invalid payload.
- `401 Unauthorized`: missing/invalid Bearer token.
- `413 Request Entity Too Large`: exceeds 8MB.
- `415 Unsupported Media Type`: not `image/*`.
- `500 Internal Server Error`: system/file/DB error.

## Notes

- If you need a configurable upload directory, extend config to read from an env var (e.g. `UPLOAD_DIR`).