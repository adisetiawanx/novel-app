package repository

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *model.User) (*model.User, error)
	FindByID(ID string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: DB,
	}
}

func (repository *userRepositoryImpl) Save(user *model.User) (*model.User, error) {
	result := repository.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindByID(ID string) (*model.User, error) {
	user := new(model.User)
	result := repository.DB.Take(user, "id = ?", ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)
	result := repository.DB.Where("email = ?", email).Take(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}
