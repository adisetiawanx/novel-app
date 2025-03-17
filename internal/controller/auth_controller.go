package controller

import (
	"errors"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/adisetiawanx/novel-app/internal/model/web"
	"github.com/adisetiawanx/novel-app/internal/model/web/request"
	"github.com/adisetiawanx/novel-app/internal/model/web/response"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController interface {
	Register(context echo.Context) error
	Login(context echo.Context) error
}

type authControllerImpl struct {
	service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &authControllerImpl{
		AuthService: service,
	}
}

func (controller *authControllerImpl) Register(ctx echo.Context) error {
	var req request.AuthRegisterRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.APIResponse{
			Message: "Invalid request format",
			Data:    nil,
		})
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.APIResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	user, err := controller.AuthService.Register(&req)
	var baseError *helper.BaseError
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, web.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusCreated, web.APIResponse{
		Message: "User Created Successfully",
		Data: web.UserResponseWrapper{
			User: response.AuthRegisterResponse{
				Id: user.ID.String(),
			},
		},
	})
}

func (controller *authControllerImpl) Login(ctx echo.Context) error {
	var req request.AuthLoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.APIResponse{
			Message: "Invalid request format",
			Data:    nil,
		})
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.APIResponse{
			Message: err.Error(),
			Data:    nil,
		})
	}

	user, accessToken, refreshToken, err := controller.AuthService.Login(&req)
	var baseError *helper.BaseError
	if errors.As(err, &baseError) {
		return ctx.JSON(baseError.StatusCode, web.APIResponse{
			Message: baseError.Message,
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, web.APIResponse{
		Message: "User Login Successfully",
		Data: web.LoginResponseWrapper{
			User: response.AuthLoginResponse{
				Id:    user.ID.String(),
				Email: user.Email,
			},
			Token: response.AuthLoginTokenResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		},
	})
}
