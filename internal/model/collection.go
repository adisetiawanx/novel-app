package model

import (
	"github.com/google/uuid"
	"time"
)

type Collection struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User  *User             `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Items []*CollectionItem `gorm:"foreignKey:CollectionID"`
}

func (Collection) TableName() string {
	return "novelku.collections"
}

type CollectionItem struct {
	CollectionID uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	NovelID      uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	Novel      *Novel      `gorm:"foreignKey:NovelID;constraint:OnDelete:CASCADE"`
	Collection *Collection `gorm:"foreignKey:CollectionID;constraint:OnDelete:CASCADE"`
}

func (CollectionItem) TableName() string {
	return "novelku.collection_items"
}
