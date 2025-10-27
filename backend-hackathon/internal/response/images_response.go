package response

import (
    "backend-hackathon/internal/domain"
    "time"
)

type ImageItem struct {
    // ID is the image identifier
    // example: 10
    ID          uint       `json:"id"`
    // UserID is the owner user id (nullable)
    // example: 1
    UserID      *uint      `json:"user_id"`
    // Filename is the original filename
    // example: holiday.jpg
    Filename    string     `json:"filename"`
    // ContentType is the MIME type
    // example: image/jpeg
    ContentType string     `json:"content_type"`
    // SizeBytes is the file size in bytes
    // example: 204800
    SizeBytes   int64      `json:"size_bytes"`
    // Path is the storage path or URL
    // example: /uploads/1/holiday.jpg
    Path        string     `json:"path"`
    // UploadedAt is the upload time (RFC3339)
    // example: 2025-01-01T12:00:00Z
    UploadedAt  time.Time  `json:"uploaded_at"`
}

type ImagesResponse struct {
    BaseResponse
    Data []ImageItem `json:"data"`
}

func FromDomainImages(images []domain.Image) []ImageItem {
    out := make([]ImageItem, 0, len(images))
    for _, im := range images {
        out = append(out, ImageItem{
            ID:          im.ID,
            UserID:      im.UserID,
            Filename:    im.Filename,
            ContentType: im.ContentType,
            SizeBytes:   im.SizeBytes,
            Path:        im.Path,
            UploadedAt:  im.UploadedAt,
        })
    }
    return out
}

func FromDomainImage(im domain.Image) ImageItem {
    return ImageItem{
        ID:          im.ID,
        UserID:      im.UserID,
        Filename:    im.Filename,
        ContentType: im.ContentType,
        SizeBytes:   im.SizeBytes,
        Path:        im.Path,
        UploadedAt:  im.UploadedAt,
    }
}