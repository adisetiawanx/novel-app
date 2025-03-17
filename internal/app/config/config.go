package config

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type TokenConfig struct {
	AccessSecret  string
	RefreshSecret string
}

type AppConfig struct {
	Database DatabaseConfig
	Token    TokenConfig
}

var App AppConfig

func InitServerConfig() {
	v := viper.New()
	v.SetConfigFile("internal/app/config/.env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	App.Database = DatabaseConfig{
		User:     v.GetString("DATABASE_USER"),
		Password: v.GetString("DATABASE_PASSWORD"),
		Host:     v.GetString("DATABASE_HOST"),
		Port:     v.GetString("DATABASE_PORT"),
		Name:     v.GetString("DATABASE_NAME"),
	}

	App.Token = TokenConfig{
		AccessSecret:  v.GetString("TOKEN_ACCESS_SECRET"),
		RefreshSecret: v.GetString("TOKEN_REFRESH_SECRET"),
	}
}
