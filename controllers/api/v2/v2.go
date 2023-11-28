package v2

import (
	"github.com/gin-gonic/gin"

	"web-service-gin/controllers/api/v2/users"
)

func Routes(v2 *gin.RouterGroup) {
	v2.GET("/users", users.FindUsers)
	v2.GET("/users/:id", users.FindUser)
}
