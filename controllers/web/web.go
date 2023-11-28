package web

import (
	"github.com/gin-gonic/gin"

	users "web-service-gin/controllers/web/users"
)

func Routes(web *gin.RouterGroup) {
	web.GET("/users", users.FindUsers)
}
