package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id     uint
	Gender string
	Style  string
	Size   string
	Price  int
}
