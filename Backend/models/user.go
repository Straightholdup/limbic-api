package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
