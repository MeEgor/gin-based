package api

import (
	"github.com/gin-gonic/gin"

	auth "web-service-gin/controllers/api/auth"
	v1 "web-service-gin/controllers/api/v1"
	v2 "web-service-gin/controllers/api/v2"
	"web-service-gin/middleware"
)

func Routes(api *gin.RouterGroup) {
	api.POST("/login", auth.Login)
	api.POST("/register", auth.Register)

	// Different way to use groups
	v1.Routes(api.Group("/v1"))

	v2Group := api.Group("/v2")
	{
		v2Group.Use(middleware.Authenticate())
		v2.Routes(v2Group)
	}
}
