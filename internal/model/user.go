package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Name       string    `gorm:"type:varchar(150);not null"`
	Email      string    `gorm:"type:varchar(255);not null"`
	Password   string    `gorm:"type:varchar(150)"`
	Profile    string    `gorm:"type:varchar(150)"`
	Role       string    `gorm:"type:varchar(20);default:visitor;not null"`
	Provider   string    `gorm:"type:varchar(50);not null"`
	ProviderID string    `gorm:"type:varchar(100)"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`

	Collections   []*Collection  `gorm:"foreignKey:UserID"`
	ReadHistories []*ReadHistory `gorm:"foreignKey:UserID"`
	Votes         []*Rating      `gorm:"foreignKey:UserID"`
	Comments      []*Comment     `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "auth.users"
}
