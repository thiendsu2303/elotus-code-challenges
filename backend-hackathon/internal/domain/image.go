package domain

import "time"

// Image represents an uploaded image metadata
type Image struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      *uint     `gorm:"index" json:"user_id"`
	Filename    string    `gorm:"not null" json:"filename"`
	ContentType string    `gorm:"not null;column:content_type" json:"content_type"`
	SizeBytes   int64     `gorm:"not null;column:size_bytes" json:"size_bytes"`
	Path        string    `gorm:"not null" json:"path"`
	UploadedAt  time.Time `gorm:"autoCreateTime;column:uploaded_at" json:"uploaded_at"`
}

// TableName specifies the table name for Image
func (Image) TableName() string {
	return "images"
}
