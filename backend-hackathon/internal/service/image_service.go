package service

import (
    "backend-hackathon/internal/domain"
    "backend-hackathon/internal/repository"
)

type ImageService interface {
    ListMyImages(userID uint) ([]domain.Image, error)
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