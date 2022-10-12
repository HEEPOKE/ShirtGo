package Controllers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func products() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Products{})

	// Create
	db.Create(&Products{Code: "D42", Price: 100})

	// Read
	var Products Products
	db.First(&Products, 1)                 // find Products with integer primary key
	db.First(&Products, "code = ?", "D42") // find Products with code D42

	// Update - update Products's price to 200
	db.Model(&Products).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&Products).Updates(Products{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&Products).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete Products
	db.Delete(&Products, 1)
}
