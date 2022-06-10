package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string  `json:"email" gorm:"uniqueIndex"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Avatar   *string `json:"avatar"`
}
