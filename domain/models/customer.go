package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
