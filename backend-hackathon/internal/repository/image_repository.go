package repository

import (
    "backend-hackathon/internal/domain"

    "gorm.io/gorm"
)

type ImageRepository interface {
    ListByUserID(userID uint) ([]domain.Image, error)
    Create(image domain.Image) (domain.Image, error)
}

type imageRepository struct {
    db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
    return &imageRepository{db: db}
}

func (r *imageRepository) ListByUserID(userID uint) ([]domain.Image, error) {
    var images []domain.Image
    if err := r.db.Where("user_id = ?", userID).Order("uploaded_at DESC").Find(&images).Error; err != nil {
        return nil, err
    }
    return images, nil
}

func (r *imageRepository) Create(image domain.Image) (domain.Image, error) {
    if err := r.db.Create(&image).Error; err != nil {
        return domain.Image{}, err
    }
    return image, nil
}