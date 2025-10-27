package response

import (
    "backend-hackathon/internal/domain"
    "time"
)

type ImageItem struct {
    ID          uint       `json:"id"`
    UserID      *uint      `json:"user_id"`
    Filename    string     `json:"filename"`
    ContentType string     `json:"content_type"`
    SizeBytes   int64      `json:"size_bytes"`
    Path        string     `json:"path"`
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