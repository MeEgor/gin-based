package models

type User struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Posts []Post `json:"posts"`
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
