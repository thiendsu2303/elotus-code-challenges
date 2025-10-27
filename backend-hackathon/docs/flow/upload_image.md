# Upload File/Image Flow

- Endpoint: `POST /api/v1/upload`
- Auth: Bearer token required
- Content-Type: `multipart/form-data`
- Form field: `file` (client sends any file)
- Size limit: 8MB (server-enforced)

## Overview

The frontend allows users to upload any file via the `file` field. The backend enforces both file type and size: only `image/*` MIME types are accepted and files must be ≤ 8MB. If validation fails, an appropriate error is returned (e.g., `415` for non-image, `413` for oversize), and the client resets the selection to let the user choose another file.

## Steps

1. Client creates `FormData` and appends `file` with the selected file (no client-side type/size restrictions). If the server returns an error, the client clears the selection and shows a “Choose another file” action.
2. JWT middleware authenticates the request and injects `userID` into context.
3. Handler `UploadImage`:
   - Read `fileHeader := c.FormFile("file")` (fallback to `data` for backward compatibility).
   - Open the file and sniff the first 512 bytes to detect `content-type`.
   - Validate that `content-type` starts with `image/` (reject otherwise with `415`).
   - Only after confirming it's an image, check `size <= 8MB` (reject with `413` if too large).
   - Extract `userAgent` and `clientIP` from the request.
   - Call `imageService.SaveUpload(userID, fileHeader, contentType, userAgent, clientIP)`.
4. Service `SaveUpload`:
   - Ensure directory `backend-hackathon/tmp` exists.
   - Save file into `tmp/img_*` (random name) using `os.CreateTemp`.
   - Persist metadata in DB: `filename`, `content_type`, `size_bytes`, `path` (e.g., `tmp/img_1234`), `user_agent`, `client_ip`.
5. Response:
   - Return `201 Created` with the image object: `id`, `filename`, `content_type`, `size_bytes`, `path`, `uploaded_at`.

## Storage

- Files are stored under `backend-hackathon/tmp/` and this directory is listed in `.gitignore`.
- The `path` field is a repo-relative path like `tmp/img_4018117933`.

## Common Errors

- `400 Bad Request`: missing file or invalid payload.
- `401 Unauthorized`: missing/invalid Bearer token.
- `413 Request Entity Too Large`: exceeds 8MB.
- `415 Unsupported Media Type`: not `image/*`.
- `500 Internal Server Error`: system/file/DB error.

## Notes

- The backend tolerates both `file` and `data` fields during transition; use `file` going forward.
- To make storage configurable, add an env var (e.g., `UPLOAD_DIR`) and read it in config.