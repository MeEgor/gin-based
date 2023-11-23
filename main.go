package main

import (
	"log"
	"net/http"
	"web-service-gin/controllers"
	"web-service-gin/models"

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

	r.GET("/users", controllers.FindUsers)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run("localhost:8000")
}
