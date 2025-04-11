package controller

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/dto"
	"github.com/adisetiawanx/novel-app/internal/dto/response"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController interface {
	GoogleLogin(context echo.Context) error
	GoogleCallback(context echo.Context) error
}

type authControllerImpl struct {
	service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &authControllerImpl{
		AuthService: service,
	}
}

func (controller *authControllerImpl) GoogleLogin(ctx echo.Context) error {
	state := helper.GenerateRandomState()

	ctx.SetCookie(&http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})

	loginURL := controller.AuthService.GenerateGoogleLoginURL(state)
	return ctx.Redirect(http.StatusTemporaryRedirect, loginURL)
}

func (controller *authControllerImpl) GoogleCallback(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	state := ctx.QueryParam("state")
	
	cookie, err := ctx.Cookie("oauth_state")
	if err != nil || cookie.Value != state {
		return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Message: "Invalid state",
			Data:    nil,
		})
	}

	if code == "" {
		return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Message: "Invalid request format",
			Data:    nil,
		})
	}

	user, accessToken, refreshToken, err := controller.AuthService.HandleGoogleLogin(code)
	var baseError *helper.BaseError
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, dto.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, dto.APIResponse{
		Message: "User Login Successfully",
		Data: dto.LoginResponseWrapper{
			User: response.AuthLoginResponse{
				Id:      user.ID.String(),
				Email:   user.Email,
				Profile: user.Profile,
			},
			Token: response.AuthLoginTokenResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		},
	})
}
