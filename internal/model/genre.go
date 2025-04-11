package model

import "github.com/google/uuid"

type Genre struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Slug string    `gorm:"type:varchar(50)"`
	Name string    `gorm:"type:varchar(50)"`

	Novels []*Novel `gorm:"many2many:novelku.novel_genres;"`
}

func (Genre) TableName() string {
	return "novelku.genres"
}
