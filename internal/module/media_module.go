package module

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/adisetiawanx/novel-app/internal/routes"
	"github.com/adisetiawanx/novel-app/internal/service"
	"github.com/labstack/echo/v4"
)

func RegisterMediaModule(apiGroup *echo.Group, mediaRepository repository.MediaRepository) {
	mediaService := service.NewMediaService(mediaRepository)
	mediaController := controller.NewMediaController(mediaService)
	routes.RegisterMediaRoutes(apiGroup, mediaController)
}
