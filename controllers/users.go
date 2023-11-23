package controllers

import (
	"net/http"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"ok": true, "users": users})
}
