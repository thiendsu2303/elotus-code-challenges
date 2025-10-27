package domain

import "time"

type Image struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      *uint     `gorm:"index" json:"user_id"`
    Filename    string    `gorm:"not null" json:"filename"`
    ContentType string    `gorm:"not null;column:content_type" json:"content_type"`
    SizeBytes   int64     `gorm:"not null;column:size_bytes" json:"size_bytes"`
    Path        string    `gorm:"not null" json:"path"`
    UserAgent   string    `gorm:"column:user_agent" json:"user_agent"`
    ClientIP    string    `gorm:"column:client_ip" json:"client_ip"`
    UploadedAt  time.Time `gorm:"autoCreateTime;column:uploaded_at" json:"uploaded_at"`
}

func (Image) TableName() string {
	return "images"
}
