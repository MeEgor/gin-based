package main

import (
	"log"
	"net/http"

	// apiV1 "web-service-gin/controllers/api/v1"
	// apiV2 "web-service-gin/controllers/api/v2"
	"web-service-gin/models"
	"web-service-gin/router"

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

	// r.GET("/users", apiV1.FindUsers)
	// r.GET("/users/:id", apiV1.FindUser)
	// r.POST("/users", apiV1.CreateUser)
	// r.PATCH("/users/:id", apiV1.UpdateUser)
	// r.DELETE("/users/:id", apiV1.DeleteUser)

	// r.GET("/v2/users", apiV2.FindUsers)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run("localhost:8000")
}
