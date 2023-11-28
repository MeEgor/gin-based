package models

import uuid "github.com/gofrs/uuid"

type Post struct {
	Base
	Title  string    `json:"title"`
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
	Author User      `json:"author,omitempty" gorm:"foreignKey:UserID"`
}

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type UpdatePostInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
