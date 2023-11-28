package v1

import (
	posts "web-service-gin/controllers/api/v1/posts"
	users "web-service-gin/controllers/api/v1/users"
	"web-service-gin/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(v1 *gin.RouterGroup) {
	v1.Use(middleware.Authenticate())

	v1.GET("/users", users.FindUsers)
	v1.GET("/users/:id", users.FindUser)
	v1.POST("/users", users.CreateUser)
	v1.PUT("/users/:id", users.UpdateUser)
	v1.PATCH("/users/:id", users.UpdateUser)
	v1.DELETE("/users/:id", users.DeleteUser)

	v1.POST("/posts", posts.CreatePost)
	v1.GET("/posts", posts.FindPosts)
	v1.GET("/posts/:id", posts.FindPost)
}
