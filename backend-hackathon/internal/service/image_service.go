package service

import (
    "backend-hackathon/internal/domain"
    "backend-hackathon/internal/repository"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
)

type ImageService interface {
    ListMyImages(userID uint) ([]domain.Image, error)
    CreateImage(userID uint, filename, contentType string, sizeBytes int64, path, userAgent, clientIP string) (domain.Image, error)
    // SaveUpload saves the uploaded file to tmp and persists metadata
    SaveUpload(userID uint, fileHeader *multipart.FileHeader, contentType string, userAgent, clientIP string) (domain.Image, error)
}

type imageService struct {
    images repository.ImageRepository
}

func NewImageService(images repository.ImageRepository) ImageService {
    return &imageService{images: images}
}

func (s *imageService) ListMyImages(userID uint) ([]domain.Image, error) {
    return s.images.ListByUserID(userID)
}

func (s *imageService) CreateImage(userID uint, filename, contentType string, sizeBytes int64, path, userAgent, clientIP string) (domain.Image, error) {
    img := domain.Image{
        UserID:      &userID,
        Filename:    filename,
        ContentType: contentType,
        SizeBytes:   sizeBytes,
        Path:        path,
        UserAgent:   userAgent,
        ClientIP:    clientIP,
    }
    return s.images.Create(img)
}

func (s *imageService) SaveUpload(userID uint, fileHeader *multipart.FileHeader, contentType string, userAgent, clientIP string) (domain.Image, error) {
    // Open the uploaded file
    src, err := fileHeader.Open()
    if err != nil {
        return domain.Image{}, err
    }
    defer src.Close()

    // Ensure local repo tmp directory exists
    if err := os.MkdirAll("tmp", 0o755); err != nil {
        return domain.Image{}, err
    }

    // Save to temp file under repo-local tmp directory
    dst, err := os.CreateTemp("tmp", "img_*")
    if err != nil {
        return domain.Image{}, err
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return domain.Image{}, err
    }

    // Persist metadata
    relPath := filepath.Join("tmp", filepath.Base(dst.Name()))
    img := domain.Image{
        UserID:      &userID,
        Filename:    fileHeader.Filename,
        ContentType: contentType,
        SizeBytes:   fileHeader.Size,
        Path:        relPath,
        UserAgent:   userAgent,
        ClientIP:    clientIP,
    }
    return s.images.Create(img)
}