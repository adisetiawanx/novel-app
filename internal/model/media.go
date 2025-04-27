package model

import (
	"github.com/google/uuid"
	"time"
)

type Media struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Url       string    `gorm:"type:varchar(255);not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Size      int
	Type      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Media) TableName() string {
	return "novelku.media"
}
