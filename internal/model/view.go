package model

import (
	"github.com/google/uuid"
	"net"
	"time"
)

type View struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	NovelID   uuid.UUID `gorm:"type:uuid;not null"`
	IPAddress net.IP    `gorm:"type:inet;not null"`
	ViewedAt  time.Time `gorm:"autoCreateTime"`

	Novel *Novel `gorm:"foreignKey:NovelID"`
}

func (View) TableName() string {
	return "novelku.views"
}
