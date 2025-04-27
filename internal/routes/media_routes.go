package routes

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/labstack/echo/v4"
)

func RegisterMediaRoutes(apiGroup *echo.Group, controller controller.MediaController) {
	adminGroup := apiGroup.Group("/admin/media")
	//adminGroup.Use(middleware.IsAdminMiddleware())

	adminGroup.POST("", controller.UploadNovelCover)
}
