package router

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/labstack/echo/v4"
)

func NewUserRouter(server *echo.Echo, controller controller.UserController) {
	group := server.Group("/api")
	group.POST("/user", controller.Create)
}
