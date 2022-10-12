package Models

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Gender string
	Style  string
	Size   string
	Price  uint
}
