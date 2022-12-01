package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id     uint   `json:"id"`
	Gender string `json:"gender"`
	Style  string `json:"style"`
	Size   string `json:"size"`
	Price  int    `json:"price"`
}
