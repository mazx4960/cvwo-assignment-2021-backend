package main

import (
	"log"
	"main/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		log.Println("development mode")
	}

	database.Connect()
	if len(os.Args) > 1 {
		if os.Args[1] == "--reset" {
			database.Reset()
		}
	}

	router = gin.Default()
	initializeRoutes(router)

	router.Run()
}
