package controllers

import (
	"web-service-gin/middleware"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// Login Innput
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	var user models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found!"})
		return
	}

	if user.Password != input.Password {
		c.JSON(401, gin.H{"error": "Incorrect password!"})
		return
	}

	token, refreshToken, _ := middleware.GenerateAllTokens(user.Email, user.Name, user.ID)
	// set tokens to response header
	c.Header("access_token", token)
	c.Header("refresh_token", refreshToken)

	c.JSON(200, gin.H{"ok": true, "user": user})
}
