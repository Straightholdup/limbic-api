package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	SystemPaymentId int `gorm:"default:0"`
	Amount          int
	UserId          int
}
