package controllers

import (
	// "ShirtGo/server/models/products/"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&products{})

	// Create
	db.Create(&products{Code: "D42", Price: 100})

	// Read
	var products products
	db.First(&products, 1)                 // find products with integer primary key
	db.First(&products, "code = ?", "D42") // find products with code D42

	// Update - update products's price to 200
	db.Model(&products).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&products).Updates(products{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&products).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete products
	db.Delete(&products, 1)
}
