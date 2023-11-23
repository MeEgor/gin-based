package usersController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"title"`
	Email string `json:"artist"`
}

var users = []user{
	{ID: "1", Name: "Blue Train", Email: "John@Coltrane"},
	{ID: "2", Name: "Jeru", Email: "Gerry@Mulligan"},
	{ID: "3", Name: "Sarah Vaughan and Clifford Brown", Email: "Sarah@Vaughan"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
