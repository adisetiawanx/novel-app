package routes

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(apiGroup *echo.Group, controller controller.AuthController) {
	group := apiGroup.Group("/auth")
	group.GET("/google/login", controller.GoogleLogin)
	group.GET("/google/callback", controller.GoogleCallback)
}
