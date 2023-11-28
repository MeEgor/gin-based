package main

import (
	"gin-based/models"
	"gin-based/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	models.ConnectDatabase()
	router.Routes(r)

	r.Run("localhost:8000")
}
