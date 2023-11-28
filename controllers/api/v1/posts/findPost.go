package posts

import (
	"github.com/gin-gonic/gin"

	"web-service-gin/models"
)

/*
 * Find single post
 */
func FindPost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Preload("Author").Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(404, gin.H{"ok": false, "error": "Post not found!"})
		return
	}

	c.JSON(200, gin.H{"ok": true, "post": post})
}
