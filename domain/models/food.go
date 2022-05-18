package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name      string `gorm:"column:name"`
	Quantity  int    `gorm:"column:quantity"`
	UnitPrice int    `gorm:"unit_price"`
}
