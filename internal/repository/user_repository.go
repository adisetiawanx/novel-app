package repository

import (
	"github.com/adisetiawanx/novel-app/internal/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: DB,
	}
}

func (repository *userRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	result := repository.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
