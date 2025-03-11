package routes

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(server *echo.Echo, controller controller.UserController) {
	group := server.Group("/api")
	group.POST("/user", controller.Create)
}
