package main

import (
	"github.com/adisetiawanx/novel-app/internal/app"
	"github.com/adisetiawanx/novel-app/internal/module"
	"github.com/adisetiawanx/novel-app/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() *echo.Echo {
	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())

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

	module.RegisterAuthModule(apiGroup, tokenRepository, userRepository)

	server.Logger.Fatal(server.Start(":3000"))
}
