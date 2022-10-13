package main

import (
	"net/http"
	"shirt/Server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/shirtgo?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {
		var products []models.Product
		db.Find(&products)
		c.JSON(200, products)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var products models.Product
		db.First(&products, id)
		c.JSON(200, products)
	})

	r.POST("/products", func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&product)
		c.JSON(200, gin.H{"RowsAffected": result.RowsAffected})
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		var product models.Product
		var updateProduct models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.First(&updateProduct, product.ID)
		updateProduct.Gender = product.Gender
		updateProduct.Style = product.Style
		updateProduct.Size = product.Size
		updateProduct.Price = product.Price
		db.Save(updateProduct)
		c.JSON(200, updateProduct)
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var products models.Product
		db.First(&products, id)
		db.Delete(&products)
		c.JSON(200, products)
	})
	r.Use(cors.Default())
	r.Run()
}
