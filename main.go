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
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
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
