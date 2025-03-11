package module

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/adisetiawanx/novel-app/internal/router"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUserModule(server *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router.NewUserRouter(server, userController)
}
