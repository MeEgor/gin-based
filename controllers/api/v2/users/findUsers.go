package users

import (
	"gin-based/models"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(200, gin.H{"ok": true, "users": users, "Hello": "Hello from API V2"})
}
