package web

import (
	"github.com/gin-gonic/gin"

	users "gin-based/controllers/web/users"
)

func Routes(web *gin.RouterGroup) {
	web.GET("/users", users.FindUsers)
}
