package module

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/adisetiawanx/novel-app/internal/routes"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
)

func RegisterNovelModule(apiGroup *echo.Group, novelRepository repository.NovelRepository) {
	novelService := service.NewNovelService(novelRepository)
	novelController := controller.NewNovelController(novelService)

	routes.RegisterNovelRoutes(apiGroup, novelController)
}
