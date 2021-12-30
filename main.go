package main

import (
	"flag"
	"log"
	"main/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func main() {
	var mode string;
	var reset bool;
	flag.StringVar(&mode, "mode", "dev", "dev or prod")
	flag.BoolVar(&reset, "reset", false, "reset database")
	flag.Parse()

	if mode == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		log.Println("Running in development mode...")

		// Reset database
		if reset {
			database.Reset()
		}
		database.Connect("development")
	} else { // production mode
		log.Println("Running in production mode...")
		
		gin.SetMode(gin.ReleaseMode)
		database.Connect("production")
	}

	router = gin.Default()
	initializeRoutesv1(router)
	initializeRoutesv2(router)

	router.Run()
}
