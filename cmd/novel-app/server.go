package main

import (
	"github.com/adisetiawanx/novel-app/internal/app"
	"github.com/adisetiawanx/novel-app/internal/module"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func InitServer() *echo.Echo {
	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	server.Validator = app.NewCustomValidator()

	return server
}

func Start() {
	app.InitServerConfig()
	db := app.NewDB()

	server := InitServer()
	apiGroup := server.Group("/api")

	userRepository := repository.NewUserRepository(db)
	tokenRepository := repository.NewTokenRepository(db)
	novelRepository := repository.NewNovelRepository(db)
	mediaRepository := repository.NewMediaRepository(db)

	module.RegisterAuthModule(server, tokenRepository, userRepository)
	module.RegisterNovelModule(apiGroup, novelRepository)
	module.RegisterMediaModule(apiGroup, mediaRepository)

	server.Logger.Fatal(server.Start(":3000"))
}
