package service

import (
	"context"
	"encoding/json"
	"github.com/adisetiawanx/novel-app/internal/dto/request"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/model"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type AuthService interface {
	GenerateGoogleLoginURL(state string) string
	HandleGoogleLogin(code string) (*model.User, error)
	GenerateTokens(user *model.User) (string, string, int64, int64, error)
	RefreshAccessToken(refreshToken string) (string, int64, error)
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

func (service *authServiceImpl) HandleGoogleLogin(code string) (*model.User, error) {
	googleConfig := helper.GetGoogleOAuthConfig()
	token, err := googleConfig.Exchange(context.TODO(), code)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	client := googleConfig.Client(context.TODO(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		return nil, helper.NewInternalServerError()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, helper.NewInternalServerError()
	}

	var googleUser request.GoogleUser

	if err = json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, helper.NewInternalServerError()
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	user, err := service.UserRepository.FindByEmail(googleUser.Email)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	if user != nil {
		return user, nil
	}

	user, err = service.UserRepository.Save(&model.User{
		ID:         newId,
		Name:       googleUser.Name,
		Email:      googleUser.Email,
		Profile:    googleUser.Picture,
		Provider:   "google",
		ProviderID: googleUser.ID,
	})

	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	return user, nil

}

func (service *authServiceImpl) GenerateTokens(user *model.User) (string, string, int64, int64, error) {
	newAccessToken, newAccessTokenExp, err := helper.CreateAccessToken(user.ID.String(), string(user.Role))
	if err != nil {
		return "", "", 0, 0, helper.NewInternalServerError()
	}

	newRefreshToken, newRefreshTokenExp, err := helper.CreateRefreshToken(user.ID.String(), string(user.Role))
	if err != nil {
		return "", "", 0, 0, helper.NewInternalServerError()
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return "", "", 0, 0, helper.NewInternalServerError()
	}

	_, err = service.TokenRepository.Save(&model.Token{
		ID:           newId,
		RefreshToken: newRefreshToken,
		ExpiresAt:    time.Unix(newRefreshTokenExp, 0),
		UserID:       user.ID,
	})

	if err != nil {
		return "", "", 0, 0, helper.NewInternalServerError()
	}

	return newAccessToken, newRefreshToken, newAccessTokenExp, newRefreshTokenExp, nil
}

func (service *authServiceImpl) RefreshAccessToken(refreshToken string) (string, int64, error) {
	claims, err := helper.VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", 0, helper.NewAuthenticationError("refresh token is invalid")
	}

	tokenDB, err := service.TokenRepository.FindByToken(refreshToken)
	if err != nil || tokenDB == nil || tokenDB.ExpiresAt.Before(time.Now()) {
		return "", 0, helper.NewAuthenticationError("refresh token is invalid")
	}

	newAccessToken, newAccessTokenExp, err := helper.CreateAccessToken(claims.UserID, claims.Role)
	if err != nil {
		return "", 0, helper.NewInternalServerError()
	}

	return newAccessToken, newAccessTokenExp, nil
}
