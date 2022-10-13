package migrations

import (
	"shirt/Server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/shirtgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})

	db.Create(&models.Product{Gender: "Men", Style: "Plain color/Red", Size: "XS", Price: 400})
	db.Create(&models.Product{Gender: "Men", Style: "Plain color/Red", Size: "S", Price: 400})
}
