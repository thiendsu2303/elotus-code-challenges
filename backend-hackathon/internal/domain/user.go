package domain

import "time"

type User struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Username     string     `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string     `gorm:"column:password_hash;not null" json:"-"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	RevokedAt    *time.Time `gorm:"column:revoked_at" json:"revoked_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}
