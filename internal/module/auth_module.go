package module

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/adisetiawanx/novel-app/internal/routes"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
)

func RegisterAuthModule(server *echo.Echo, tokenRepository repository.TokenRepository, userRepository repository.UserRepository) {
	authService := service.NewAuthService(userRepository, tokenRepository)
	authController := controller.NewAuthController(authService)
	routes.RegisterAuthRoutes(server, authController)
}
