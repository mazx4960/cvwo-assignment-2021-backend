package database

import (
	"fmt"
	"main/models"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(mode string) {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var connection *gorm.DB
	var err error
	if mode == "production" {
		connection, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgres",
			DSN:        dbinfo + " sslmode=disable",
		}))
	} else {
		connection, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	}

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
