package users

import (
	models "web-service-gin/models"
	apiModels "web-service-gin/models/api"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	var users []apiModels.User
	models.DB.Find(&users)

	c.JSON(200, gin.H{"ok": true, "users": users, "message": "Hello from API V1"})
}

func FindUser(c *gin.Context) {
	var user apiModels.User

	if err := models.DB.Preload("Posts").Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"ok": false, "error": "User not found!"})
		return
	}

	c.JSON(200, gin.H{"ok": true, "user": user, "message": "Hello from API V1"})
}

func CreateUser(c *gin.Context) {
	var input models.CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"ok": false, "error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Email: input.Email}
	models.DB.Create(&user)

	c.JSON(200, gin.H{"ok": true, "user": user, "message": "Hello from API V1"})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"ok": false, "error": "User not found!"})
		return
	}

	var input models.UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"ok": false, "error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(200, gin.H{"ok": true, "user": user, "message": "Hello from API V1"})
}

func DeleteUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"ok": false, "error": "User not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(200, gin.H{"ok": true, "message": "User deleted successfully!"})
}
