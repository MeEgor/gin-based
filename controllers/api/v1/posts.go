package controllers

import (
	"fmt"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func CreatePost(c *gin.Context) {
	var input models.CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Print the input struct to the console
	fmt.Printf("Input: %+v\n", input)

	post := models.Post{Title: input.Title, Body: input.Body, UserID: uuid.FromStringOrNil("1ee8c387-66a0-651a-84fd-d00eb29192d5")}
	models.DB.Create(&post)

	c.JSON(200, gin.H{"ok": true, "post": post})
}
