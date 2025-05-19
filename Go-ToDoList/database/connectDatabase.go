package database

import (
	"log"

	"github.com/NurochmanRizal/go-ToDoList/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase()  {
	dsn := "root:mrn07052005@tcp(127.0.0.1:3306)/todolist_app?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	DB.AutoMigrate(&models.User{}, &models.Todos{})
}