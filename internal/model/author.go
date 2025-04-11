package model

import "github.com/google/uuid"

type Author struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Slug string    `gorm:"type:varchar(50);unique;not null"`
	Name string    `gorm:"type:varchar(50);unique;not null"`

	Novels []*Novel `gorm:"many2many:novelku.novel_authors;joinForeignKey:AuthorID;joinReferences:NovelID"`
}

func (Author) TableName() string {
	return "novelku.authors"
}
