package repository

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/model/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindByID(ID string) (*entity.User, error)
	IsEmailExist(email string) (bool, error)
	FindByEmail(email string) (*entity.User, error)
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: DB,
	}
}

func (repository *userRepositoryImpl) Save(user *entity.User) (*entity.User, error) {
	result := repository.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindByID(ID string) (*entity.User, error) {
	user := new(entity.User)
	result := repository.DB.Take(user, "id = ?", ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return user, nil
}

func (repository *userRepositoryImpl) IsEmailExist(email string) (bool, error) {
	var count int64
	result := repository.DB.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return count > 0, result.Error
	}

	return count > 0, nil
}

func (repository *userRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	result := repository.DB.Where("email = ?", email).Take(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}
