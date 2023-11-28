package auth

import (
	"fmt"
	"web-service-gin/helpers/password"
	"web-service-gin/middleware"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// Login innput
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register input
type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	// PassConfirm string `json:"pass_confirm" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	var user models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// print input
	fmt.Printf("Login Input: %+v\n", input)

	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(422, gin.H{"error": "E-Mail or Password is incorrect"})
		return
	}

	if ok := password.CheckPasswordHash(input.Password, user.HashedPassword); !ok {
		c.JSON(422, gin.H{"error": "E-Mail or Password is incorrect"})
		return
	}

	token, refreshToken, _ := middleware.GenerateAllTokens(user.Email, user.Name, user.ID)
	// set tokens to response header
	c.Header("access_token", token)
	c.Header("refresh_token", refreshToken)

	c.JSON(200, gin.H{"ok": true, "user": user})
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"ok": false, "error": err.Error()})
		return
	}
	// print input
	fmt.Printf("Register Input: %+v\n", input)
	// create password
	hashedPassword, _ := password.HashPassword(input.Password)
	// print hashed password
	fmt.Printf("Hashed Password: %+v\n", hashedPassword)

	user := models.User{Name: input.Name, Email: input.Email, HashedPassword: hashedPassword}
	models.DB.Create(&user)

	c.JSON(200, gin.H{"ok": true, "user": user})
}
