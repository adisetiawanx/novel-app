package routes

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(server *echo.Echo, controller controller.AuthController) {
	group := server.Group("/auth")
	group.GET("/google/login", controller.GoogleLogin)
	group.GET("/google/callback", controller.GoogleCallback)
	group.POST("/refresh", controller.RefreshToken)
}
