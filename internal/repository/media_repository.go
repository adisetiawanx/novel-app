package repository

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/model"
	"gorm.io/gorm"
)

type MediaRepository interface {
	Save(media *model.Media) (*model.Media, error)
	FindByUrl(url string) (*model.Media, error)
}

type mediaRepositoryImpl struct {
	DB *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepositoryImpl{DB: db}
}

func (repository *mediaRepositoryImpl) Save(media *model.Media) (*model.Media, error) {
	result := repository.DB.Create(media)

	if result.Error != nil {
		return nil, result.Error
	}

	return media, nil
}

func (repository *mediaRepositoryImpl) FindByUrl(url string) (*model.Media, error) {
	var mediaDB model.Media
	err := repository.DB.Where("url = ?", url).First(&mediaDB).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &mediaDB, nil
}
