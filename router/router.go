package router

import (
	"gin-based/controllers/api"
	"gin-based/controllers/web"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	api.Routes(r.Group("/api"))
	web.Routes(r.Group("/web"))
}
