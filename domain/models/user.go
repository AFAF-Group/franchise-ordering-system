package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"column:email;size:50;uniqueIndex"`
	Password string `gorm:"column:password"`
}
