package model

import (
	"github.com/google/uuid"
	"time"
)

type Rating struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	NovelID   uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Rating    float32   `gorm:"type:numeric(4,1);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User  *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Novel *Novel `gorm:"foreignKey:NovelID;references:ID;constraint:OnDelete:CASCADE"`
}

func (Rating) TableName() string {
	return "novelku.ratings"
}
