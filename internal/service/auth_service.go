package service

import (
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/model/entity"
	"github.com/adisetiawanx/novel-app/internal/model/web/request"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/google/uuid"
	"time"
)

type AuthService interface {
	Register(request *request.AuthRegisterRequest) (*entity.User, error)
	Login(request *request.AuthLoginRequest) (*entity.User, string, string, error)
}

type authServiceImpl struct {
	repository.UserRepository
	repository.TokenRepository
}

func NewAuthService(userRepository repository.UserRepository, tokenRepository repository.TokenRepository) AuthService {
	return &authServiceImpl{
		UserRepository:  userRepository,
		TokenRepository: tokenRepository,
	}
}

func (service *authServiceImpl) Register(request *request.AuthRegisterRequest) (*entity.User, error) {
	emailExist, err := service.UserRepository.IsEmailExist(request.Email)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	if emailExist {
		return nil, helper.NewConflictError("email already exist")
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	user, err := service.UserRepository.Save(&entity.User{
		ID:       newId,
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: hashedPassword,
		Role:     entity.Visitor,
	})

	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	return user, nil
}

func (service *authServiceImpl) Login(request *request.AuthLoginRequest) (*entity.User, string, string, error) {
	user, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	if user == nil {
		return nil, "", "", helper.NewAuthenticationError("email or password is wrong")
	}

	isPasswordSame := helper.ComparePassword(user.Password, request.Password)
	if !isPasswordSame {
		return nil, "", "", helper.NewAuthenticationError("email or password is wrong")
	}

	newAccessToken, err := helper.CreateAccessToken(user.ID.String(), string(user.Role))
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	newRefreshToken, err := helper.CreateRefreshToken(user.ID.String(), string(user.Role))
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	newTokenId, err := uuid.NewV7()
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	_, err = service.TokenRepository.Save(&entity.Token{
		ID:           newTokenId,
		RefreshToken: newRefreshToken,
		ExpiresAt:    time.Now().Add(time.Hour * 720),
		UserID:       user.ID,
	})
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	return user, newAccessToken, newRefreshToken, nil
}
