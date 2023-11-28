package posts

import (
	"github.com/gin-gonic/gin"

	"web-service-gin/models"
)

func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(200, gin.H{"ok": true, "posts": posts})
}
