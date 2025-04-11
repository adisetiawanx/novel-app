package model

import (
	"github.com/google/uuid"
	"time"
)

type Novel struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	PostStatus       string    `gorm:"type:varchar(20);not null;default:publish"`
	Title            string    `gorm:"type:varchar(150);not null"`
	Slug             string    `gorm:"type:varchar(100);not null"`
	AlternativeTitle string    `gorm:"type:varchar(150)"`
	Synopsis         string    `gorm:"type:text"`
	Status           string    `gorm:"type:varchar(20);not null;default:unknown"`
	ReleaseYear      int16     `gorm:"type:smallint;not null"`
	Country          string    `gorm:"type:varchar(20);not null;default:unknown"`
	RatingTotal      float32   `gorm:"type:numeric(4,1);not null"`
	ChapterTotal     int       `gorm:"default:0;not null"`
	CommentTotal     int       `gorm:"default:0;not null"`
	VoteTotal        int       `gorm:"default:0;not null"`
	BookmarkTotal    int       `gorm:"default:0;not null"`
	ViewTotal        int       `gorm:"default:0;not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`

	Genres   []*Genre   `gorm:"many2many:novelku.novel_genres;joinForeignKey:NovelID;joinReferences:GenreID"`
	Authors  []*Author  `gorm:"many2many:novelku.novel_authors;joinForeignKey:NovelID;joinReferences:AuthorID"`
	Artists  []*Artist  `gorm:"many2many:novelku.novel_artists;joinForeignKey:NovelID;joinReferences:ArtistID"`
	Chapters []*Chapter `gorm:"foreignKey:NovelID"`
}

func (Novel) TableName() string {
	return "novelku.novels"
}
