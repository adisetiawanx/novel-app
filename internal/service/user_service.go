package service

import (
	"github.com/adisetiawanx/novel-app/internal/model/domain"
	"github.com/adisetiawanx/novel-app/internal/model/web/request"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"gorm.io/gorm"
)

type UserService interface {
	Create(request *request.UserCreateRequest) (*domain.User, error)
}

type userServiceImpl struct {
	repository.UserRepository
	*gorm.DB
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *userServiceImpl) Create(request *request.UserCreateRequest) (*domain.User, error) {
	user, err := service.UserRepository.Create(&domain.User{
		Name:  request.Name,
		Email: request.Email,
		Phone: request.Phone,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
