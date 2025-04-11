package service

import (
	"context"
	"encoding/json"
	"github.com/adisetiawanx/novel-app/internal/dto"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/model"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"net/http"
)

type AuthService interface {
	HandleGoogleLogin(code string) (*model.User, string, string, error)
	GenerateGoogleLoginURL(state string) string
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

func (service *authServiceImpl) GenerateGoogleLoginURL(state string) string {
	googleConfig := helper.GetGoogleOAuthConfig()
	return googleConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (service *authServiceImpl) HandleGoogleLogin(code string) (*model.User, string, string, error) {
	googleConfig := helper.GetGoogleOAuthConfig()
	token, err := googleConfig.Exchange(context.TODO(), code)
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	client := googleConfig.Client(context.TODO(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", "", helper.NewInternalServerError()
	}

	var googleUser dto.GoogleUser

	if err = json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	user, err := service.UserRepository.Save(&model.User{
		ID:         newId,
		Name:       googleUser.Name,
		Email:      googleUser.Email,
		Profile:    googleUser.Picture,
		Provider:   "google",
		ProviderID: googleUser.ID,
	})

	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	newAccessToken, err := helper.CreateAccessToken(user.ID.String(), string(user.Role))
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	newRefreshToken, err := helper.CreateRefreshToken(user.ID.String(), string(user.Role))
	if err != nil {
		return nil, "", "", helper.NewInternalServerError()
	}

	return user, newAccessToken, newRefreshToken, nil

}
