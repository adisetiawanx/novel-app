package model

import (
	"github.com/google/uuid"
	"time"
)

type Chapter struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	NovelID   uuid.UUID `gorm:"type:uuid;not null"`
	Slug      string    `gorm:"type:varchar(100);unique;not null"`
	Title     string    `gorm:"type:varchar(150);not null"`
	Number    int       `gorm:"autoIncrement;not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Novel *Novel `gorm:"foreignKey:NovelID"`
}

func (Chapter) TableName() string {
	return "novelku.chapters"
}
