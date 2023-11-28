package api

import (
	models "web-service-gin/models"
)

type User struct {
	models.Base
	Name  string `json:"name"`
	Email string `json:"email"`
}
