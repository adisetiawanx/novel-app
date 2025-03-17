package module

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/adisetiawanx/novel-app/internal/routes"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterAuthModule(apiGroup *echo.Group, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	tokenRepository := repository.NewTokenRepository(db)
	authService := service.NewAuthService(userRepository, tokenRepository)
	authController := controller.NewAuthController(authService)

	routes.RegisterAuthRoutes(apiGroup, authController)
}
