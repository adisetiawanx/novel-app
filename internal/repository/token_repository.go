package repository

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/model"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Save(token *model.Token) (*model.Token, error)
	FindByToken(token string) (*model.Token, error)
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

func (repository *tokenRepositoryImpl) FindByToken(token string) (*model.Token, error) {
	var tokenDB model.Token
	err := repository.DB.Where("refresh_token = ?", token).First(&tokenDB).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &tokenDB, nil
}
