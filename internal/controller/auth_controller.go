package controller

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/dto"
	"github.com/adisetiawanx/novel-app/internal/dto/request"
	"github.com/adisetiawanx/novel-app/internal/dto/response"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AuthController interface {
	GoogleLogin(context echo.Context) error
	GoogleCallback(context echo.Context) error
	RefreshToken(context echo.Context) error
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
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
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

	ctx.SetCookie(&http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	if code == "" {
		return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Message: "Invalid code",
			Data:    nil,
		})
	}

	var baseError *helper.BaseError
	user, err := controller.AuthService.HandleGoogleLogin(code)
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, dto.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	accessToken, refreshToken, accessTokenExp, refreshTokenExp, err := controller.AuthService.GenerateTokens(user)
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, dto.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	userAgent := ctx.Request().Header.Get("User-Agent")
	isMobile := helper.IsMobileDevice(userAgent)

	if isMobile {
		return ctx.JSON(http.StatusOK, dto.APIResponse{
			Message: "User Login Successfully",
			Data: response.LoginResponseWrapper{
				Token: response.AuthLoginTokenResponse{
					AccessToken:  accessToken,
					RefreshToken: refreshToken,
				},
			},
		})
	} else {
		ctx.SetCookie(&http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			MaxAge:   int(refreshTokenExp - time.Now().Unix()),
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})

		ctx.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    accessToken,
			Path:     "/",
			MaxAge:   int(accessTokenExp - time.Now().Unix()),
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})

		callbackUrl, err := ctx.Cookie("callback_login_url")
		if err != nil {
			return err
		}

		return ctx.Redirect(http.StatusPermanentRedirect, callbackUrl.Value)
	}
}

func (controller *authControllerImpl) RefreshToken(ctx echo.Context) error {
	var token string

	userAgent := ctx.Request().Header.Get("User-Agent")
	isMobile := helper.IsMobileDevice(userAgent)

	if isMobile {
		var body request.RefreshTokenRequest
		if err := ctx.Bind(&body); err != nil {
			return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
				Message: "Invalid request format",
				Data:    nil,
			})
		}
		token = body.RefreshToken
	} else {
		cookie, err := ctx.Cookie("refresh_token")
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, dto.APIResponse{
				Message: "Missing refresh token",
				Data:    nil,
			})
		}
		token = cookie.Value
	}

	var baseError *helper.BaseError
	newAccessToken, newAccessTokenExp, err := controller.AuthService.RefreshAccessToken(token)
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, dto.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	if isMobile {
		return ctx.JSON(http.StatusOK, dto.APIResponse{
			Message: "Access token refreshed",
			Data:    response.AuthRefreshTokenResponse{AccessToken: newAccessToken},
		})
	} else {
		ctx.SetCookie(&http.Cookie{
			Name:     "token",
			Value:    newAccessToken,
			Path:     "/",
			MaxAge:   int(newAccessTokenExp - time.Now().Unix()),
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})

		return ctx.JSON(http.StatusOK, dto.APIResponse{
			Message: "Access token refreshed",
			Data:    nil,
		})
	}
}
