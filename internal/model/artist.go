package model

import "github.com/google/uuid"

type Artist struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Slug string    `gorm:"type:varchar(50);unique;not null"`
	Name string    `gorm:"type:varchar(50);unique;not null"`

	Novels []*Novel `gorm:"many2many:novelku.novel_artists;joinForeignKey:ArtistID;joinReferences:NovelID"`
}

func (Artist) TableName() string {
	return "novelku.artists"
}
