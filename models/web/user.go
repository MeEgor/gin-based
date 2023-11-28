package models

import "gorm.io/gorm"

type User struct {
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
