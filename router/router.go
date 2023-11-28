package router

import (
	"web-service-gin/controllers/api"
	"web-service-gin/controllers/web"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	api.Routes(r.Group("/api"))
	web.Routes(r.Group("/web"))
}
