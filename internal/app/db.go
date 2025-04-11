package app

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", Config.Database.Host, Config.Database.User, Config.Database.Password, Config.Database.Name, Config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}
