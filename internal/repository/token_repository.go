package repository

import (
	"github.com/adisetiawanx/novel-app/internal/model"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Save(token *model.Token) (*model.Token, error)
}

type tokenRepositoryImpl struct {
	DB *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepositoryImpl{DB: db}
}

func (repository *tokenRepositoryImpl) Save(token *model.Token) (*model.Token, error) {
	result := repository.DB.Create(token)

	if result.Error != nil {
		return nil, result.Error
	}

	return token, nil
}
