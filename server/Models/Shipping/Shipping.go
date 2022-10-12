package Models

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	status     string
	start_date time.Time
	end_date   time.Time
}
