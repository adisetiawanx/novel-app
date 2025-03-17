package main

import (
	"github.com/adisetiawanx/novel-app/internal/app"
	"github.com/adisetiawanx/novel-app/internal/app/config"
	"github.com/adisetiawanx/novel-app/internal/app/module"
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
	config.InitServerConfig()
	db := app.NewDB()

	server := InitServer()
	apiGroup := server.Group("/api")

	module.RegisterAuthModule(apiGroup, db)

	server.Logger.Fatal(server.Start(":3000"))
}
