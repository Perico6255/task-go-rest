package config

import (
	"Perico6255/task-go-rest/internal/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:password@tcp(127.0.0.1:3306)/go_api"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error conectandose a la base de datos: %w")
	}
	db.AutoMigrate(models.UserModel{})
	DB = db
  fmt.Println("CONECTED DATABASE")
}
