package config

import (
	"Perico6255/task-go-rest/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/api-go"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error conectandose a la base de datos: %w")
	}
  db.AutoMigrate(models.UserModel{})
	DB = db
}