package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unit_price"`
}
