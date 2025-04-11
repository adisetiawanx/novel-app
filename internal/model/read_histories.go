package model

import (
	"github.com/google/uuid"
	"time"
)

type ReadHistory struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	ChapterID uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	NovelID   uuid.UUID `gorm:"type:uuid;not null"`
	ReadAt    time.Time `gorm:"autoCreateTime"`

	User    *User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Chapter *Chapter `gorm:"foreignKey:ChapterID;constraint:OnDelete:CASCADE"`
	Novel   *Novel   `gorm:"foreignKey:NovelID;constraint:OnDelete:CASCADE"`
}

func (ReadHistory) TableName() string {
	return "novelku.read_histories"
}
