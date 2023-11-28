package router

import (
	apiAuthController "web-service-gin/controllers/api"
	apiV1Controllers "web-service-gin/controllers/api/v1"
	apiV2Controllers "web-service-gin/controllers/api/v2"
	webControllers "web-service-gin/controllers/web"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	api := r.Group("/api")
	{
		// Auth
		api.POST("/login", apiAuthController.Login)

		v1 := api.Group("/v1")
		{
			v1.GET("/users", apiV1Controllers.FindUsers)
			v1.GET("/users/:id", apiV1Controllers.FindUser)
			v1.POST("/users", apiV1Controllers.CreateUser)
			v1.PATCH("/users/:id", apiV1Controllers.UpdateUser)
			v1.DELETE("/users/:id", apiV1Controllers.DeleteUser)

			v1.POST("/posts", apiV1Controllers.CreatePost)
		}

		v2 := api.Group("/v2")
		{
			v2.GET("/users", apiV2Controllers.FindUsers)
		}
	}

	web := r.Group("/web")
	{
		web.GET("/users", webControllers.FindUsers)
	}
}
