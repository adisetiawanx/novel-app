package model

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	RefreshToken string    `gorm:"type:varchar(255);not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (Token) TableName() string {
	return "auth.tokens"
}
