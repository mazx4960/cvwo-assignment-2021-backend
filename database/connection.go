package database

import (
	"fmt"
	"main/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbinfo := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	connection, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Task{}, &models.Tag{}, &models.TaskTag{}, &models.TaskList{})
}

func Reset() {
	DB.Migrator().DropTable(&models.User{})
	DB.Migrator().DropTable(&models.Task{}, &models.Tag{}, &models.TaskTag{}, &models.TaskList{})
	
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Task{}, &models.Tag{}, &models.TaskTag{}, &models.TaskList{})
}