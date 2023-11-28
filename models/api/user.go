package api

import (
	models "gin-based/models"
)

type User struct {
	models.Base
	Name  string `json:"name"`
	Email string `json:"email"`
}
