package controller

import (
	"github.com/adisetiawanx/novel-app/internal/model/web"
	"github.com/adisetiawanx/novel-app/internal/model/web/request"
	"github.com/adisetiawanx/novel-app/internal/model/web/response"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController interface {
	Create(context echo.Context) error
}

type userControllerImpl struct {
	service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userControllerImpl{
		UserService: service,
	}
}

func (controller *userControllerImpl) Create(ctx echo.Context) error {
	var req request.UserCreateRequest

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

	user, err := controller.UserService.Create(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, web.APIResponse{
			Message: "Failed to create user",
			Data:    nil,
		})
	}

	if user == nil {
		return ctx.JSON(http.StatusInternalServerError, web.APIResponse{
			Message: "Failed to create user",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusCreated, web.APIResponse{
		Message: "User Created Successfully",
		Data: web.UserResponseWrapper{
			User: response.UserCreateResponse{
				Id: user.ID,
			},
		},
	})
}
