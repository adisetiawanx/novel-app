package routes

import (
	"github.com/adisetiawanx/novel-app/internal/controller"
	"github.com/labstack/echo/v4"
)

func RegisterNovelRoutes(apiGroup *echo.Group, controller controller.NovelController) {
	adminGroup := apiGroup.Group("/admin/novel")
	//adminGroup.Use(middleware.IsAdminMiddleware())

	adminGroup.POST("", controller.Create)
	adminGroup.GET("", controller.FindAll)
}
