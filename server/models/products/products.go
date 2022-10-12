package models

import (
	"gorm.io/gorm"
)

type products struct {
	gorm.Model
	id     uint
	shirt  string
	price  uint
	style  string
	size   uint
	gender string
}
