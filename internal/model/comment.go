package model

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;not null"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	NovelID   uuid.UUID  `gorm:"type:uuid;not null"`
	ChapterID *uuid.UUID `gorm:"type:uuid"`
	ParentID  *uuid.UUID `gorm:"type:uuid"`
	Content   string     `gorm:"type:text;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`

	User    *User    `gorm:"foreignKey:UserID"`
	Novel   *Novel   `gorm:"foreignKey:NovelID"`
	Chapter *Chapter `gorm:"foreignKey:ChapterID"`

	Parent   *Comment   `gorm:"foreignKey:ParentID"`
	Children []*Comment `gorm:"foreignKey:ParentID"`
}

func (Comment) TableName() string {
	return "novelku.comments"
}
