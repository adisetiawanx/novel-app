package app

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
	Database    DatabaseConfig
	Token       TokenConfig
	GoogleOauth GoogleOauthConfig
}

type GoogleOauthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

var Config AppConfig

func InitServerConfig() {
	v := viper.New()
	v.SetConfigFile("configs/.env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	Config.Database = DatabaseConfig{
		User:     v.GetString("DATABASE_USER"),
		Password: v.GetString("DATABASE_PASSWORD"),
		Host:     v.GetString("DATABASE_HOST"),
		Port:     v.GetString("DATABASE_PORT"),
		Name:     v.GetString("DATABASE_NAME"),
	}

	Config.Token = TokenConfig{
		AccessSecret:  v.GetString("TOKEN_ACCESS_SECRET"),
		RefreshSecret: v.GetString("TOKEN_REFRESH_SECRET"),
	}

	Config.GoogleOauth = GoogleOauthConfig{
		ClientID:     v.GetString("GOOGLE_CLIENT_ID"),
		ClientSecret: v.GetString("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  v.GetString("GOOGLE_REDIRECT_URL"),
	}
}
