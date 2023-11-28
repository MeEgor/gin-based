package controllers

import (
	models "web-service-gin/models"
	webModels "web-service-gin/models/web"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	var users []webModels.User
	models.DB.Find(&users)

	c.JSON(200, gin.H{"ok": true, "users": users, "message": "Hello from WEB"})
}
