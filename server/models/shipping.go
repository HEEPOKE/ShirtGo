package models

import (
	"time"

	"gorm.io/gorm"
)

type shipping struct {
	gorm.Model
	id           uint
	shippings_id uint
	status       string
	start_date   time.Time
	end_date     time.Time
}
