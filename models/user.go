package models

type User struct {
	Base
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`

	Posts []Post `json:"posts,omitempty"`
}

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
