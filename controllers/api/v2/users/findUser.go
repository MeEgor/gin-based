package users

import (
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"ok": false, "error": "User not found!"})
		return
	}

	c.JSON(200, gin.H{"ok": true, "user": user, "message": "Hello from API V2"})
}
