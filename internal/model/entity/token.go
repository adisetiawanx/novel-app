package entity

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	RefreshToken string    `gorm:"size:255;unique;not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (u *Token) TableName() string {
	return "tokens"
}
