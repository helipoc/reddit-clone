package database

import (
	"os"

	"github.com/helipoc/goapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	data, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URI")))
	if err != nil {
		return DB.Error
	}

	data.AutoMigrate(&models.User{})
	data.AutoMigrate(&models.Post{})

	DB = data

	return nil
}
