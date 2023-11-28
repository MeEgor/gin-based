package posts

import (
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

	userId := c.MustGet("id").(uuid.UUID)
	post := models.Post{Title: input.Title, Body: input.Body, UserID: userId}
	models.DB.Create(&post)

	c.JSON(200, gin.H{"ok": true, "post": post})
}
